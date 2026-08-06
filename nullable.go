package nullable

import (
	"bytes"
	"encoding/json"
	"errors"
)

// Value is a generic type, which implements a field that can be one of three states:
//
// - field is not set in the request
// - field is explicitly set to `null` in the request
// - field is explicitly set to a valid value in the request
//
// Value is intended to be used with JSON marshalling and unmarshalling.
//
// Internal implementation details:
//
// - map[true]T means a value was provided
// - map[false]T means an explicit null was provided
// - nil or zero map means the field was not provided
//
// If the field is expected to be optional, add the `omitempty` JSON tags. Do NOT use `*Value`!
//
// Adapted from https://github.com/golang/go/issues/64515#issuecomment-1841057182
type Value[T any] map[bool]T

// Nullable is the original name of [Value], and is retained as an alias so that
// existing code continues to compile. Because it is an alias rather than a
// distinct type, `Nullable[T]` and `Value[T]` are interchangeable everywhere:
// no conversion is needed to pass one where the other is expected.
//
// New code should prefer [Value].
type Nullable[T any] = Value[T]

// NewValue is a convenience helper to allow constructing a `Value` with a given value, for instance to construct a field inside a struct, without introducing an intermediate variable
func NewValue[T any](t T) Value[T] {
	var n Value[T]
	n.Set(t)
	return n
}

// NewNullValue is a convenience helper to allow constructing a `Value` with an explicit `null`, for instance to construct a field inside a struct, without introducing an intermediate variable
func NewNullValue[T any]() Value[T] {
	var n Value[T]
	n.SetNull()
	return n
}

// NewNullableWithValue is the original name of [NewValue], and is retained so
// that existing code continues to compile.
//
// New code should prefer [NewValue].
func NewNullableWithValue[T any](t T) Nullable[T] {
	return NewValue(t)
}

// NewNullNullable is the original name of [NewNullValue], and is retained so
// that existing code continues to compile.
//
// New code should prefer [NewNullValue].
func NewNullNullable[T any]() Nullable[T] {
	return NewNullValue[T]()
}

// Get retrieves the underlying value, if present, and returns an error if the value was not present
func (t Value[T]) Get() (T, error) {
	var empty T
	if t.IsNull() {
		return empty, errors.New("value is null")
	}
	if !t.IsSpecified() {
		return empty, errors.New("value is not specified")
	}
	return t[true], nil
}

// GetOrEmpty retrieves the underlying value or returns empty value if not present or was `null`. Use Get to distinguish between these cases.
func (t Value[T]) GetOrEmpty() T {
	var empty T
	if !t.IsSpecified() || t.IsNull() {
		return empty
	}
	return t[true]
}

// MustGet retrieves the underlying value, if present, and panics if the value was not present
func (t Value[T]) MustGet() T {
	v, err := t.Get()
	if err != nil {
		panic(err)
	}
	return v
}

// Set sets the underlying value to a given value
func (t *Value[T]) Set(value T) {
	*t = map[bool]T{true: value}
}

// IsNull indicate whether the field was sent, and had a value of `null`
func (t Value[T]) IsNull() bool {
	_, foundNull := t[false]
	return foundNull
}

// SetNull indicate that the field was sent, and had a value of `null`
func (t *Value[T]) SetNull() {
	var empty T
	*t = map[bool]T{false: empty}
}

// IsSpecified indicates whether the field was sent
func (t Value[T]) IsSpecified() bool {
	return len(t) != 0
}

// SetUnspecified indicate whether the field was sent
func (t *Value[T]) SetUnspecified() {
	*t = map[bool]T{}
}

func (t Value[T]) MarshalJSON() ([]byte, error) {
	// if field was specified, and `null`, marshal it
	if t.IsNull() {
		return []byte("null"), nil
	}

	// if field was unspecified, and `omitempty` is set on the field's tags, `json.Marshal` will omit this field

	// otherwise: we have a value, so marshal it
	return json.Marshal(t[true])
}

func (t *Value[T]) UnmarshalJSON(data []byte) error {
	// if field is unspecified, UnmarshalJSON won't be called

	// if field is specified, and `null`
	if bytes.Equal(data, []byte("null")) {
		t.SetNull()
		return nil
	}
	// otherwise, we have an actual value, so parse it
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	t.Set(v)
	return nil
}
