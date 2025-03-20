package gobasics

import "fmt"

/*
Exported methods start with upperCase
and unexported methods start with lowerCase
*/

func Strings() {
	var multiLineString string = `Hi this a multi
	line string`
	fmt.Println(multiLineString)

	var b string = "cde"
	fmt.Println(b + "def") //cdedef
	fmt.Println(b)
	a := "abc"
	/*
		string are built in type in golang, string are immutable in go. They can always be reassigned but not modified.
		strings can be compared with ==, !=, < , >
		you can concatenate them using + operator.

		len() method in string can be used to print the length of things, but it works differently for everything.
		In case of string len() would return the number of byte, so for simple string in case they don't have any special characters it would work well
		but in case of string having special runes or characters, it would return the number of byte used for the special characters.
		it needs to be discussed furthur.
	*/
	fmt.Println(len(b))   // 3
	fmt.Println(len("a")) // 1
	fmt.Println(len("~")) // ?
	fmt.Println(a == b)   //false
	fmt.Println(a < b)    //true
	fmt.Println(a > b)    //false

	/*
		Character in go, a character in go is called a rune, rune is type in go.
		rune is stored in memory as int32, so rune can be stored as int32.
		Notice when we print the rune.
	*/
	var e int32 = 'a' //legal but should not be preferred.
	var d rune = 'A'
	fmt.Println(e) //prints the ASCII value of e i.e 97
	fmt.Println(d) // prints the ASCII value of rune i.e 65
}

func WaysToDeclareVariables() {
	/*	var a int = 1
		var b = 1
		c := 1

		examples of type inference

		var x1, x2 = 1, 2
		y1, y2 := 1, 2
	*/
}
