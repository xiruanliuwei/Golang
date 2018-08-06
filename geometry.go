
package main

import "fmt"
import "math"

type Point struct {
    X, Y float64
}

func Distance(p, q Point) float64 {
    fmt.Printf("function Distance is called\n")
    return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p Point) Distance(q Point) float64 {
    fmt.Printf("method Distance is called\n")
    return math.Hypot(q.X - p.X, q.Y - p.Y)
}

type Path []Point  // slice

func (path Path) Distance() float64 {
    sum := 0.0
    for i, val := range path {
        if i > 0 {
            sum += path[i-1].Distance(path[i])
            fmt.Println(val)	
        }
    }

    return sum
}

func main() {

    var v Point = Point{1, 2}  // struct literal
    fmt.Println(v)

    p := Point{1, 2}  // struct literal
    q := Point{4, 6}

    fmt.Println(Distance(p, q))
    fmt.Println(p.Distance(q))

    perim := Path{  // slice literal
    	{1, 1},
    	{5, 1},
    	{5, 4},
    	{1, 1},
    }

    fmt.Println(perim.Distance())
}
