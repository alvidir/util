package algorithm

import (
	"fmt"
	"strconv"
)

// Nav gets recursively a deep value inside a map, considering the values for each key in args are already
// other maps of interfaces, navigables as well
func Nav(content map[interface{}]interface{}, args ...interface{}) (value interface{}, err error) {
	if len(args) == 0 {
		return content, nil
	}

	if content == nil {
		err = fmt.Errorf("Cannot navigate the current content map %v", content)
		return
	}

	key, exists := args[0], false
	if value, exists = content[key]; !exists {
		err = fmt.Errorf("Cannot get resource for provided key %v", key)
		return content, err
	}

	if parse, ok := value.(map[interface{}]interface{}); ok {
		value, err = Nav(parse, args[1:])
	} else if len(args) > 1 {
		err = fmt.Errorf("Cannot complete navigability, %v is not a map", value)
	}

	return
}

// Address return the pointer address as an int64
func Address(i interface{}) (ival int64) {
	iparse := fmt.Sprintf("%x", i)
	ival, _ = strconv.ParseInt(iparse, 16, 64)
	return
}
