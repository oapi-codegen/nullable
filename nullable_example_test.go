package nullable_test

import (
	"encoding/json"
	"fmt"

	"github.com/oapi-codegen/nullable"
)

func ExampleNewNullNullable() {
	p := struct {
		N nullable.Nullable[int]
	}{}

	p.N = nullable.NewNullNullable[int]()

	fmt.Printf("Specified: %v\n", p.N.IsSpecified())
	fmt.Printf("Null: %v\n", p.N.IsNull())
	// Output:
	// Specified: true
	// Null: true
}

func ExampleNewNullableWithValue() {
	p := struct {
		N nullable.Nullable[int]
	}{}

	p.N = nullable.NewNullableWithValue(123)

	fmt.Println(p.N.Get())
	// Output:
	// 123 <nil>
}

func ExampleNullable_marshalRequired() {
	obj := struct {
		ID nullable.Nullable[int] `json:"id"`
	}{}

	// when it's not set (by default)
	b, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Unspecified:")
	fmt.Printf(`JSON: %s`+"\n", b)
	fmt.Println("---")

	// when it's not set (explicitly)
	obj.ID.SetUnspecified()

	b, err = json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Unspecified:")
	fmt.Printf(`JSON: %s`+"\n", b)
	fmt.Println("---")

	// when it's set explicitly to nil
	obj.ID.SetNull()

	b, err = json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Null:")
	fmt.Printf(`JSON: %s`+"\n", b)
	fmt.Println("---")

	// when it's set explicitly to the zero value
	var v int
	obj.ID.Set(v)

	b, err = json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Zero value:")
	fmt.Printf(`JSON: %s`+"\n", b)
	fmt.Println("---")

	// when it's set explicitly to a specific value
	v = 12345
	obj.ID.Set(v)

	b, err = json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Value:")
	fmt.Printf(`JSON: %s`+"\n", b)
	fmt.Println("---")

	// Output:
	// Unspecified:
	// JSON: {"id":0}
	// ---
	// Unspecified:
	// JSON: {"id":0}
	// ---
	// Null:
	// JSON: {"id":null}
	// ---
	// Zero value:
	// JSON: {"id":0}
	// ---
	// Value:
	// JSON: {"id":12345}
	// ---
}

func ExampleNullable_marshalOptional() {
	obj := struct {
		ID nullable.Nullable[int] `json:"id,omitempty"`
	}{}

	// when it's not set (by default)
	b, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Unspecified:")
	fmt.Printf(`JSON: %s`+"\n", b)
	fmt.Println("---")

	// when it's not set (explicitly)
	obj.ID.SetUnspecified()

	b, err = json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Unspecified:")
	fmt.Printf(`JSON: %s`+"\n", b)
	fmt.Println("---")

	// when it's set explicitly to nil
	obj.ID.SetNull()

	b, err = json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Null:")
	fmt.Printf(`JSON: %s`+"\n", b)
	fmt.Println("---")

	// when it's set explicitly to the zero value
	var v int
	obj.ID.Set(v)

	b, err = json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Zero value:")
	fmt.Printf(`JSON: %s`+"\n", b)
	fmt.Println("---")

	// when it's set explicitly to a specific value
	v = 12345
	obj.ID.Set(v)

	b, err = json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Value:")
	fmt.Printf(`JSON: %s`+"\n", b)
	fmt.Println("---")

	// Output:
	// Unspecified:
	// JSON: {}
	// ---
	// Unspecified:
	// JSON: {}
	// ---
	// Null:
	// JSON: {"id":null}
	// ---
	// Zero value:
	// JSON: {"id":0}
	// ---
	// Value:
	// JSON: {"id":12345}
	// ---
}

func ExampleNullable_unmarshalRequired() {
	obj := struct {
		Name nullable.Nullable[string] `json:"name"`
	}{}

	// when it's not set
	err := json.Unmarshal([]byte(`
		{
		}
		`), &obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Unspecified:")
	fmt.Printf("obj.Name.IsSpecified(): %v\n", obj.Name.IsSpecified())
	fmt.Printf("obj.Name.IsNull(): %v\n", obj.Name.IsNull())
	fmt.Println("---")

	// when it's set explicitly to nil
	err = json.Unmarshal([]byte(`
		{
		"name": null
		}
		`), &obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Null:")
	fmt.Printf("obj.Name.IsSpecified(): %v\n", obj.Name.IsSpecified())
	fmt.Printf("obj.Name.IsNull(): %v\n", obj.Name.IsNull())
	fmt.Println("---")

	// when it's set explicitly to the zero value
	err = json.Unmarshal([]byte(`
		{
		"name": ""
		}
		`), &obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Zero value:")
	fmt.Printf("obj.Name.IsSpecified(): %v\n", obj.Name.IsSpecified())
	fmt.Printf("obj.Name.IsNull(): %v\n", obj.Name.IsNull())
	val, err := obj.Name.Get()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("obj.Name.Get(): %#v <nil>\n", val)
	fmt.Printf("obj.Name.MustGet(): %#v\n", obj.Name.MustGet())
	fmt.Println("---")

	// when it's set explicitly to a specific value
	err = json.Unmarshal([]byte(`
		{
		"name": "foo"
		}
		`), &obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Value:")
	fmt.Printf("obj.Name.IsSpecified(): %v\n", obj.Name.IsSpecified())
	fmt.Printf("obj.Name.IsNull(): %v\n", obj.Name.IsNull())
	val, err = obj.Name.Get()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("obj.Name.Get(): %#v <nil>\n", val)
	fmt.Printf("obj.Name.MustGet(): %#v\n", obj.Name.MustGet())
	fmt.Println("---")

	// Output:
	// Unspecified:
	// obj.Name.IsSpecified(): false
	// obj.Name.IsNull(): false
	// ---
	// Null:
	// obj.Name.IsSpecified(): true
	// obj.Name.IsNull(): true
	// ---
	// Zero value:
	// obj.Name.IsSpecified(): true
	// obj.Name.IsNull(): false
	// obj.Name.Get(): "" <nil>
	// obj.Name.MustGet(): ""
	// ---
	// Value:
	// obj.Name.IsSpecified(): true
	// obj.Name.IsNull(): false
	// obj.Name.Get(): "foo" <nil>
	// obj.Name.MustGet(): "foo"
	// ---
}

func ExampleNullable_unmarshalOptional() {
	obj := struct {
		// Note that there is no pointer for nullable.Nullable when it's
		Name nullable.Nullable[string] `json:"name,omitempty"`
	}{}

	// when it's not set
	err := json.Unmarshal([]byte(`
		{
		}
		`), &obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Unspecified:")
	fmt.Printf("obj.Name.IsSpecified(): %v\n", obj.Name.IsSpecified())
	fmt.Printf("obj.Name.IsNull(): %v\n", obj.Name.IsNull())
	fmt.Println("---")

	// when it's set explicitly to nil
	err = json.Unmarshal([]byte(`
		{
		"name": null
		}
		`), &obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Null:")
	fmt.Printf("obj.Name.IsSpecified(): %v\n", obj.Name.IsSpecified())
	fmt.Printf("obj.Name.IsNull(): %v\n", obj.Name.IsNull())
	fmt.Println("---")

	// when it's set explicitly to the zero value
	err = json.Unmarshal([]byte(`
		{
		"name": ""
		}
		`), &obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Zero value:")
	fmt.Printf("obj.Name.IsSpecified(): %v\n", obj.Name.IsSpecified())
	fmt.Printf("obj.Name.IsNull(): %v\n", obj.Name.IsNull())
	val, err := obj.Name.Get()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("obj.Name.Get(): %#v <nil>\n", val)
	fmt.Printf("obj.Name.MustGet(): %#v\n", obj.Name.MustGet())
	fmt.Println("---")

	// when it's set explicitly to a specific value
	err = json.Unmarshal([]byte(`
		{
		"name": "foo"
		}
		`), &obj)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Value:")
	fmt.Printf("obj.Name.IsSpecified(): %v\n", obj.Name.IsSpecified())
	fmt.Printf("obj.Name.IsNull(): %v\n", obj.Name.IsNull())
	val, err = obj.Name.Get()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("obj.Name.Get(): %#v <nil>\n", val)
	fmt.Printf("obj.Name.MustGet(): %#v\n", obj.Name.MustGet())
	fmt.Println("---")

	// Output:
	// Unspecified:
	// obj.Name.IsSpecified(): false
	// obj.Name.IsNull(): false
	// ---
	// Null:
	// obj.Name.IsSpecified(): true
	// obj.Name.IsNull(): true
	// ---
	// Zero value:
	// obj.Name.IsSpecified(): true
	// obj.Name.IsNull(): false
	// obj.Name.Get(): "" <nil>
	// obj.Name.MustGet(): ""
	// ---
	// Value:
	// obj.Name.IsSpecified(): true
	// obj.Name.IsNull(): false
	// obj.Name.Get(): "foo" <nil>
	// obj.Name.MustGet(): "foo"
	// ---
}
