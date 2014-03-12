package main

import (
  "fmt"
  "probability/statistics"
)

func main() {
  set1 := []float32{1, 2, 3, 4, 5}
  fmt.Println(statistics.Mean(set1))  // Gives 3

  set2 := []float32{0.94, 1.20, 1.04, 1.01, 1.14}
  fmt.Println(statistics.Mean(set2))  // Gives 1.066
}
