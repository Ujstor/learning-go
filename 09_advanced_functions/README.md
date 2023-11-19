# First Class and Higher Order Functions

A programming language is said to have "first-class functions" when functions in that language are treated like any other variable. For example, in such a language, a function can be passed as an argument to other functions, can be returned by another function and can be assigned as a value to a variable.

A function that returns a function or accepts a function as input is called a Higher-Order Function.

Go supports [first-class](https://developer.mozilla.org/en-US/docs/Glossary/First-class_Function) and higher-order functions.  Another way to think of this is that a function is just another type -- just like `int`s and `string`s and `bool`s.

For example, to accept a function as a parameter:

```go
func add(x, y int) int {
  return x + y
}

func mul(x, y int) int {
  return x * y
}

// aggregate applies the given math function to the first 3 inputs
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
  return arithmetic(arithmetic(a, b), c)
}

func main(){
  fmt.Println(aggregate(2,3,4, add))
  // prints 9
  fmt.Println(aggregate(2,3,4, mul))
  // prints 24
}
```

# Currying

Function currying is the practice of writing a function that takes a function (or functions) as input, and returns a new function.

For example:

```go
func main() {
  squareFunc := selfMath(multiply)
  doubleFunc := selfMath(add)
  
  fmt.Println(squareFunc(5))
  // prints 25

  fmt.Println(doubleFunc(5))
  // prints 10
}

func multiply(x, y int) int {
  return x * y
}

func add(x, y int) int {
  return x + y
}

func selfMath(mathFunc func(int, int) int) func (int) int {
  return func(x int) int {
    return mathFunc(x, x)
  }
}
```

# Defer

The defer keyword is a fairly unique feature of Go. It allows a function to be executed automatically *just before* its enclosing function returns.

The deferred calls arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

Deferred functions are typically used to close database connections, file handlers and the like.

For example:

```go
// CopyFile copies a file from srcName to dstName on the local filesystem.
func CopyFile(dstName, srcName string) (written int64, err error) {

  // Open the source file
  src, err := os.Open(srcName)
  if err != nil {
    return
  }
  // Close the source file when the CopyFile function returns
  defer src.Close()

  // Create the destination file
  dst, err := os.Create(dstName)
  if err != nil {
    return
  }
  // Close the destination file when the CopyFile function returns
  defer dst.Close()

  return io.Copy(dst, src)
}
```

In the above example, the `src.Close()` function is not called until after the `CopyFile` function was called but immediately before the `CopyFile` function returns.

Defer is a great way to **make sure** that something happens at the end of a function, even if there are multiple return statements.

# Closures

A closure is a function that references variables from outside its own function body. The function may access and *assign* to the referenced variables.

In this example, the `concatter()` function returns a function that has reference to an *enclosed* `doc` value. Each successive call to `harryPotterAggregator` mutates that same `doc` variable.

```go
func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func main() {
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
	// Mr. and Mrs. Dursley of number four, Privet Drive
}
```

# Anonymous Functions

Anonymous functions are true to form in that they have *no name*. We've been using them throughout this chapter, but we haven't really talked about them yet.

Anonymous functions are useful when defining a function that will only be used once or to create a quick [closure](https://en.wikipedia.org/wiki/Closure_(computer_programming)).

```go
// doMath accepts a function that converts one int into another
// and a slice of ints. It returns a slice of ints that have been
// converted by the passed in function.
func doMath(f func(int) int, nums []int) []int {
	var results []int
	for _, n := range nums {
		results = append(results, f(n))
	}
	return results
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	
    // Here we define an anonymous function that doubles an int
    // and pass it to doMath
	allNumsDoubled := doMath(func(x int) int {
	    return x + x
	}, nums)
	
	fmt.Println(allNumsDoubled)
    // prints:
    // [2 4 6 8 10]
}
```