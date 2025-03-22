package test_utils

import (
	"encoding/json"
	"fmt"
	"time"
)

// JsonBuilder helps construct complex JSON payloads for testing.
type JsonBuilder struct {
	data map[string]any
}

// NewJsonBuilder creates a new MockBuilder instance.
func NewJsonBuilder() *JsonBuilder {
	return &JsonBuilder{
		data: make(map[string]any),
	}
}

// SetString adds a string field to the mock payload.
func (mb *JsonBuilder) SetString(key string, value string) *JsonBuilder {
	mb.data[key] = value
	return mb
}

// SetInt adds an integer field to the mock payload.
func (mb *JsonBuilder) SetInt(key string, value int) *JsonBuilder {
	mb.data[key] = value
	return mb
}

// SetBool adds a boolean field to the mock payload.
func (mb *JsonBuilder) SetBool(key string, value bool) *JsonBuilder {
	mb.data[key] = value
	return mb
}

// SetFloat64 adds a float64 field to the mock payload.
func (mb *JsonBuilder) SetFloat64(key string, value float64) *JsonBuilder {
	mb.data[key] = value
	return mb
}

// SetTime adds a time.Time field to the mock payload.
func (mb *JsonBuilder) SetTime(key string, value time.Time) *JsonBuilder {
	mb.data[key] = value
	return mb
}

// SetObject adds a nested object field to the mock payload.
func (mb *JsonBuilder) SetObject(key string, obj map[string]any) *JsonBuilder {
	mb.data[key] = obj
	return mb
}

// SetArray adds an array field to the mock payload.
func (mb *JsonBuilder) SetArray(key string, arr []any) *JsonBuilder {
	mb.data[key] = arr
	return mb
}

// Build returns the JSON representation of the mock payload.
func (mb *JsonBuilder) Build() ([]byte, error) {
	return json.Marshal(mb.data)
}

// BuildString returns the JSON representation of the mock payload as a string.
func (mb *JsonBuilder) BuildString() (string, error) {
	data, err := mb.Build()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Print builds and prints the JSON payload to the console for debugging.
func (mb *JsonBuilder) Print() {
	data, err := mb.BuildString()
	if err != nil {
		fmt.Println("Error building mock JSON:", err)
		return
	}
	fmt.Println(data)
}
