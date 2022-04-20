# Iterator

The iterator library allows users to iterate over collections using a functional approach. 

## Motivation
Being new to the world of Go, while having a background in `C#` and `Rust`, I missed the capabilities of `Linq` and `Iterator` respectively. Furthermore, in search of a better way to check and handle errors, and inspired by Rob Pike's great article [Errors are values](https://go.dev/blog/errors-are-values), this attempt was made.

Some of the code is inspired by the great [Soft/iter](https://github.com/Soft/iter) repository, check them out!

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

## TODO's
- Benchmarking iterator approach vs for loop
- Implementing other iterator generators/adapters/consumers
- Add error wrapping
