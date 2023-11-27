# Introduction to Pointers

As we have learned, a variable is a named location in memory that stores a value. We can manipulate the value of a variable by assigning a new value to it or by performing operations on it. When we assign a value to a variable, we are storing that value in a specific location in memory.

```go
x := 42
// "x" is the name of a location in memory. That location is storing the integer value of 42
```

## A pointer is a variable

A pointer is a variable that stores the *memory address* of another variable. This means that a pointer "points to" the *location* of where the data is stored *NOT* the actual data itself.

The `*` syntax defines a pointer:

```go
var p *int
```

The `&` operator generates a pointer to its operand.

```go
myString := "hello"
myStringPtr = &myString
```

## Why are pointers useful?

Pointers allow us to manipulate data in memory directly, without making copies or duplicating data. This can make programs more efficient and allow us to do things that would be difficult or impossible without them.

# Pointers

Pointers hold the memory address of a value.

The `*` syntax defines a pointer:

```go
var p *int
```

A pointer's zero value is `nil`

The & operator generates a pointer to its operand.

```go
myString := "hello"
myStringPtr = &myString
```

The * dereferences a pointer to gain access to the value

```go
fmt.Println(*myStringPtr) // read myString through the pointer
*myStringPtr = "world"    // set myString through the pointer 
```

Unlike C, Go has no [pointer arithmetic](https://www.tutorialspoint.com/cprogramming/c_pointer_arithmetic.htm)

## Just because you can doesn't mean you should

We're doing this exercise to understand that pointers **can** be used in this way. That said, pointers can be *very* dangerous. It's generally a better idea to have your functions accept non-pointers and return new values rather than mutating pointer inputs.

# Nil Pointers

Pointers can be very dangerous.

If a pointer points to nothing (the zero value of the pointer type) then dereferencing it will cause a runtime error (a [panic](https://gobyexample.com/panic)) that crashes the program. Generally speaking, whenever you're dealing with pointers you should check if it's `nil` before trying to dereference it.

# Pointer Receivers

A receiver type on a method can be a pointer.

Methods with pointer receivers can modify the value to which the receiver points. Since methods often need to modify their receiver, pointer receivers are *more common* than value receivers.

## Pointer receiver

```go
type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}

func main() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "blue"
}
```

## Non-pointer receiver

```go
type car struct {
	color string
}

func (c car) setColor(color string) {
	c.color = color
}

func main() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "white"
}
```

Mthods with pointer receivers don't require that a pointer is used to call the method. The pointer will automatically be derived from the value.

```go
type circle struct {
	x int
	y int
    radius int
}

func (c *circle) grow(){
    c.radius *= 2
}

func main(){
    c := circle{
        x: 1,
        y: 2,
        radius: 4,
    }

    // notice c is not a pointer in the calling function
    // but the method still gains access to a pointer to c
    c.grow()
    fmt.Println(c.radius)
    // prints 8
}
```
