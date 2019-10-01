package algorithm

import "fmt"

// Nav gets recursively a deep value inside a map, considering the values for each key in args are already
// other maps navigables as well
func Nav(mapp map[interface{}]interface{}, args []string) (value interface{}, err error) {
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

	if parse, ok := value.(map[interface{}]interface{}); ok {
		value, err = Nav(parse, args[1:])
	} else if len(args) > 1 {
		err = fmt.Errorf("Cannot complete navigability, %v is not a map", value)
	}

	return
}
