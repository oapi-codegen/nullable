# oapi-codegen/nullable

> An implementation of a `Value` (formerly `Nullable`) type for JSON bodies, indicating whether the field is absent, set to null, or set to a value

Unlike other known implementations, this makes it possible to both marshal and unmarshal the value, as well as represent all three states:

- the field is _not set_
- the field is _explicitly set to null_
- the field is _explicitly set to a given value_

And can be embedded in structs, for instance with the following definition:

```go
obj := struct {
		// RequiredID is a required, nullable field
		RequiredID     nullable.Value[int]     `json:"id"`
		// OptionalString is an optional, nullable field
		// NOTE that no pointer is required, only `omitempty`
		OptionalString nullable.Value[string] `json:"optionalString,omitempty"`
}{}
```

## Usage

> [!IMPORTANT]
> Although this project is under the [oapi-codegen org](https://github.com/oapi-codegen) for the `oapi-codegen` OpenAPI-to-Go code generator, this is intentionally released as a separate, standalone library which can be used by other projects.

First, add to your project with:

```sh
go get github.com/oapi-codegen/nullable
```

Check out the examples in [the package documentation on pkg.go.dev](https://pkg.go.dev/github.com/oapi-codegen/nullable) for more details.

## Naming note

The type was originally called `Nullable[T]`, and is now called `Value[T]`. `Nullable[T]` remains available as a type alias, so the two are interchangeable — existing code keeps compiling, and you can pass a `Nullable[T]` anywhere a `Value[T]` is expected without a conversion.

Both sets of constructors are available and do the same thing:

```go
// Preferred
n := nullable.NewValue(123)
nNull := nullable.NewNullValue[int]()

// Original names, still supported
o := nullable.NewNullableWithValue(123)
oNull := nullable.NewNullNullable[int]()
```

## Credits

- [KumanekoSakura](https://github.com/KumanekoSakura), [via](https://github.com/golang/go/issues/64515#issuecomment-1842973794)
- [Sebastien Guilloux](https://github.com/sebgl), [via](https://github.com/sebgl/nullable/)

As well as contributions from:

- [Jamie Tanna](https://www.jvt.me)
- [Ashutosh Kumar](https://github.com/sonasingh46)

## License

Licensed under the Apache-2.0 license.
