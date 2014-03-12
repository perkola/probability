package main

import (
  "fmt"
  "probability/statistics"
)

func main() {
  set1 := []float32{1, 2, 3, 4, 5, 6}
  fmt.Println(stat.Mean(set1))    // Gives 3.5
  fmt.Println(stat.Median(set1))  // Gives 3.5

  set2 := []float32{0.94, 1.20, 1.04, 1.01, 1.14}
  fmt.Println(stat.Mean(set2))    // Gives 1.066
  fmt.Println(stat.Median(set2))  // Gives 1.04
}
