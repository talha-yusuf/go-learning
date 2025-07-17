package main

import "fmt"
import "strconv"
	
// Global Variables

var globalString = "I'm a global variable"
var globalInt int = 100

// Multiple package level variable
var (
	packageVar1 string = "Package variable 1"
	packageVar2 int = 200
	packageVar3 bool = true
)

//Constants

// Constants cannot change after declaration
const PI = 3.14159
const (
	StatusActive = 1
	StatusInactive = 0
	StatusPending = 2
)

func main() {
	fmt.Println("=== GO VARIABLES COMPLETE GUIDE ===")

	//================================
	// 1. Variable Declaration Method
	//================================
	fmt.Println("1. VARIABLE DECLARATION METHODS:")

	// Method 1: Using var keyword without initialization
	var name string
	var age int
	var isActive bool
	var height float64

	fmt.Printf("Zero values - name: '%s', age: %d, isActive: %t, height: %.2f\n", name, age, isActive, height)

	// Method 2: Declaration with Initialization
	var fullName string = "Talha Yusuf"
	var userAge int = 25
	var isVerified bool = true

	fmt.Printf("Initialized - fullname: %s, userAge: %d, isVerified: %t\n", fullName, userAge, isVerified)

	// Method 3: Type Inference (Go infers the type)
	var city = "Lahore"
	var population = 8000000
	var isCapital = false
	var temperature = 72.5

	fmt.Printf("Type Ineference - city: %s, population: %d, isCapital: %t, temperature: %.1f\n", city, population, isCapital, temperature)

	// Method 4: Short variable declaration (:=) only inside functions
	country := "Pakistan"
	zipcode := 54000
	isValid := true

	fmt.Printf("Short Declaration - country: %s, zipcode: %d, isValid: %t\n", country, zipcode, isValid)

	//================================
	// 2. Multiple Variable Declaration
	//================================
	fmt.Println("2. MULTIPLE VARIABLE DECLARATION:")

	//Multiple variables of same type
	var x,y,z int
	fmt.Printf("Multiple variables of same type (zero values) - x: %d, y: %d, z: %d\n", x,y,z)

	//Multiple variables with initialization
	var a,b,c int = 1,2,3
	fmt.Printf("Multiple variables with initialization - a: %d, b: %d, c: %d\n", a,b,c)

	// Multiple variables with different types
	var (
		studentName string = "Talha"
		studentAge  int    = 30
		studentGPA  float64 = 3.8
	)
	fmt.Printf("Multiple different types - name: %s, age: %d, GPA: %.1f\n", studentName, studentAge, studentGPA)

	// Short declaration with multiple variables
	firstName, lastName, grade := "Talha", "Yusuf", 85
	fmt.Printf("Multiple short declaration - firstName: %s, lastName: %s, grade: %d\n", firstName, lastName, grade)

	//================================
	// 3. Zero Values Demonstration
	//================================
	fmt.Println("3. ZERO VALUES:")

	var (
		zeroInt 	int
		zeroFloat 	float64
		zeroString 	string
		zeroBool 	bool
		zeroSlice 	[]int
		zeroMap 	map[string]int
		zeroPointer *int
	)

	fmt.Printf("int: %d\n", zeroInt)
	fmt.Printf("float64: %f\n", zeroFloat)
	fmt.Printf("string: '%s'\n", zeroString)
	fmt.Printf("bool: %t\n", zeroBool)
	fmt.Printf("slice: %v (nil: %t)\n", zeroSlice, zeroSlice == nil)
	fmt.Printf("map: %v (nil: %t)\n", zeroMap, zeroMap == nil)
	fmt.Printf("pointer: %v (nil: %t)\n", zeroPointer, zeroPointer == nil)

	//================================
	// 4. Variable Scope Demonstration
	//================================
	fmt.Println("4. VARIABLE SCOPE:")

	// Accessing package-level variables
	fmt.Printf("Global string: %s\n", globalString)
	fmt.Printf("Global int: %d\n", globalInt)

	// Function-level variable
	functionVar := "I'm a function variable"
	fmt.Printf("Function variable: %s\n", functionVar)

	//Block level variables
	if true {
		blockVar := "I'm in a block"
		fmt.Printf("Block variable: %s\n", blockVar)

		// Nested Block
		{
			nestedVar := "I'm in nested block"
			fmt.Printf("Nested block variable: %s\n", nestedVar)
		}
		// nestedVar is not accessible here
	}
	// blockVar is not accessible here

	// Loop scope
	for i:=0; i<3; i++ {
		loopVar := fmt.Sprintf("Loop iteration %d", i) // Creates string
		fmt.Printf("Loop variable: %s\n", loopVar) // Prints the string
	}
	// i and loopVar are not accessible here

	fmt.Println()

	//================================
	// 5. Variable Shadowing
	//================================
	fmt.Println("5. VARIABLE SHADOWING:")

	shadowDemo := "outer" 
	fmt.Printf("Outer scope: %s\n", shadowDemo)

	{
		shadowDemo := "inner" // This shadows the outer variable
		fmt.Printf("Inner scope: %s\n", shadowDemo)
	}
	fmt.Printf("Back to outer scope: %s\n", shadowDemo)

	//================================
	// 6. REDECLARATION WITH SHORT DECLARATION
	//================================
	fmt.Println("6. REDECLARATION WITH SHORT DECLARATION:")
	
	username := "talha"
	fmt.Printf("Initial username: %s\n", username)
	
	// Redeclaring username with a new variable
	username, userID := "yusuf", 12345
	fmt.Printf("Redeclared username: %s, new userID: %d\n", username, userID)

	//================================
	// 7. Type Conversion
	//================================
	fmt.Println("7. TYPE CONVERSION:")

	var intVal int = 42
	var floatVal float64 = float64(intVal)
	var stringVal string = strconv.Itoa(intVal)

	fmt.Printf("Original int: %d\n", intVal)
	fmt.Printf("Converted to float64: %f\n", floatVal)
	fmt.Printf("Converted to string: %s\n", stringVal)

	// String to Int conversion
	numStr := "123"
	if num, err := strconv.Atoi(numStr); err == nil {
		fmt.Printf("String '%s' converted to int: %d\n", numStr, num)
	}

	fmt.Println()

	// ====================
	// 8. CONSTANTS DEMONSTRATION
	// ====================
	fmt.Println("8. CONSTANTS:")
	
	fmt.Printf("PI constant: %f\n", PI)
	fmt.Printf("Status Active: %d\n", StatusActive)
	fmt.Printf("Status Inactive: %d\n", StatusInactive)
	fmt.Printf("Status Pending: %d\n", StatusPending)
	
	fmt.Println()

	// ====================
	// 9. COMMON PATTERNS
	// ====================
	fmt.Println("9. COMMON PATTERNS:")

	// Swapping variables

	val1, val2 := 10,20
	fmt.Printf("Before swap: val1=%d, val2=%d\n", val1, val2)
	val1, val2 = val2, val1
	fmt.Printf("After swap: val1=%d, val2=%d\n", val1, val2)

	// ignoring values with blank identifier
	result , _ := divideNumbers(10, 2) // ignore error return
	fmt.Printf("Division result: %d\n", result)

	// Multiple assignment
	p1, p2, p3 := getPersonDetails()
	fmt.Printf("Person details: %s, %d, %s\n", p1, p2, p3)

	fmt.Println()

	// ====================
	// 10. VARIABLE NAMING CONVENTIONS
	// ====================

	fmt.Println("10. VARIABLE NAMING EXAMPLES:")
	
	// Good naming examples
	var userName string = "alice"           // camelCase
	var employeeAge int = 25                // descriptive
	var isLoggedIn bool = true              // boolean with 'is' prefix
	var maxRetries int = 3                  // constants-like naming
	var dbConnection string = "localhost"   // abbreviations OK when clear
	
	fmt.Printf("Good naming: userName=%s, employeeAge=%d, isLoggedIn=%t, maxRetries=%d, dbConnection=%s\n", userName, employeeAge, isLoggedIn, maxRetries, dbConnection)

	// Variables with underscores (less common but valid)
	var user_id int = 12345
	var max_file_size int = 1024
	
	fmt.Printf("Underscore naming: user_id=%d, max_file_size=%d\n", user_id, max_file_size)
	
	fmt.Println()


	// ====================
	// 11. ADVANCED PATTERNS
	// ====================
	fmt.Println("11. ADVANCED PATTERNS:")
	
	// Pointer variables
	var ptr *int
	num := 42
	ptr = &num
	fmt.Printf("Pointer value: %d, Address: %p\n", *ptr, ptr)

	// Slice and Map initialization
	numbers := []int{1,2,3,4,5}
	grades := map[string]int{"math": 90, "science": 85}

	fmt.Printf("Slice: %v\n", numbers)
	fmt.Printf("Map: %v\n", grades)

	// Struct variable
	type Person struct {
		Name	string
		Age		int
	}

	person1 := Person{Name: "John", Age: 30}
	var person2 Person // zero value struct

	fmt.Printf("Struct 1: %+v\n", person1)
	fmt.Printf("Struct 2 (zero value): %+v\n", person2)

	fmt.Println("\n=== END OF VARIABLES GUIDE ===")
}

	//=========================
	// Helper Functions
	//=========================

	func divideNumbers(a,b int) (int, error) {
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a/b, nil
	}

	func getPersonDetails() (string, int, string) {
		return "Talha Yusuf", 28, "BSE"
	}

