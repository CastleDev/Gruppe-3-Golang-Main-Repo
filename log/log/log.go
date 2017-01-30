package log

import "math"


func Log2(x float64) int  {

answer := math.Log2(x)
var intAnswer int = int(answer)

return intAnswer

}

func Log10(x float64) int {

  answer := math.Log10(x)
  var intAnswer int = int(answer)

  return intAnswer
}
