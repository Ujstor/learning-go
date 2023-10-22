## Struct 

We can define new types of containers of other properties or fields in Go just like in other programming languages. For example, we can create a type called `person` to represent a person, with fields name and age. We call this kind of type a `struct`.
```Go
type person struct {
	name string
	age int
}
```
Look how easy it is to define a `struct`!

There are two fields.

- `name` is a `string` used to store a person's name.
- `age` is a `int` used to store a person's age.

Let's see how to use it.
```Go
type person struct {
	name string
	age int
}

var P person  // p is person type

P.name = "Astaxie"  // assign "Astaxie" to the field 'name' of p
P.age = 25  // assign 25 to field 'age' of p
fmt.Printf("The person's name is %s\n", P.name)  // access field 'name' of p
```
There are three more ways to initialize a struct.

- Assign initial values by order
```Go
P := person{"Tom", 25}
```
- Use the format `field:value` to initialize the struct without order
```Go
P := person{age:24, name:"Bob"}
```
- Define an anonymous struct, then initialize it
```Go
P := struct{name string; age int}{"Amy",18}
```		
Let's see a complete example.

```Go
package main

import "fmt"

// define a new type
type person struct {
	name string
	age  int
}

// struct is passed by value
// compare the age of two people, then return the older person and differences of age
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var tom person

	tom.name, tom.age = "Tom", 18
	bob := person{age: 25, name: "Bob"}
	paul := person{"Paul", 43}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, bob.name, tb_Older.name, tb_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, paul.name, tp_Older.name, tp_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", bob.name, paul.name, bp_Older.name, bp_diff)
}
```
### embedded fields in struct

I've just introduced to you how to define a struct with field names and type. In fact, Go supports fields without names, but with types. We call these embedded fields.

When the embedded field is a struct, all the fields in that struct will implicitly be the fields in the struct in which it has been embedded.

Let's see one example.
```Go
package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human     // embedded field, it means Student struct includes all fields that Human has.
	specialty string
}

func main() {
	// instantiate and initialize a student
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}

	// access fields
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His specialty is ", mark.specialty)

	// modify mark's specialty
	mark.specialty = "AI"
	fmt.Println("Mark changed his specialty")
	fmt.Println("His specialty is ", mark.specialty)

	fmt.Println("Mark become old. He is not an athlete anymore")
	mark.age = 46
	mark.weight += 60
	fmt.Println("His age is", mark.age)
	fmt.Println("His weight is", mark.weight)
}

```
![](https://i.imgur.com/bTPevk6.png)

Figure 2.7 Embedding in Student and Human

We see that we can access the `age` and `name` fields in Student just like we can in Human. This is how embedded fields work. It's very cool, isn't it? Hold on, there's  something cooler! You can even use Student to access Human in this embedded field!
```Go
mark.Human = Human{"Marcus", 55, 220}
mark.Human.age -= 1
```
All the types in Go can be used as embedded fields.
```Go
package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human     // struct as embedded field
	Skills    // string slice as embedded field
	int       // built-in type as embedded field
	specialty string
}

func main() {
	// initialize Student Jane
	jane := Student{Human: Human{"Jane", 35, 100}, specialty: "Biology"}
	// access fields
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her specialty is ", jane.specialty)
	// modify value of skill field
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// modify embedded field
	jane.int = 3
	fmt.Println("Her preferred number is ", jane.int)
}

```
In the above example, we can see that all types can be embedded fields and we can use functions to operate on them.

There is one more problem however. If Human has a field called `phone` and Student has a field with same name, what should we do?

Go use a very simple way to solve it. The outer fields get upper access levels, which means when you access `student.phone`, we will get the field called phone in student, not the one in the Human struct. This feature can be simply seen as field `overloading`.
```Go
package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string // Human has phone field
}

type Employee struct {
	Human
	specialty string
	phone     string // phone in employee
}

func main() {
	Bob := Employee{Human{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}

	fmt.Println("Bob's work phone is:", Bob.phone)
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
}
```

# Object-oriented

We talked about functions and structs in the last two sections, but did you ever consider using functions as fields of a struct? In this section, I will introduce you to another form of function that has a receiver, which is called a `method`.

## method

Suppose you define a "rectangle" struct and you want to calculate its area. We'd typically use the following code to achieve this goal.
```Go
package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of r2 is: ", area(r2))
}

```
The above example can calculate a rectangle's area. We use the function called `area`, but it's not a method of the rectangle struct (like class methods in classic object-oriented languages). The function and struct are two independent things as you may notice.

It's not a problem so far. However, if you also have to calculate the area of a circle, square, pentagon, or any other kind of shape, you are going to need to add additional functions with very similar names.

![](https://i.imgur.com/KExe0qy.png)

Figure 2.8 Relationship between function and struct

Obviously that's not cool. Also, the area should really be the property of a circle or rectangle.

This is where a `method` comes to play. The `method` is a function affiliated with a type. It has similar syntax as function except, after the `func` keyword has a parameter called the `receiver`, which is the main body of that method.

Using the same example, `Rectangle.Area()` belongs directly to rectangle, instead of as a peripheral function. More specifically, `length`, `width` and `Area()` all belong to rectangle.

As Rob Pike said.

	"A method is a function with an implicit first argument, called a receiver."

Syntax of method.
```Go
func (r ReceiverType) funcName(parameters) (results)
```
Let's change our example using `method` instead.
```Go
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

// method
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

// method
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	c1 := Circle{10}
	c2 := Circle{25}
	r1 := Rectangle{9, 4}
	r2 := Rectangle{12, 2}

	fmt.Println("Area of c1 is: ", c1.Area())
	fmt.Println("Area of c2 is: ", c2.Area())
	fmt.Println("Area of r1 is: ", r1.Area())
	fmt.Println("Area of r2 is: ", r2.Area())
}
```

Notes for using methods.

- If the name of methods are the same but they don't share the same receivers, they are not the same.
- Methods are able to access fields within receivers.
- Use `.` to call a method in the struct, the same way fields are called.

![](https://i.imgur.com/BrsAkIk.png)

Figure 2.9 Methods are different in different structs

In the example above, the Area() methods belong to both Rectangle and Circle respectively, so the receivers are Rectangle and Circle.

One thing that's worth noting is that the method with a dotted line means the receiver is passed by value, not by reference. The difference between them is that a method can change its receiver's values when the receiver is passed by reference, and it gets a copy of the receiver when the receiver is passed by value.

Can the receiver only be a struct? Of course not. Any type can be the receiver of a method. You may be confused about customized types. Struct is a special kind of customized type -there are more customized types.

Use the following format to define a customized type.
```Go
type typeName typeLiteral
```
Examples of customized types:

```Go
type age int
type money float32
type months map[string]int

m := months {
	"January":31,
	"February":28,
	...
	"December":31,
}
```

I hope that you know how to use customized types now. Similar to `typedef` in C, we use `ages` to substitute `int` in the above example.

Let's get back to talking about `method`.

You can use as many methods in custom types as you want.
```Go
package main

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Box struct {
	width, height, depth float64
	color Color
}
type Color byte
type BoxList []Box //a slice of boxes

// method
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

// method with a pointer receiver
func (b *Box) SetColor(c Color) {
	b.color = c
}

// method
func (bl BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

// method
func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

// method
func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestsColor().String())

	// Let's paint them all black
	boxes.PaintItBlack()

	fmt.Println("The color of the second one is", boxes[1].color.String())
	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestsColor().String())
}
```

We define some constants and customized types.

- Use `Color` as alias of `byte`.
- Define a struct `Box` which has fields height, width, length and color.
- Define a struct `BoxList` which has `Box` as its field.

Then we defined some methods for our customized types.

- `Volume()` uses Box as its receiver and returns the volume of Box.
- `SetColor(c Color)` changes Box's color.
- `BiggestsColor()` returns the color which has the biggest volume.
- `PaintItBlack()` sets color for all Box in BoxList to black.
- `String()` use Color as its receiver, returns the string format of color name.

Is it much clearer when we use words to describe our requirements? We often write our requirements before we start coding.

### Use pointer as receiver

Let's take a look at `SetColor` method. Its receiver is a pointer of Box. Yes, you can use `*Box` as a receiver. Why do we use a pointer here? Because we want to change Box's color in this method. Thus, if we don't use a pointer, it will only change the value inside a copy of Box.

If we see that a receiver is the first argument of a method, it's not hard to understand how it works.

You might be asking why we aren't using `(*b).Color=c` instead of `b.Color=c` in the `SetColor()` method. Either one is OK here because Go knows how to interpret the assignment. Do you think Go is more fascinating now?

You may also be asking whether we should use `(&bl[i]).SetColor(BLACK)` in `PaintItBlack` because we pass a pointer to `SetColor`. Again, either one is OK because Go knows how to interpret it!

### Inheritance of method

We learned about inheritance of fields in the last section. Similarly, we also have method inheritance in Go. If an anonymous field has methods, then the struct that contains the field will have all the methods from it as well.
```Go
package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // anonymous field
	school string
}

type Employee struct {
	Human
	company string
}

// define a method in Human
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}

	sam.SayHi()
	mark.SayHi()
}
```
### Method Overriding

If we want Employee to have its own method `SayHi`, we can define a method that has the same name in Employee, and it will hide `SayHi` in Human when we call it.
```Go
package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func main() {
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}

	sam.SayHi()
	mark.SayHi()
}

```
You are able to write an Object-oriented program now, and methods use rule of capital letter to decide whether public or private as well.
