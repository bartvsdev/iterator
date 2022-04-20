# Iterator

The iterator library allows users to iterate over collections using a functional approach. 

## Requirements

Iterator requires Go version 1.18 or later.

## Usage

A simple example using the well-known `Map` and `Fold` functions. 

```go
func main() {
	quad := func(a int) int { return a * a }
	add := func(a, b int) int { return a + b }

	iter := iterator.Slice([]int{1, 2, 3})
	quads := iterator.Map(iter, quad)
	sum, _ := iterator.Fold(quads, 0, add)
	fmt.Println(sum)
}
```

A more involved sample using `MapErr` (possibly erroneous) mapping strings to integers using `Atoi` from the standard library. An error is returned on the value `"three"` when the iterator is consumed by `ToSlice`.

```go
func main() {
	iter := iterator.Slice([]string{"3", "03", "three", "-3"})
	mapped := iterator.MapErr(iter, strconv.Atoi)
	numbers, err := iterator.ToSlice(mapped)
	if err != nil {
		// panic: strconv.Atoi: parsing "three": invalid syntax
		panic(err)
	}
	// Never reached
	fmt.Println(numbers)
}
```

## Benchmarks
*WIP*

