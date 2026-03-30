#   Go Constants
Go constants is simply declearing or creating a varible that can not be changed 
the value of the varible is fixed and cant be change.

#   Example of a constant varible
```go
package main 

import "fmt"

func main(){
    const Sum = 30

    const Studen_Name string = "Dickson"
    //Constants can be declared both inside and outside of a function
}

```
#   Multiple Constants Declaration
Multiple constants can be grouped together into a block this makes it easy to read and understand:

```go
package main
import ("fmt")

const (
  X int = 20
  Y = 4.0
  Z = "hellow"
)

func main() {

}

```