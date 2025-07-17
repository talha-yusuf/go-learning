# Go Variables Complete Guide

This directory contains a comprehensive demonstration of Go variable concepts, including declaration methods, scope, constants, and type conversion.

## Table of Contents

- [Overview](#overview)
- [Variable Declaration Methods](#variable-declaration-methods)
- [Multiple Variable Declaration](#multiple-variable-declaration)
- [Zero Values](#zero-values)
- [Variable Scope](#variable-scope)
- [Variable Shadowing](#variable-shadowing)
- [Redeclaration with Short Declaration](#redeclaration-with-short-declaration)
- [Type Conversion](#type-conversion)
- [Constants](#constants)
- [Running the Code](#running-the-code)

## Overview

This example demonstrates all major aspects of variable handling in Go, from basic declaration to advanced concepts like shadowing and type conversion.

## Variable Declaration Methods

### 1. Using `var` keyword without initialization
- Variables get their **zero values** when declared without initialization
- Zero values: `string` = `""`, `int` = `0`, `bool` = `false`, `float64` = `0.0`

### 2. Declaration with Initialization
- Explicitly specify type and initial value
- Example: `var name string = "value"`

### 3. Type Inference
- Go automatically infers the type from the initial value
- Example: `var city = "Lahore"` (Go infers string type)

### 4. Short Variable Declaration (`:=`)
- Most concise way to declare and initialize variables
- **Note**: Can only be used inside functions
- Example: `name := "value"`

## Multiple Variable Declaration

### Same Type
- Declare multiple variables of the same type in one line
- Can be initialized with different values

### Different Types
- Use parentheses to declare multiple variables of different types
- Each variable can have its own type and initial value

### Short Declaration with Multiple Variables
- Use `:=` to declare multiple variables at once
- At least one variable must be new for redeclaration

## Zero Values

Go automatically assigns zero values to uninitialized variables:

| Type | Zero Value |
|------|------------|
| `int` | `0` |
| `float64` | `0.0` |
| `string` | `""` (empty string) |
| `bool` | `false` |
| `slice` | `nil` |
| `map` | `nil` |
| `pointer` | `nil` |

## Variable Scope

### Package-Level Variables (Global)
- Declared outside any function
- Accessible throughout the entire package
- Can be declared individually or in groups

### Function-Level Variables
- Declared inside a function
- Accessible only within that function
- Destroyed when function execution ends

### Block-Level Variables
- Declared within code blocks (if statements, loops, etc.)
- Accessible only within the block
- Can have nested blocks with their own variables

### Loop Scope
- Variables declared in loops are scoped to the loop
- Loop counter variables are also scoped to the loop

## Variable Shadowing

- When a variable in an inner scope has the same name as one in an outer scope
- The inner variable "shadows" (hides) the outer variable
- Outer variable becomes inaccessible within the inner scope
- Outer variable regains visibility when inner scope ends

## Redeclaration with Short Declaration

- You can redeclare variables using `:=` when at least one new variable is introduced
- Useful for updating existing variables while declaring new ones
- Example: `name, id := "newName", 123` (redeclares name, declares new id)

## Type Conversion

### Explicit Type Conversion
- Use type casting to convert between compatible types
- Example: `float64(intValue)` converts int to float64

### String Conversion
- Use `strconv` package for string conversions
- `strconv.Itoa()` converts int to string
- `strconv.Atoi()` converts string to int (returns error if conversion fails)

## Constants

- Declared using the `const` keyword
- Cannot be changed after declaration
- Must be initialized at declaration
- Can be declared individually or in groups
- Can be used in expressions
- Can be declared at package level or function level

## Running the Code

To run this example:

```bash
cd go-fundamentals/01-variables
go run main.go
```

The program demonstrates all variable concepts with practical examples and output.

## Key Takeaways

1. **Zero Values**: Go automatically initializes variables with appropriate zero values
2. **Type Inference**: Go can infer types from initial values
3. **Scope Rules**: Variables are accessible only within their declared scope
4. **Shadowing**: Inner scope variables can shadow outer scope variables
5. **Short Declaration**: `:=` is convenient but only works inside functions
6. **Constants**: Use `const` for values that shouldn't change
7. **Type Safety**: Go is statically typed, requiring explicit conversion between types

This comprehensive example covers all the fundamental concepts you need to understand variables in Go! 