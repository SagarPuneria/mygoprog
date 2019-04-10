package main

import (
	"fmt"
)

/* https://medium.com/@vivek_syngh/getting-started-with-golang-interfaces-d27121b8fa8c
Composition of Interfaces:
 Interfaces can embed methods from another interface.
 It embeds all exported (name starting with capital case) and non-exported (name starts with small case) methods from other interfaces.
*/

type Interface1 interface {
	Method1(string) string
	method2(string) string
	//Interface2 // Circular embedding of interfaces NOT allowed. Go compiler will be raise error: invalid recursive type Interface2
}

type Interface2 interface {
	Method3(string) string
	Interface1 //Embedding Interface1
}

type T struct{}

func (*T) Method1(str string) string {
	return str + " Method1"
}

func (*T) method2(str string) string {
	return str + " method2"
}

func (*T) Method3(str string) string {
	return str + " Method3"
}

/* define a method for Interface2 */
func getMethods(i2 Interface2, str string) {
	fmt.Println("Method3:", i2.Method3(str))
	fmt.Println("method2:", i2.method2(str))
	fmt.Println("Method1:", i2.Method1(str))
}
func main() {
	t := &T{}
	getMethods(t, "Interface define common method/methods that can be implemented by many different struct/one struct")
}

/*
Method3: Interface define common method/methods that can be implemented by many different struct/one struct Method3
method2: Interface define common method/methods that can be implemented by many different struct/one struct method2
Method1: Interface define common method/methods that can be implemented by many different struct/one struct Method1
*/
