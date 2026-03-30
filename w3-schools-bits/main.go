package main
import ("fmt")

func main() {
  myslice1 := make([]int, 10, 20)
  fmt.Println(myslice1)
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))


}