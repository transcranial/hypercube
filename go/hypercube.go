package main

import "github.com/fogleman/ln/ln"

type Hypercube struct {
        Min ln.Vector
        Max ln.Vector
}

func NewHypercube(min, max ln.Vector) *Hypercube {
	return &Hypercube{min, max}
}

func (c *Hypercube) Paths() ln.Paths {
	x1, y1, z1 := c.Min.X, c.Min.Y, c.Min.Z
	x2, y2, z2 := c.Max.X, c.Max.Y, c.Max.Z
	paths := ln.Paths{
		{{x1, y1, z1}, {x1, y1, z2}},
		{{x1, y1, z1}, {x1, y2, z1}},
		{{x1, y1, z1}, {x2, y1, z1}},
		{{x1, y1, z2}, {x1, y2, z2}},
		{{x1, y1, z2}, {x2, y1, z2}},
		{{x1, y2, z1}, {x1, y2, z2}},
		{{x1, y2, z1}, {x2, y2, z1}},
		{{x1, y2, z2}, {x2, y2, z2}},
		{{x2, y1, z1}, {x2, y1, z2}},
		{{x2, y1, z1}, {x2, y2, z1}},
		{{x2, y1, z2}, {x2, y2, z2}},
		{{x2, y2, z1}, {x2, y2, z2}},
	}
	return paths
	paths = paths[:0]
	for i := 0; i <= 10; i++ {
		p := float64(i) / 10
		var x, y float64
		x = x1 + (x2-x1)*p
		y = y1 + (y2-y1)*p
		paths = append(paths, ln.Path{{x, y1, z1}, {x, y1, z2}})
		paths = append(paths, ln.Path{{x, y2, z1}, {x, y2, z2}})
		paths = append(paths, ln.Path{{x1, y, z1}, {x1, y, z2}})
		paths = append(paths, ln.Path{{x2, y, z1}, {x2, y, z2}})
	}
	return paths
}

func main() {
        scene := ln.Scene{}
        scene.Add(ln.NewCube(ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}))

        eye := ln.Vector{4, 3, 2}         // camera position
        center := ln.Vector{0, 0, 0}      // camera looks at
        up := ln.Vector{0, 0, 1}          // up direction
        width := 1024.0                   // rendered width
        height := 1024.0                  // rendered height
        fovy := 50.0                      // vertical field of view, degrees
        znear := 0.1                      // near z plane
        zfar := 10.0                      // far z plane
        step := 0.01                      // how finely to chop the paths for visibility testing

        paths := scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)

	paths.WriteToPNG("hypercube.png", width, height)
	//paths.WriteToSVG("hypercube.svg", width, height)
}
