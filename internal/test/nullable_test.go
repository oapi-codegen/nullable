package nullable_test

import (
	"encoding/json"
	"testing"

	"github.com/oapi-codegen/nullable"

	"github.com/stretchr/testify/require"
)

type Obj struct {
	Foo nullable.Nullable[string] `json:"foo,omitempty"` // note "omitempty" is important for fields that are optional
}

func TestNullable(t *testing.T) {
	// --- parsing from json and serializing back to JSON

	// -- case where there is an actual value
	data := `{"foo":"bar"}`
	// deserialize from json
	myObj := parse(data, t)
	require.Equal(t, myObj, Obj{Foo: nullable.Nullable[string]{true: "bar"}})
	require.False(t, myObj.Foo.IsNull())
	require.True(t, myObj.Foo.IsSpecified())
	value, err := myObj.Foo.Get()
	require.NoError(t, err)
	require.Equal(t, "bar", value)
	require.Equal(t, "bar", myObj.Foo.MustGet())
	// serialize back to json: leads to the same data
	require.Equal(t, data, serialize(myObj, t))

	// -- case where no value is specified: parsed from JSON
	data = `{}`
	// deserialize from json
	myObj = parse(data, t)
	require.Equal(t, myObj, Obj{Foo: nil})
	require.False(t, myObj.Foo.IsNull())
	require.False(t, myObj.Foo.IsSpecified())
	_, err = myObj.Foo.Get()
	require.ErrorContains(t, err, "value is not specified")
	// serialize back to json: leads to the same data
	require.Equal(t, data, serialize(myObj, t))

	// -- case where the specified value is explicitly null
	data = `{"foo":null}`
	// deserialize from json
	myObj = parse(data, t)
	require.Equal(t, myObj, Obj{Foo: nullable.Nullable[string]{false: ""}})
	require.True(t, myObj.Foo.IsNull())
	require.True(t, myObj.Foo.IsSpecified())
	_, err = myObj.Foo.Get()
	require.ErrorContains(t, err, "value is null")
	require.Panics(t, func() { myObj.Foo.MustGet() })
	// serialize back to json: leads to the same data
	require.Equal(t, data, serialize(myObj, t))

	// --- building objects from a Go client

	// - case where there is an actual value
	myObj = Obj{}
	myObj.Foo.Set("bar")
	require.Equal(t, `{"foo":"bar"}`, serialize(myObj, t))

	// - case where the value should be unspecified
	myObj = Obj{}
	// do nothing: unspecified by default
	require.Equal(t, `{}`, serialize(myObj, t))
	// explicitly mark unspecified
	myObj.Foo.SetUnspecified()
	require.Equal(t, `{}`, serialize(myObj, t))

	// - case where the value should be null
	myObj = Obj{}
	myObj.Foo.SetNull()
	require.Equal(t, `{"foo":null}`, serialize(myObj, t))
}

func parse(data string, t *testing.T) Obj {
	var myObj Obj
	err := json.Unmarshal([]byte(data), &myObj)
	require.NoError(t, err)
	return myObj
}

func serialize(o Obj, t *testing.T) string {
	data, err := json.Marshal(o)
	require.NoError(t, err)
	return string(data)
}
