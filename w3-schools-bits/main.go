package main
import ("fmt")

func main() {
  sd := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  // Original slice

  // Create copy with only needed numbers
  s := sd[:len(sd)-0]
  y := make([]int, len(s))
  copy(y, s)

  fmt.Printf("numbersCopy = %v\n", y)
  
}