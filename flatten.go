package seq

import (
	"fmt"
	"strconv"
)

func propertyPath(key string, name string) string {
	if key == "" {
		return name
	}
	return key + "." + name
}

func flatten(key string, x interface{}) map[string]string {

	var res = make(map[string]string)
	switch vv := x.(type) {
	case string:
		res[key] = vv
	case []interface{}:
		res[fmt.Sprintf("%s.len", key)] = strconv.Itoa(len(vv))
		for ii, iv := range vv {
			var prefix = fmt.Sprintf("%s[%v]", key, ii)
			for ivk, ivv := range flatten(prefix, iv) {
				res[ivk] = ivv
			}
		}

	case Map:
		for ik, iv := range vv {
			var prefix = propertyPath(key, ik)
			for ivk, ivv := range flatten(prefix, iv) {
				res[ivk] = ivv
			}
		}
	case map[string]interface{}:
		for ik, iv := range vv {
			var prefix = propertyPath(key, ik)
			for ivk, ivv := range flatten(prefix, iv) {
				res[ivk] = ivv
			}
		}
	default:
		res[key] = string(marshal(vv))
	}

	return res
}
