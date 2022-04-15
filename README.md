# gocol

Gocol (Go collection) is a generic functions package for slice (I will add functions for map in the future).

For example, you can use Map function to apply specific function to elements of a slice.
```go
s := []int{1, 2, 3}
doubled := slices.Map(s, func(i int) int { return i%2 == 0 })
fmt.Printf("doubled: %v", doubled)
// Output:
// doubled: [2 4 6]
```
