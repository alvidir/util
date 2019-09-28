package struct

import (
	"fmt"
)

// An InterMap represents a bilateral map of unknown types
type InterMap map[interface{}]interface{}

// Nav gets recursively a deep value inside a map, considering the values for each key in args are already
// other maps of interfaces, navigables as well
func (imap InterMap) Nav(args ...interface{}) (value interface{}, err error) {
	if imap == nil {
		err = fmt.Errorf("Cannot navigate the current content map %v", imap)
		return
	}

	if len(args) == 0 {
		return imap, nil
	}

	key, exists := args[0], false
	if value, exists = imap[key]; !exists {
		err = fmt.Errorf("Cannot get resource for provided key %v", key)
		return imap, err
	}

	if parse, ok := value.(map[interface{}]interface{}); ok {
		value, err = imap.Nav(parse, args[1:])
	} else if len(args) > 1 {
		err = fmt.Errorf("Cannot complete navigability, %v is not a map", value)
	}

	return
}
