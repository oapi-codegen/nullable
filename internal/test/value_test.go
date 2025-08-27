package nullable_test

import (
	"encoding/json"
	"testing"

	"github.com/oapi-codegen/nullable"

	"github.com/stretchr/testify/require"
)

type ObjV struct {
	Foo nullable.Value[string] `json:"foo,omitempty"` // note "omitempty" is important for fields that are optional
}

func TestValue(t *testing.T) {
	// --- parsing from json and serializing back to JSON

	// -- case where there is an actual value
	data := `{"foo":"bar"}`
	// deserialize from json
	myObj := parseV(data, t)
	require.Equal(t, myObj, ObjV{Foo: nullable.Value[string]{true: "bar"}})
	require.False(t, myObj.Foo.IsNull())
	require.True(t, myObj.Foo.IsSpecified())
	value, err := myObj.Foo.Get()
	require.NoError(t, err)
	require.Equal(t, "bar", value)
	require.Equal(t, "bar", myObj.Foo.MustGet())
	// serialize back to json: leads to the same data
	require.Equal(t, data, serializeV(myObj, t))

	// -- case where no value is specified: parsed from JSON
	data = `{}`
	// deserialize from json
	myObj = parseV(data, t)
	require.Equal(t, myObj, ObjV{Foo: nil})
	require.False(t, myObj.Foo.IsNull())
	require.False(t, myObj.Foo.IsSpecified())
	_, err = myObj.Foo.Get()
	require.ErrorContains(t, err, "value is not specified")
	// serialize back to json: leads to the same data
	require.Equal(t, data, serializeV(myObj, t))

	// -- case where the specified value is explicitly null
	data = `{"foo":null}`
	// deserialize from json
	myObj = parseV(data, t)
	require.Equal(t, myObj, ObjV{Foo: nullable.Value[string]{false: ""}})
	require.True(t, myObj.Foo.IsNull())
	require.True(t, myObj.Foo.IsSpecified())
	_, err = myObj.Foo.Get()
	require.ErrorContains(t, err, "value is null")
	require.Panics(t, func() { myObj.Foo.MustGet() })
	// serialize back to json: leads to the same data
	require.Equal(t, data, serializeV(myObj, t))

	// --- building objects from a Go client

	// - case where there is an actual value
	myObj = ObjV{}
	myObj.Foo.Set("bar")
	require.Equal(t, `{"foo":"bar"}`, serializeV(myObj, t))

	// - case where the value should be unspecified
	myObj = ObjV{}
	// do nothing: unspecified by default
	require.Equal(t, `{}`, serializeV(myObj, t))
	// explicitly mark unspecified
	myObj.Foo.SetUnspecified()
	require.Equal(t, `{}`, serializeV(myObj, t))

	// - case where the value should be null
	myObj = ObjV{}
	myObj.Foo.SetNull()
	require.Equal(t, `{"foo":null}`, serializeV(myObj, t))
}

func TestValueConstructors(t *testing.T) {
	// NewValue sets a concrete value
	v := nullable.NewValue(123)
	require.True(t, v.IsSpecified())
	require.False(t, v.IsNull())
	require.Equal(t, 123, v.MustGet())

	// NewNullValue sets an explicit null
	n := nullable.NewNullValue[int]()
	require.True(t, n.IsSpecified())
	require.True(t, n.IsNull())
	_, err := n.Get()
	require.ErrorContains(t, err, "value is null")
}

func parseV(data string, t *testing.T) ObjV {
	var myObj ObjV
	err := json.Unmarshal([]byte(data), &myObj)
	require.NoError(t, err)
	return myObj
}

func serializeV(o ObjV, t *testing.T) string {
	data, err := json.Marshal(o)
	require.NoError(t, err)
	return string(data)
}
