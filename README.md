# sqhelp
`import "github.com/earthrot/sqhelp"`

sqhelp is a library with helper functions for the `database/sql` types

## Helper functions
There are helper functions for these datatypes:

```go
NullString(value string, valid ...bool) sql.NullString
NullTime(value time.Time, valid ...bool) sql.NullTime
NullInt16(value int16, valid ...bool) sql.NullInt16
NullInt32(value int32, valid ...bool) sql.NullInt32
NullInt64(value int64, valid ...bool) sql.NullInt64
NullFloat64(value float64, valid ...bool) sql.NullFloat64
```

They all share the same parameters: the first is the value, and the second wether it is valid or not. This second parameter is optional and defaults to `true`.

## Example usage
```go
package main

import "github.com/earthrot/sqhelp"

func main() {
	a := sqhelp.NullInt32(42)
	b := sqhelp.NullString("foo")
	c := sqhelp.NullString("", false)
}
```
