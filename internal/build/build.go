package build

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/scbizu/vasques/internal/pb"
	"github.com/scbizu/vasques/internal/reader"
)

var pkg string

// SetPackageName sets the proto package
func SetPackageName(p string) {
	pkg = p
}

// BuildFromReader builds the proto3 structure from the given reader
func BuildFromReader(r *reader.Reader) (*pb.Proto3, error) {
	pp := pb.NewPB3(pkg)
	for varName, message := range flatten(r.Readout()) {
		fsMap := make(map[string]pb.Field)
		for varName, f := range message {
			fsMap[varName] = f
		}
		pp.AddMessage(varName, fsMap)
	}
	return pp, nil
}

// flatten makes flattenMessage happy
func flatten(src map[string]interface{}) map[string]map[string]pb.Field {
	for messageName, tp := range src {
		switch tp.(type) {
		case map[string]interface{}:
			return flattenMessage(messageName, tp.(map[string]interface{}))
		}
	}
	panic("bad json formation")
}

// flattenMessage returns two layers
// the first is the message layer while the second is the message releated fields layer
func flattenMessage(message string, src map[string]interface{}) map[string]map[string]pb.Field {
	flatted := make(map[string]map[string]pb.Field)
	if len(flatted[message]) == 0 {
		flatted[message] = make(map[string]pb.Field)
	}
	for name, p := range src {
		switch p.(type) {
		case map[string]interface{}:
			n := UppercaseName(name)
			flatted[n] = make(map[string]pb.Field)
			flatted[message][n] = pb.NewField(pb.NewType(n), false, name)
			f := flattenMessage(n, p.(map[string]interface{}))
			for m, fs := range f {
				flatted[m] = fs
			}
		case string:
			_, err := strconv.ParseInt(p.(string), 10, 64)
			if err != nil {
				// parse failed, really a string
				flatted[message][name] = pb.NewField(pb.KeywordTypeString, false, name)
			} else {
				flatted[message][name] = pb.NewField(pb.KeywordTypeInt64, false, name)
			}
		case float64:
			f := big.NewFloat(p.(float64))
			if f.IsInt() {
				flatted[message][name] = pb.NewField(pb.KeywordTypeInt32, false, name)
			} else {
				flatted[message][name] = pb.NewField(pb.KeywordTypeDouble, false, name)
			}
		case bool:
			flatted[message][name] = pb.NewField(pb.KeywordTypeBool, false, name)
			// TODO: slice
		case []interface{}:
			i := p.([]interface{})[0]
			switch i.(type) {
			case string:
				if _, err := strconv.ParseInt(i.(string), 10, 64); err != nil {
					flatted[message][name] = pb.NewField(pb.KeywordTypeString, true, name)
					break
				} else {
					flatted[message][name] = pb.NewField(pb.KeywordTypeInt64, true, name)
				}
			case float64:
				f := big.NewFloat(i.(float64))
				if f.IsInt() {
					flatted[message][name] = pb.NewField(pb.KeywordTypeInt32, true, name)
				} else {
					flatted[message][name] = pb.NewField(pb.KeywordTypeDouble, true, name)
				}
			case bool:
				flatted[message][name] = pb.NewField(pb.KeywordTypeBool, true, name)
			case map[string]interface{}:
				n := UppercaseName(name)
				flatted[n] = make(map[string]pb.Field)
				flatted[message][n] = pb.NewField(pb.NewType(n), true, name)
				f := flattenMessage(n, i.(map[string]interface{}))
				for m, fs := range f {
					flatted[m] = fs
				}
			}
		}
	}
	return flatted
}

func UppercaseName(n string) string {
	bn := []byte(n)
	return fmt.Sprintf("%s%s", strings.ToUpper(string(bn[0])), string(bn[1:]))
}
