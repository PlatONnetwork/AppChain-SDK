package transaction

import (
	"errors"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
	"github.com/PlatONnetwork/PlatON-Go/common/json"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"reflect"
	"strings"
)

func Send(method string, inputs []string, t reflect.Type, v reflect.Value, opt *bind.TransactOpts) (*types.Transaction, error) {
	lowerMethod := strings.ToLower(method)

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		//fmt.Println(m.Name)
		if strings.ToLower(m.Name) == lowerMethod {
			num := m.Type.NumIn()
			//fmt.Println("field string", m.Type.In(0).String(), m.Type.In(1).String())

			if len(inputs) != num-2 {
				return nil, fmt.Errorf("inputs mismatch, expect:%d, actual:%d", num, len(inputs))
			}
			args := []reflect.Value{reflect.ValueOf(opt)}
			for i := 2; i < num; i++ {
				field := m.Type.In(i)
				//fmt.Println("field string", field.String())

				isPointer := false
				newType := field
				if field.Kind() == reflect.Ptr {
					newType = field.Elem()
					isPointer = true
				}
				s := reflect.New(newType)

				c := reflect.ValueOf(s.Interface())
				call := c.MethodByName("UnmarshalText")
				if call.IsValid() {
					call.Call([]reflect.Value{reflect.ValueOf([]byte(inputs[i-2]))})
				} else {
					if err := json.Unmarshal([]byte(inputs[i-2]), s.Interface()); err != nil {
						return nil, err
					}
				}
				if !isPointer {
					args = append(args, reflect.ValueOf(s.Elem().Interface()))
				} else {
					args = append(args, s)
				}

			}
			stakeCall := v.MethodByName(m.Name)
			res := stakeCall.Call(args)
			if res[0].IsNil() {
				return nil, res[1].Interface().(error)
			} else {
				return res[0].Interface().(*types.Transaction), nil
			}
		}
	}
	return nil, errors.New("not found method")
}
func Call(method string, inputs []string, t reflect.Type, v reflect.Value, opt *bind.CallOpts) ([]byte, error) {
	lowerMethod := strings.ToLower(method)

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		//fmt.Println(m.Name)
		if strings.ToLower(m.Name) == lowerMethod {
			num := m.Type.NumIn()
			//fmt.Println("field string", m.Type.In(0).String(), m.Type.In(1).String())

			if len(inputs) != num-2 {
				return nil, fmt.Errorf("inputs mismatch, expect:%d, actual:%d", num, len(inputs))
			}
			args := []reflect.Value{reflect.ValueOf(opt)}
			for i := 2; i < num; i++ {
				field := m.Type.In(i)
				//fmt.Println("field string", field.String())

				isPointer := false
				newType := field
				if field.Kind() == reflect.Ptr {
					newType = field.Elem()
					isPointer = true
				}
				s := reflect.New(newType)

				c := reflect.ValueOf(s.Interface())
				call := c.MethodByName("UnmarshalText")
				if call.IsValid() {
					call.Call([]reflect.Value{reflect.ValueOf([]byte(inputs[i-2]))})
				} else {
					if err := json.Unmarshal([]byte(inputs[i-2]), s.Interface()); err != nil {
						return nil, err
					}
				}
				//c.MethodByName("UnmarshalText").Call([]reflect.Value{reflect.ValueOf([]byte(inputs[i-2]))})
				if !isPointer {
					args = append(args, reflect.ValueOf(s.Elem().Interface()))
				} else {
					args = append(args, s)
				}

			}
			stakeCall := v.MethodByName(m.Name)
			res := stakeCall.Call(args)
			var values []interface{}

			if !res[len(res)-1].IsNil() {
				return nil, res[1].Interface().(error)
			} else {
				for _, v := range res[:len(res)-1] {
					if bv, ok := v.Interface().([]byte); ok {
						values = append(values, hexutil.Bytes(bv))
					} else {
						values = append(values, v.Interface())
					}
				}
				return json.MarshalIndent(values, " ", " ")
			}
		}
	}
	return nil, errors.New("not found method")
}
