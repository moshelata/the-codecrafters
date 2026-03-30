#   Go Output Functions

Go output functions is the output that apears when we run our go code and this output functions
can be run in three different ways.
```go

Print()
Println()
Printf()

```
# Print() function 
The Print() output functions is used to print the actual value in their default format 
## Example 
```go
package main 

import "fmt"

func main (){
    var x, y = "ada" "john"

    fmt.Print(x)
    fmt.Print(y)
}
```
# Println() function 

If you want to print your output in a new line then you can use the Print\n to print your 
argument on a new line.
## Example 

```go
package main

import "fmt"

func main() {
	var My_Num, No_Hours = 70, 60
	id_Num := 75
	candidateName := "Canada"

	fmt.Print(My_Num, "\n")
	fmt.Print(No_Hours, "\n")

	fmt.Println(id_Num)
	fmt.Println(candidateName)

}
```
Note if you want to add a space between string argument we use " "
## Example
```go
package main

import "fmt"

func main() {
	My_Num, No_Hours, Name, Location:= 70, 60, "John" "Canada"

	fmt.Print(My_Num, " " No_Hours)
	
}

```
#   It is also possible to only use one Print() for printing multiple variables.
## Example
```go
package main 

import "fmt"

func main(){
    x, y, z := 10,2,"age"

    fmt.Print(x,"\n", y,"\n", z)
}
```
# The Printf() Function

The Printf() function first formats its argument based on the given formatting verb and then prints them.
%v is used to print the value of the arguments
%T is used to print the type of the arguments

## Example

```go
package main

import (
	"fmt"
)

func main() {
	carName, carNum := "Forld", 07456

	fmt.Printf("the car name is %v and the car number is %T \n", carName, carNum)
	fmt.Printf("the carname is of the type %T and the carnumber is of the type %T,\n", carName, carNum)
}

```

#   Go Formatting Verbs
Formatting Verbs for Printf() go offers so many verbs than can be used with the printf functions
this verbs can be used to with all data type:
##  Example
```go
package main

import "fmt"

func main() {
	i := 69.55
    d := "Shekiri"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

    fmt.Printf("%v\n", d)
	fmt.Printf("%#v\n", d)
	fmt.Printf("%T\n", d)
}

```
#   Integer Formatting Verbs
Integer formatting verbs are used across various go particularly in the printf function, to control the output representation of integer values.


```go
package main
import ("fmt")

func main() {
  var i = 345
 
  fmt.Printf("%b\n", i)
  fmt.Printf("%d\n", i)
  fmt.Printf("%+d\n", i)
  fmt.Printf("%o\n", i)
  fmt.Printf("%O\n", i)
  fmt.Printf("%x\n", i)
  fmt.Printf("%X\n", i)
  fmt.Printf("%#x\n", i)
  fmt.Printf("%4d\n", i)
  fmt.Printf("%-4d\n", i)
  fmt.Printf("%04d\n", i)
}
```


#   String Formatting Verbs
String formatting verbs (also known as format specifiers or placeholders) are codes used within a format string to define how a variable's value should be displayed
## Example

```go
package main
import ("fmt")

func main() {
  var txt = "Hello Word!"
 
  fmt.Printf("%s\n", txt)
  fmt.Printf("%q\n", txt)
  fmt.Printf("%8s\n", txt)
  fmt.Printf("%-8s\n", txt)
  fmt.Printf("%x\n", txt)
  fmt.Printf("% x\n", txt)
}
```

#   Boolean Formatting Verbs


