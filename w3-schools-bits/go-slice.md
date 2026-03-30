#   Go Slices
Slices are similar to arr.slices are also used to store grouped values of the same type in a single variable.Unlike arrays, the length of a slice can be increased or decreased as you want it to be.
In Go, there are several ways to create a slice:

    Using the []datatype{values} format
    Create a slice from an array
    Using the make() function

# Example of go slice
```go
package main
import ("fmt")

func main() {

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)
} 

```
#   Create a Slice From an Array
You can create a slice by slicing an array
```go
package main 

import "fmt"

func main(){
    var array = [8]int{1,2,3,4,5,6,7,8}
    var slice = array[:5]

    fmt.Println(slice)
}
```
#   Create a Slice With The make() Function
```go
package main
import ("fmt")

func main() {
  myslice1 := make([]int, 10, 20)
  fmt.Println(myslice1)
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))


}
```
#   Access Elements of a Slice
we can always access a specific slice element by referring to the index number. The example below 
shows how to access the first and the third element in the slice [start:end] the first index is 0.
# Example of how to access elements in slice 
```go 
package main
import ("fmt")

func main() {
  num_student := []int{3,20,30}

  fmt.Println(prices[0])
  fmt.Println(prices[3])
}
```
#   Change Elements of a Slice
You can also change a specific slice element by referring to the index number
## Example of how to change an element in a slice
```go 
package main 

import "fmt"

func main (){
    slice_name := []int{1,3,4,5,6}
    slice_name[1]  = 50

    fmt.Println(slice_name)
    fmt.Println(slice_name[1])
}

```
#   Append Elements To a Slice
this a way of adding element to your slices using the apend function append()
#   Example of append
```go
package main 

import "fmt"

func main (){
    myslice := []int{10,12,13,14,15}
    myslice = append(myslice, 20,30,40)

    fmt.Println(myslice)
    fmt.Println(len(myslice))
    fmt.Println(cap(myslice))
}

```
#   Append One Slice To Another Slice
we append all the elements of one slice to another slice, using the append()function:
# Example of appending one sice to another slice
```go
package main

import "fmt"

func main (){
    myslice := []int{10,11,12,13,14}
    myslice01 := []int{21,30,40,50}

    slice = append(myslice, myslice01...)

    fmt.Println(myslice)
    fmt.Println(myslice01)
     fmt.Println(slice)
}
```
