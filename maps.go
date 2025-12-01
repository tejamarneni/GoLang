package main

import "fmt"

func main() {
	fmt.Println("--- Complex Map using interface{} ---")

	// 1. Declaration
	// Key type: string
	// Value type: interface{} (can hold any type)
	person := map[string]interface{}{
		"name":    "Alex",
		"age":     28,
		"city":    "Seattle",
		// The value for "hobbies" is a slice of strings ([]string)
		"hobbies": []string{"hiking", "photography", "coding"},
	}

	fmt.Printf("\nPerson Map (Type: map[string]interface{}):\n")
	// Using %+v prints the map in a detailed key:value format
	fmt.Printf("%+v\n", person) 

	// ----------------------------------------------------
	// 2. Accessing Values with Type Assertion
	// ----------------------------------------------------

	fmt.Println("\n--- Accessing Values ---")

	// When you retrieve a value from a map[string]interface{}, 
	// Go gives you an 'interface{}' type.
	// You must use "Type Assertion" (value.(Type)) to convert it back to its specific type.

	// Accessing the 'name' (a string)
	name, ok := person["name"].(string)
	if ok {
		fmt.Printf("Name: %s (Type: %T)\n", name, name)
	}

	// Accessing the 'age' (an integer)
	age, ok := person["age"].(int)
	if ok {
		// Note: If you initialize with 28, Go defaults to 'int'
		fmt.Printf("Age: %d (Type: %T)\n", age, age)
	}

	// Accessing the 'hobbies' (a slice of strings)
	hobbies, ok := person["hobbies"].([]string)
	if ok {
		// You can now iterate over the slice as usual
		fmt.Printf("Hobbies: %v (Type: %T)\n", hobbies, hobbies)
		fmt.Printf("First hobby: %s\n", hobbies[0])
	} else {
        fmt.Println("Error: Hobbies could not be asserted as []string")
    }

	// Type assertion is crucial because if you assert to the wrong type,
	// the 'ok' variable will be false, allowing you to handle the error safely.
	
	// Example of failed assertion:
	badAge, ok := person["age"].(string)
	if !ok {
		fmt.Printf("\nFailed Access: Tried to assert 'age' as a string. 'ok' is %t, value is '%v'\n", ok, badAge)
	}
}
