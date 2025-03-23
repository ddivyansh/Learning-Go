package main

/*
Custom imports of packages using multi-imports
modulesName/package name
*/
import (
	"fmt"
	"learninggo/gobasics"
)

func main() {
	gobasics.Strings()
	gobasics.WaysToDeclareVariables()
	var ans int = gobasics.Divide(6, 2)
	var err error = nil
	fmt.Printf("the answer is : %v \n", ans)
	ans, err = gobasics.DivideAndThrowError(0, 2)
	if err != nil {
		fmt.Println("The error encountered is", err.Error())
	} else {
		fmt.Printf("the answer is : %v", ans)
	}
	gobasics.ArraysInGo()
	gobasics.SlicingInGo()
}
