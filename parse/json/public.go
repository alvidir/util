package json

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	adapt "github.com/alvidir/util/parse/adapter"
)

// UnmarshalResponse decodes a json response from a source into target interface
func UnmarshalResponse(ctx context.Context, addr string, target interface{}) (err error) {
	client := &http.Client{}
	if deadline, ok := ctx.Deadline(); ok {
		timeout := time.Until(deadline)
		client.Timeout = timeout
	}

	var res *http.Response
	if res, err = client.Get(addr); err != nil {
		return
	}

	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}

// Unmarshal decodes an json definition to interface object
func Unmarshal(filepath string, manifest interface{}) (err error) {
	unmarshal := adapt.NewUnmarshaler(json.Unmarshal)
	return unmarshal.Path(filepath, manifest)
}

// Marshal encode an interface object to corresponding json definition
func Marshal(filepath string, content interface{}) (err error) {
	marshal := adapt.NewMarshaler(json.Marshal)
	return marshal.Path(filepath, content)
}

// Adapter builds a new adapter for json marshaling
func Adapter() adapt.Adapter {
	marshal := *adapt.NewMarshaler(json.Marshal)
	unmarshal := *adapt.NewUnmarshaler(json.Unmarshal)
	return adapt.NewEncoder(marshal, unmarshal)
}
