package main

import "fmt"

var a = "this is stored in the variable a"     // package scope
var b, c string = "stored in b", "stored in c" // package scope
var d string                                   // package scope

func main() {

	d = "stored in d" // declaration above; assignment here; package scope
	var e = 42        // function scope - subsequent variables have func scope:
	f := 43
	g := "stored in g"
	h, i := "stored in h", "stored in i"
	j, k, l, m := 44.7, true, false, 'm' // single quotes
	n := "n"                             // double quotes
	o := `o`                             // back ticks

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m)
	fmt.Println("n - ", n)
	fmt.Println("o - ", o)

	fmt.Printf("a - %T \n", a)
	fmt.Printf("b - %T \n", b)
	fmt.Printf("c - %T \n", c)
	fmt.Printf("d - %T \n", d)
	fmt.Printf("e - %T \n", e)
	fmt.Printf("f - %T \n", f)
	fmt.Printf("g - %T \n", g)
	fmt.Printf("h - %T \n", h)
	fmt.Printf("i - %T \n", i)
	fmt.Printf("j - %T \n", j)
	fmt.Printf("k - %T \n", k)
	fmt.Printf("l - %T \n", l)
	fmt.Printf("m - %T \n", m)
	fmt.Printf("n - %T \n", n)
	fmt.Printf("o - %T \n", o)
}
