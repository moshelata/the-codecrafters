# go varible 

go varible, has to do with storing of value just like the way we store things in our phones
wallet, cuboard and so on so talking about varible in go is simple a way of storing data 

## Why do we need to varible 
one of the most reason why we use varible is to meke coding easy for us, imagine have a chunk of list
of the names of people in learn2ear and you have to list them one after the other each time you want to
make use of them that we be a whole lot of stress but with the help of varible to store those chunk of list of the names of every candidate in learn2earn makes is more fast and efficent to code 

## Example of delcearcing a varible 
```go
//Using var as a keyword, the name of the varible the type and the value it holds
//Note: when creating a varible you must always specify either type or value (or both).

var Cluster_Name string = "Debugers"

// We also have the short declearation where the compiler automaticaly asign the type of the varible
// This is called an inferred 

dog_years := 20


```
## Varible declearation without value
what this means is simply declearing a varible without a value this type of decleara usually take te defualt value of the type.

## Example of a varible without a value 
```go
// when this code is run it takes the default value for the type 

package main 

import"fmt"

func main(){
    var n stirng

    var sum int

    var x bool

    fmt.Println(string)
    fmt.Println(int)
    fmt.Println(bool)
}
```

## Value assigned after declearation 
This practies is good only when you dont know the value of the varible 
Note: when you are using the short declearation ":=" you must assgin it a value
## Example of value assigned after declearation

```go
package main 

import "fmt"

func main(){
    var surname string
    
    surname = "Otokpa"
}
```
##  The difference between the var and := 
-   var can be using inside and outside of a function  
    := can only be used inside a function
-   var the varible declaration and value assign can be done sepratly
    eg. 
     var sumTotal int
     sumTotal = 1000
    := varible declaration and value assign can not be done sepratly it must be done on the same line
    eg.
     sumTotal := 1000 

#   Go Multiple Variable Declaration
    In go you can declear as many varible as you want that is mutiple variable declaration

## Example of multiple varible declaration 
```go
pacakge main 

import "fmt"

var month, days, name string = "march", "monday", "blessing"

func main (){

    age, num, hight := 20, 40, 80
}
```
Note if you use the type keyword you can not assign any other type of varible
Example

```go
package main

import "fmt"

func main(){
    var name, num = "Oche", 84
    word, dogage := "firm", 10
}
```
# Go Variable Declaration in a Block
The go variable declaration in block can be grouped together for readablity 

## Example
```go
package main 

import "fmt"

func main(){
    var (
        x = "adams"
        y int = 30
        z string 
    )
}
```

#   Go Variable Naming Rules
Why go naming rule is very important is to avoid errors from occuring like naming your varible "func,
import ". A variable can have a short name (like x and y) 
Go variable naming rules:
-   A varible name must start with a letter or _
-   can not start a varible name with number
-   varible name are case sensitive
-   A variable can not contain spacing 
-   Variable can not take go keywords

#   Multi-Word Variable Names
When the name of the Variable is more than one word it makes it defficut to read for this reason 
the best way to handle this case is as follows.
-   Sanke Case: this can take and underscore (_) in each words "my_varible_name"
-   Pasca Case: every first word is to be capitalized "MyVaribleName"
-   Camel Case: every word can take a capital case but exception of the frist "myVaribleName"

