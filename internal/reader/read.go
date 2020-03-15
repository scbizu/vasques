// Package reader reads the JSON
package reader

import (
	"encoding/json"
	"fmt"
)

type Reader struct {
	raw map[string]interface{}
}

type valueType struct {
	v interface{}
	t string
}

func NewReader(j []byte) (*Reader, error) {
	r := new(Reader)
	r.raw = make(map[string]interface{})
	if err := json.Unmarshal(j, &r.raw); err != nil {
		return nil, fmt.Errorf("reader: %w", err)
	}
	return r, nil
}

func (r Reader) Readout() map[string]interface{} {
	return map[string]interface{}{
		"AutoGenerate": r.raw,
	}
}

// func (r *Reader) selfConvert() error {
// 	for k, raw := range r.raw {
// 		switch raw.(type) {
// 		case string:
// 			_, err := strconv.ParseInt(raw.(string), 10, 64)
// 			if err != nil {
// 				// string
// 				r.vt =append(r.vt, ) valueType{
// 					v: raw,
// 					t: "string",
// 				}
// 			}
// 			r.vt[k] = valueType{
// 				v: raw,
// 				t: "int64",
// 			}
// 		case int64:
// 			r.vt[k] = valueType{
// 				v: raw,
// 				t: "int64",
// 			}
// 		case bool:
// 			r.vt[k] = valueType{
// 				v: raw,
// 				t: "bool",
// 			}
// 			// TODO: []string , []int
// 		case map[string]interface{}:
// 			// struct
// 		}
// 	}
// 	return nil
// }
