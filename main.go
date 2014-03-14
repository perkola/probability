package main

import (
  "fmt"
  "probability/statistics"
)

func main() {
  set1 := []float32{1, 1, 3, 4, 5, 6}
  fmt.Println(stat.Mean(set1))    // Gives 3.5
  fmt.Println(stat.Median(set1))  // Gives 3.5
  fmt.Println(stat.Mode(set1))    // Gives [1]

  set2 := []float32{1.20, 0.94, 1.14, 1.04, 1.01, 1.14, 1.20}
  fmt.Println(stat.Mean(set2))    // Gives 1.066
  fmt.Println(stat.Median(set2))  // Gives 1.04
  fmt.Println(stat.Mode(set2))    // Gives [1.14 1.2]

  set3 := []float32{8, 7, 6, 5}
  fmt.Println(stat.Mode(set3))    // Gives []
}
