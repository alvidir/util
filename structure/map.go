package structure

import (
	"fmt"
)

// An Map represents a bilateral map of unknown types
type Map map[interface{}]interface{}

// Nav gets recursively a deep value inside a map, considering the values for each key in args are already
// other maps of interfaces, navigables as well
func (mapp Map) Nav(args []string) (value interface{}, err error) {
	if mapp == nil {
		err = fmt.Errorf("Cannot navigate the current content map %v", mapp)
		return
	}

	if len(args) == 0 {
		return mapp, nil
	}

	key, exists := args[0], false
	if value, exists = mapp[key]; !exists {
		err = fmt.Errorf("Cannot get resource for provided key %v", key)
		return mapp, err
	}

	if parse, ok := value.(Map); ok {
		value, err = parse.Nav(args[1:])
	} else if len(args) > 1 {
		err = fmt.Errorf("Cannot complete navigability, %v is not a map", value)
	}

	return
}
