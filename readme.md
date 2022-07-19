Package atox converts ASCII to numbers. It assumes base 10. (TODO: Should it use base 0, i.e. instead accept arbitrary Go constants, like "0b1011" and "10_000"?)

It is a generic wrapper for the Parse* functions in the strconv package.

It is mainly helpful when you are writing other generic code that needs to parse number out of strings.

Sample use:

```go
half, err := atox.N[float64]("1.5")

half := atox.Must[float32]("1.5")

eleven, err := atox.N[int]("11")
```
