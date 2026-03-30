# Introduction to Go Fundamentals

## What is Go?
Go is an open-source programming language created by Google in 2007.  
Go is easy to learn and understand compared to other programming languages.

## What is Go used for?
Go is used for server-side web development and for developing network-based programs.

## Let's get started with our first Go program

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```


##  My first Go program consists of the following parts:
-   package main
-   import
-   func main
-   fmt.Println

Package main: In Go, every program belongs to a package, and execution starts from the main package.

import: This line helps include code from other packages.

func main: This is the starting point of the program.

fmt.Println: This is a standard library function used to print output in Go.