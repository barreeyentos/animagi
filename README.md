# Animagi

A smart object mapper for Golang structures.

## Feature list
- handles copy of same Types and aliased Types
- handles nested structures
- handles pointers
    - primitive types to pointer of the same type (new memory allocated for pointer): int -> *int
    - pointer of primitive to the primitive type: *string -> string
    - pointer of primitive to pointer of primitive (deep copy, new memory allocated): *uint -> *uint
    - all of the above where the types are aliased: myint -> *int or *mystring -> string, etc.

## Usage

The `Transform` method takes in a source and destination.  The source will be mapped to the destination which means the destination must be settable (pass a pointer to the variable).

```golang
type mystring string
type myint int

src := struct {
    A int
    B mystring
    C string
}{42, "a string", "just another string"}

var dst struct {
    A myint
    B string
    D uint8
}

err := animagi.Transform(src, &dst)
```
In the above `dst` will have A and B set to `42` and `a string` and D will be default value of `0`.