package stat

import "sort"

type Set []float32

func (a Set) Len() int {
  return len(a)
}

func (a Set) Swap(i, j int) {
  a[i], a[j] = a[j], a[i]
}

func (a Set) Less(i, j int) bool {
  return a[i] < a[j]
}

func Mean(set []float32) float32 {
  var sum float32 = 0
  for i := 0; i < len(set); i++ {
    sum += set[i]
  }
  mean := sum/float32(len(set))
  return mean
}

func Median(set []float32) float32 {
  sort.Sort(Set(set))
  l := len(set)
  var median float32
  if l % 2 == 0 {
    i := set[l/2]
    j := set[l/2-1]
    median = Mean([]float32{i, j})
  } else {
    median = set[len(set)/2]
  }
  return median
}
