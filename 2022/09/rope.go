package main

import "fmt"

type rope struct {
	Points []point
	Head   *point
	Tail   *point
}

func NewRope(length int) rope {
	pts := make([]point, length)
	return rope{
		Points: pts,
		Head:   &pts[0],
		Tail:   &pts[len(pts)-1],
	}
}

func (r *rope) Move(dir string) {
	switch dir {
	case "R":
		r.Head.X++
	case "L":
		r.Head.X--
	case "U":
		r.Head.Y++
	case "D":
		r.Head.Y--
	default:
		panic(fmt.Sprintf("invalid direction: '%s'", dir))
	}

	for i := 0; i < len(r.Points)-1; i++ {
		if r.Points[i].DistanceTo(r.Points[i+1]) > 1 {
			if r.Points[i].X-r.Points[i+1].X > 1 {
				r.Points[i+1].X++
				if r.Points[i].Y > r.Points[i+1].Y {
					r.Points[i+1].Y++
				} else if r.Points[i].Y < r.Points[i+1].Y {
					r.Points[i+1].Y--
				}
			} else if r.Points[i].X-r.Points[i+1].X < -1 {
				r.Points[i+1].X--
				if r.Points[i].Y > r.Points[i+1].Y {
					r.Points[i+1].Y++
				} else if r.Points[i].Y < r.Points[i+1].Y {
					r.Points[i+1].Y--
				}
			} else if r.Points[i].Y-r.Points[i+1].Y > 1 {
				r.Points[i+1].Y++
				if r.Points[i].X > r.Points[i+1].X {
					r.Points[i+1].X++
				} else if r.Points[i].X < r.Points[i+1].X {
					r.Points[i+1].X--
				}
			} else if r.Points[i].Y-r.Points[i+1].Y < -1 {
				r.Points[i+1].Y--
				if r.Points[i].X > r.Points[i+1].X {
					r.Points[i+1].X++
				} else if r.Points[i].X < r.Points[i+1].X {
					r.Points[i+1].X--
				}
			}
		}
	}
}
