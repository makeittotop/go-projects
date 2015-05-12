package main

import (
    "fmt"
    "unsafe"
)

type MyType struct {
    Value1 int
    Value2 string
}

func main() {
    // Allocate an value of type MyType
    myValue := &MyType{10, "Bill"}

    // Create a pointer to the memory for myValue
    //  For Display Purposes
    pointer := unsafe.Pointer(myValue)

    // Display the address and values
    fmt.Printf("Addr: %v Value1 : %d Value2: %s\n",
        pointer,
        myValue.Value1,
        myValue.Value2)

    // Change the values of myValue
    ChangeMyValue(myValue)

    // Display the address and values
    fmt.Printf("Addr: %v Value1 : %d Value2: %s\n",
        pointer,
        myValue.Value1,
        myValue.Value2)
}

func ChangeMyValue(myValue *MyType) {
    // Change the values of myValue
    myValue.Value1 = 20
    myValue.Value2 = "Jill"

    foo := map[int]string{
    	1: "foo",
    	2: "boo",
    };
    
    fmt.Println(foo);
    
    // Create a pointer to the memory for myValue
    pointer := unsafe.Pointer(& myValue)

    // Display the address and values
    fmt.Printf("Addr: %v Value1 : %d Value2: %s\n",
        pointer,
        myValue.Value1,
        myValue.Value2)
}