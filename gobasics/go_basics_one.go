package gobasics

import (
	"errors"
	"fmt"
)

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

/*
this method divides numerator/denominator and the return type is int
In go we can return multiple values instead of 1.
*/
func Divide(numerator int, denominator int) int {
	fmt.Println(numerator / denominator)
	var ans int = numerator / denominator
	return ans
}

/*
This method divides two numbers and returns the ans or error.
For object default value is nil such as for error types.
We can return multiple types such as (int, error)

In go, the else and else-if statement has to be in the same line as the last }
*/
func DivideAndThrowError(numerator int, denominator int) (int, error) {
	var err error
	if numerator == 0 {
		err = errors.New("numerator can't be 0")
		return 0, err
	} else {
		return numerator / denominator, nil
	}
}

/*
Arrays are fixed size, contiguous block of memory, of the same type.
int 32 -- 4 byte
Intialisation of [] in Go

var arr [] int = [] int {1,2,3}
var arr = [] int {1,2,3}
var arr [] int -- arr with default value i.e 0
*/
func ArraysInGo() {
	//The size of the arr is 3 with 12 bytes. T
	//The values of every element is 0(default value)
	var arr [3]int32
	fmt.Println("the unitialised array would have every element as default value : ", arr)
	// end index is not included.
	fmt.Println("we can slice the array as here 1 is starting index, uptil 3 index : ", arr[1:3])
	fmt.Println("we can print the memory location of the arr as", &arr[0])
	fmt.Println("we can slice the array as arr[:3]: ", arr[:3])
	fmt.Println("we can print the size of [] : ", len(arr))

	/*
		The way to initialise the [] in go
	*/
	var arr2 [3]int = [3]int{1, 2, 3}
	fmt.Println("the initialised array is : ", arr2)
}

/*
Slicing happens to array using append() method
Slicing is a powerfule tool in go, which uses an [] under the hood
append(sliceArr, 3), here 3 is appended to sliceArr, if the capacity wasn't enough then a new array is formed
And the elements are copied into the new array
the capacity of the newly formed array is different or increased, so the nextslicing operation doesn't result in generation of a new array
however if one tries to access the indexes where there's no element present so we'd runinto a runtime panic, "index out of bound"

We can also slice or append another [] instead of a single elemnent.

	var sliceArrTwo []int = []int{5, 6, 7}
	sliceArrTwo = append(sliceArr, sliceArrTwo...)
	We use ... spread operator to do so.

len(obj) -- returns the number of characters in obj
cap(obj) -- returns the capacity of obj
*/
func SlicingInGo() {
	var sliceArr []int = []int{1, 2, 3} //Here the length of [] is automatically infered
	sliceArr = append(sliceArr, 4)
	fmt.Println("the initialised slice is : ", sliceArr)

	var sliceArrTwo []int = []int{5, 6, 7}
	sliceArrTwo = append(sliceArr, sliceArrTwo...)
	fmt.Println("the sliceArrayTwo is : ", sliceArrTwo)
	fmt.Printf("the new size of the original [] is %v and the capacity is %v ", len(sliceArrTwo), cap(sliceArrTwo))

}

/*
This function is about map data structure in golang

	here key is string, and value is an int
	var mapOne map[string]int = map[string]int{"one": 1, "two": 2}
	var mapTwo = make(map[string]int)

Accessing a key in map

	map["key"]
	var val, ok = mapOne["one"]
	if any key is not found in golang map, the value which is returned in go is default value of the key, SO be careful
	but go also return a boolean variable i.e ok (true if the key is found and false if the key is not found.

Deletion of a key in golang

	delete(mapName,KeyTobeDeleted)
*/
func MapsInGo() {
	var mapOne map[string]int = map[string]int{"one": 1, "two": 2}
	fmt.Println("the initialised map is : ", mapOne)
	var mapTwo = make(map[string]int)
	fmt.Println("the initialised map using make method is : ", mapTwo)
	var val, ok = mapOne["one"]
	if ok == true {
		fmt.Println("the value at key 'one' is : ", val)
	} else {
		fmt.Println("the key with 'one is not found '")
	}
}

/*
Iteration in case of golang.
We make use of the range keyword

	for i := range arr {
			fmt.Println(i)
		}

	for key, value := range golangMap {
			fmt.Println(key, value)
		}
*/
func IterationInGo() {
	var arr []int = []int{4, 5, 6}
	for i := range arr {
		fmt.Println(i)
	}
	var golangMap map[string]int = map[string]int{"one": 1, "two": 2}
	for key, value := range golangMap {
		fmt.Println(key, value)
	}
	/*
		A while loop in golang
	*/
	i := 0
	for {
		if i > 5 {
			break
		}
		fmt.Println("The value of 'one' is : ", i)
		i++
	}

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
}
