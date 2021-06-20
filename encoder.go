package util

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	JsonEncoder = NewEncoder(NewMarshaler(json.Marshal), NewUnmarshaler(json.Unmarshal))
	YamlEncoder = NewEncoder(NewMarshaler(yaml.Marshal), NewUnmarshaler(yaml.Unmarshal))
)

type Marshal func(interface{}) ([]byte, error)
type Unmarshal func([]byte, interface{}) error

type Marshaler interface {
	Marshal(interface{}) ([]byte, error)
	Path(string, interface{}) error
}

type Unmarshaler interface {
	Unmarshal([]byte, interface{}) error
	Path(string, interface{}) error
}

type Encoder interface {
	Marshaler() Marshaler
	Unmarshaler() Unmarshaler
}

func NewMarshaler(marshal Marshal) Marshaler {
	return &marshaler{marshal}
}

func NewUnmarshaler(unmarshal Unmarshal) Unmarshaler {
	return &unmarshaler{unmarshal}
}

func NewEncoder(marshal Marshaler, unmarshal Unmarshaler) Encoder {
	return &encoder{marshal, unmarshal}
}

type encoder struct {
	marshaler   Marshaler
	unmarshaler Unmarshaler
}

func (encoder *encoder) Marshaler() Marshaler {
	return encoder.marshaler
}

func (encoder *encoder) Unmarshaler() Unmarshaler {
	return encoder.unmarshaler
}

type marshaler struct {
	marshal Marshal
}

func (marshaler *marshaler) Marshal(v interface{}) ([]byte, error) {
	return marshaler.marshal(v)
}

func (marshaler *marshaler) Path(p string, v interface{}) (err error) {
	var data []byte
	if data, err = marshaler.marshal(v); err != nil {
		return
	}

	if _, err = os.Stat(p); err != nil {
		// if file does not exists
		return
	}

	return ioutil.WriteFile(p, data, 0644)
}

type unmarshaler struct {
	unmarshal Unmarshal
}

func (unmarshaler *unmarshaler) Unmarshal(s []byte, v interface{}) error {
	return unmarshaler.unmarshal(s, v)
}

func (unmarshaler *unmarshaler) Path(p string, v interface{}) (err error) {
	if _, err = os.Stat(p); err != nil {
		// if there is no file for the given path
		return
	}

	var content []byte
	if content, err = ioutil.ReadFile(p); err != nil {
		return
	}

	return unmarshaler.unmarshal(content, v)
}
