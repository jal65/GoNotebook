package main
import . "fmt"

const shades = " .,o#&0%"

func main() {
  for y := -1.0; y < 1; y += 0.03 {
    var row []byte
    for x := -1.0; x < 1; x += 0.015 {
      if insideCircle(x, y) {
        row = append(row, shades[shadePixel(x, y)])
      } else {
        row = append(row, shades[0])
      }
    }
    Println(string(row))
  }
}

func insideCircle(x, y float64) bool {
  return x * x + y * y < 1
}

func distanceFromEdge(x, y float64) float64 {
  return 1 - x * x - y * y
}

func shadePixel(x, y float64) int {
  return int(4 * (0.3 * x + 0.6 * (y + distanceFromEdge(x, y)) + 1))
}