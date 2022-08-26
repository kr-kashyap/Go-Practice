package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

/*
	func (t *T) M() {
		if t == nil {
			fmt.Println("<nil>")
			return
		}
		fmt.Println(t.S)
	}
*/
func (f F) M() {
	fmt.Println(f)
}

func interfaces_10_to_14_26() {
	fmt.Println("\nInterface Notes")
	fmt.Println("- var i I // Creates a nil interace")
	fmt.Println("- Calling a method M() on Nil interface is a RUNTIME Error.\n")
	fmt.Println("- var i1 interface{} // Empty Interface")
	fmt.Println("- An empty interface may hold values of any type. (Every type implements at least zero methods.\n")
	fmt.Println("- i = F(3.14) // Interface on a float64 of Type F.")
	fmt.Println("- i = T{\"hello\"} // Interface on a Struct.\n\n")

	// Nil interface
	var i I
	//Return nil nil I
	fmt.Println("\nCalling a descirbe(i) on Nil interface.")
	describe(i)
	//Calling a method on Nil interface is a Runtime Error
	// i.M()

	i = F(3.14)
	fmt.Println("\nCalling a M() on i = F(3.14).")
	i.M()

	i = T{"hello"}
	fmt.Println("\nCalling a M() on i = T{\"hello\"}.")
	i.M()

	/*
		var t *T
		i = t
		fmt.Println("\nCalling a describe on i = *T.\n")
		describe(i)
		fmt.Println("\nCalling a M() on i = *T.\n")
		i.M()

		i = &T{"hello"}
		describe(i)
		i.M()
	*/

	//Empty interface
	var i1 interface{}
	fmt.Println("\nCalling a emptyDescribe() on i1 Empty interface.")
	emptyDescribe(i1)
	i1 = 42
	fmt.Println("\nCalling a emptyDescribe() on i1 Empty interface.")
	emptyDescribe(i1)

}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func emptyDescribe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
