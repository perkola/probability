package statistics

func Mean(set []float32) float32 {
  var sum float32 = 0
  for i := 0; i < len(set); i++ {
    sum += set[i]
  }
  mean := sum/float32(len(set))
  return mean
}
