package athengo

import (
	"athengo/helper"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/athena/athenaiface"
)

type rows struct {
	api        athenaiface.AthenaAPI
	out        *athena.GetQueryResultsOutput
	queryId    string
	done       bool
	skipHeader bool
}

type Rows interface {
	GetColumns() ([]string, error)
	GetResults(value interface{}) error
	GetRawResult(token *string) (*athena.GetQueryResultsOutput, error)
}

func NewRows(api athenaiface.AthenaAPI, queryId string) (Rows, error) {
	if queryId == "" {
		return nil, errors.New("the query id is required")
	}
	if api == nil {
		return nil, errors.New("the athena api is required")
	}

	return &rows{
		api:        api,
		queryId:    queryId,
		skipHeader: true,
		done:       false,
	}, nil
}

func (r *rows) GetColumns() ([]string, error) {
	var columns []string
	if r.out == nil {
		return columns, errors.New("please get result first")
	}
	for _, colInfo := range r.out.ResultSet.ResultSetMetadata.ColumnInfo {
		columns = append(columns, *colInfo.Name)
	}

	return columns, nil
}

func (r *rows) GetResults(value interface{}) error {
	if r.skipHeader {
		next, err := r.nextPage(nil, value)
		if err != nil {
			return err
		}
		r.done = !next
	}

	for !r.done {
		if r.out.NextToken == nil || *r.out.NextToken == "" {
			return nil
		}

		nextPage, err := r.nextPage(r.out.NextToken, value)
		if err != nil {
			return err
		}
		r.done = !nextPage
	}
	return nil
}

func (r *rows) GetRawResult(token *string) (*athena.GetQueryResultsOutput, error) {
	return r.api.GetQueryResults(&athena.GetQueryResultsInput{
		QueryExecutionId: aws.String(r.queryId),
		NextToken:        token,
	})
}

func (r *rows) nextPage(token *string, value interface{}) (bool, error) {
	var err error
	r.out, err = r.api.GetQueryResults(&athena.GetQueryResultsInput{
		QueryExecutionId: aws.String(r.queryId),
		NextToken:        token,
	})

	if err != nil {
		return false, err
	}

	offset := 0
	if r.skipHeader {
		offset = 1
		r.skipHeader = false
	}

	if len(r.out.ResultSet.Rows) < offset+1 {
		return false, nil
	}

	r.out.ResultSet.Rows = r.out.ResultSet.Rows[offset:]
	columns := r.out.ResultSet.ResultSetMetadata.ColumnInfo
	for i := 0; i < len(r.out.ResultSet.Rows); i++ {
		err := r.convertRow(columns, r.out.ResultSet.Rows[i].Data, value)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (r *rows) convertValue(athenaType string, rawValue *string) (interface{}, error) {
	if rawValue == nil {
		return nil, nil
	}

	val := *rawValue
	switch athenaType {
	case "tinyint":
		return strconv.ParseInt(val, 10, 8)
	case "smallint":
		return strconv.ParseInt(val, 10, 16)
	case "integer":
		return strconv.ParseInt(val, 10, 32)
	case "bigint":
		return strconv.ParseInt(val, 10, 64)
	case "boolean":
		switch val {
		case "true":
			return true, nil
		case "false":
			return false, nil
		}
		return nil, fmt.Errorf("cannot parse '%s' as boolean", val)
	case "float":
		return strconv.ParseFloat(val, 32)
	case "double", "decimal":
		return strconv.ParseFloat(val, 64)
	case "varchar", "string":
		return val, nil
	case "timestamp":
		return time.Parse(helper.TIMESTAMP_LAYOUT, val)
	case "timestamp with time zone":
		return time.Parse(helper.TIMESTAMP_WITH_TIME_ZONE, val)
	case "date":
		return time.Parse(helper.DATE_LAYOUT, val)
	default:
		return nil, fmt.Errorf("unknown type %s with value %s", athenaType, val)
	}
}

func (r *rows) convertRow(columns []*athena.ColumnInfo, in []*athena.Datum, ptr interface{}) error {
	v := reflect.ValueOf(ptr)
	if !v.IsValid() {
		return errors.New("the ptr is invalid")
	}

	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	} else {
		return errors.New("the input ptr is not a pointer")
	}
	var (
		crv     reflect.Value
		curLen  = 0
		element reflect.Value
	)
	if v.Kind() == reflect.Slice {
		curLen = v.Len()
		typ := v.Type().Elem()
		if typ.Kind() == reflect.Ptr {
			element = reflect.New(typ.Elem())
		} else if typ.Kind() == reflect.Struct {
			element = reflect.New(typ).Elem()
		} else {
			return errors.New("the slice ptr is not a support type")
		}
		v.Set(reflect.Append(v, element))
		temp := v.Index(curLen)
		if temp.Kind() == reflect.Ptr && !v.IsNil() {
			temp = temp.Elem()
		} else {
			return errors.New("the slice element is not a pointer")
		}
		crv = temp
	} else if v.Kind() == reflect.Struct {
		crv = v
	} else {
		return errors.New("the ptr is not a support type")
	}

	var fieldValues = make(map[string]int)
	for i := 0; i < crv.NumField(); i++ {
		if tName, exist := crv.Type().Field(i).Tag.Lookup(helper.TAG_NAME); exist {
			fieldValues[tName] = i
		}
	}

	for i, val := range in {
		record, err := r.convertValue(*columns[i].Type, val.VarCharValue)
		if err != nil {
			return err
		}
		tagVal := *columns[i].Name

		if itemNum, exist := fieldValues[tagVal]; exist {
			if crv.Field(itemNum).Kind() != reflect.ValueOf(record).Kind() {
				record, err = matchType(crv.Field(itemNum).Kind(), record)
				if err != nil {
					return err
				}
			}
			crv.Field(itemNum).Set(reflect.ValueOf(record))
		}
	}
	return nil
}
func matchType(require reflect.Kind, record interface{}) (interface{}, error) {
	switch require {
	case reflect.Uint:
		if reflect.ValueOf(record).Kind() == reflect.Int64 {
			return uint(record.(int64)), nil
		} else if reflect.ValueOf(record).Kind() == reflect.String {
			u64, err := strconv.ParseUint(record.(string), 10, 64)
			if err != nil {
				return nil, err
			}
			return uint(u64), nil
		}
	case reflect.Uint8:
		if reflect.ValueOf(record).Kind() == reflect.Int64 {
			return uint8(record.(int64)), nil
		} else if reflect.ValueOf(record).Kind() == reflect.String {
			u64, err := strconv.ParseUint(record.(string), 10, 8)
			if err != nil {
				return nil, err
			}
			return uint8(u64), nil
		}
	case reflect.Uint16:
		if reflect.ValueOf(record).Kind() == reflect.Int64 {
			return uint16(record.(int64)), nil
		} else if reflect.ValueOf(record).Kind() == reflect.String {
			u64, err := strconv.ParseUint(record.(string), 10, 16)
			if err != nil {
				return nil, err
			}
			return uint16(u64), nil
		}
	case reflect.Uint32:
		if reflect.ValueOf(record).Kind() == reflect.Int64 {
			return uint32(record.(int64)), nil
		} else if reflect.ValueOf(record).Kind() == reflect.String {
			u64, err := strconv.ParseUint(record.(string), 10, 32)
			if err != nil {
				return nil, err
			}
			return uint32(u64), nil
		}
	case reflect.Uint64:
		if reflect.ValueOf(record).Kind() == reflect.Int64 {
			return uint64(record.(int64)), nil
		} else if reflect.ValueOf(record).Kind() == reflect.String {
			u64, err := strconv.ParseUint(record.(string), 10, 64)
			if err != nil {
				return nil, err
			}
			return uint64(u64), nil
		}
	case reflect.Int:
		if reflect.ValueOf(record).Kind() == reflect.Int64 {
			return int(record.(int64)), nil
		} else if reflect.ValueOf(record).Kind() == reflect.String {
			u64, err := strconv.ParseInt(record.(string), 10, 64)
			if err != nil {
				return nil, err
			}
			return int(u64), nil
		}
	case reflect.Int8:
		if reflect.ValueOf(record).Kind() == reflect.Int64 {
			return int8(record.(int64)), nil
		} else if reflect.ValueOf(record).Kind() == reflect.String {
			u64, err := strconv.ParseInt(record.(string), 10, 8)
			if err != nil {
				return nil, err
			}
			return int8(u64), nil
		}
	case reflect.Int16:
		if reflect.ValueOf(record).Kind() == reflect.Int64 {
			return int16(record.(int64)), nil
		} else if reflect.ValueOf(record).Kind() == reflect.String {
			u64, err := strconv.ParseInt(record.(string), 10, 16)
			if err != nil {
				return nil, err
			}
			return int16(u64), nil
		}
	case reflect.Int32:
		if reflect.ValueOf(record).Kind() == reflect.Int64 {
			return int32(record.(int64)), nil
		} else if reflect.ValueOf(record).Kind() == reflect.String {
			u64, err := strconv.ParseInt(record.(string), 10, 32)
			if err != nil {
				return nil, err
			}
			return int32(u64), nil
		}
	case reflect.Int64:
		if reflect.ValueOf(record).Kind() == reflect.Int64 {
			return record, nil
		} else if reflect.ValueOf(record).Kind() == reflect.String {
			u64, err := strconv.ParseInt(record.(string), 10, 64)
			if err != nil {
				return nil, err
			}
			return int64(u64), nil
		}
	case reflect.Float32:
		if reflect.ValueOf(record).Kind() == reflect.Float64 {
			return float32(record.(float64)), nil
		} else if reflect.ValueOf(record).Kind() == reflect.String {
			u64, err := strconv.ParseFloat(record.(string), 32)
			if err != nil {
				return nil, err
			}
			return float32(u64), nil
		}
	case reflect.Float64:
		if reflect.ValueOf(record).Kind() == reflect.Float64 {
			return record, nil
		} else if reflect.ValueOf(record).Kind() == reflect.String {
			u64, err := strconv.ParseFloat(record.(string), 64)
			if err != nil {
				return nil, err
			}
			return float64(u64), nil
		}
	}
	return nil, errors.New("the require type is not support")
}
