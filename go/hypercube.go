package main

import "github.com/fogleman/ln/ln"
//import "math"
import "fmt"

func CalcPath(p1 ln.Vector, p2 ln.Vector, p3 ln.Vector, p4 ln.Vector, t float64) ln.Vector {

  newVec := ln.Vector{0, 0, 0}

  if t >= 0 && t < 0.25 {
    newVec.X = p4.X + (p1.X - p4.X ) * t * 4;
    newVec.Y = p4.Y + (p1.Y - p4.Y ) * t * 4;
    newVec.Z = p4.Z + (p1.Z - p4.Z ) * t * 4;
  } else if t >= 0.25 && t < 0.5 {
    newVec.X = p1.X + (p2.X - p1.X ) * (t - 0.25) * 4;
    newVec.Y = p1.Y + (p2.Y - p1.Y ) * (t - 0.25) * 4;
    newVec.Z = p1.Z + (p2.Z - p1.Z ) * (t - 0.25) * 4;
  } else if t >= 0.5 && t < 0.75 {
    newVec.X = p2.X + (p3.X - p2.X ) * (t - 0.5) * 4;
    newVec.Y = p2.Y + (p3.Y - p2.Y ) * (t - 0.5) * 4;
    newVec.Z = p2.Z + (p3.Z - p2.Z ) * (t - 0.5) * 4;
  } else if t >= 0.75 && t < 1.0 {
    newVec.X = p3.X + (p4.X - p3.X ) * (t - 0.75) * 4;
    newVec.Y = p3.Y + (p4.Y - p3.Y ) * (t - 0.75) * 4;
    newVec.Z = p3.Z + (p4.Z - p3.Z ) * (t - 0.75) * 4;
  }

  return newVec;
}

type Hypercube struct {
        Vertices [16]ln.Vector
	Box ln.Box
}

func NewHypercube(t float64) *Hypercube {
        min, max := ln.Vector{-1,-1,-1}, ln.Vector{ 1, 1, 1}
        
        verticesStart := [16]ln.Vector{
                ln.Vector{-0.5, -0.5, -0.5},
                ln.Vector{ 0.5, -0.5, -0.5},
                ln.Vector{ 0.5, 0.5, -0.5},
                ln.Vector{-0.5, 0.5, -0.5},
                
                ln.Vector{-0.5, -0.5, 0.5},
                ln.Vector{ 0.5, -0.5, 0.5},
                ln.Vector{ 0.5, 0.5, 0.5},
                ln.Vector{-0.5, 0.5, 0.5},
                
                ln.Vector{-1, -1, -1},
                ln.Vector{ 1, -1, -1},
                ln.Vector{ 1, 1, -1},
                ln.Vector{-1, 1, -1},
                
                ln.Vector{-1, -1, 1},
                ln.Vector{ 1, -1, 1},
                ln.Vector{ 1, 1, 1},
                ln.Vector{-1, 1, 1},
        }
        
        verticesNow := [16]ln.Vector{
                CalcPath(verticesStart[8], verticesStart[9], verticesStart[1], verticesStart[0], t),
                CalcPath(verticesStart[0], verticesStart[8], verticesStart[9], verticesStart[1], t),
                CalcPath(verticesStart[3], verticesStart[11], verticesStart[10], verticesStart[2], t),
                CalcPath(verticesStart[11], verticesStart[10], verticesStart[2], verticesStart[3], t),
                CalcPath(verticesStart[12], verticesStart[13], verticesStart[5], verticesStart[4], t),
                CalcPath(verticesStart[4], verticesStart[12], verticesStart[13], verticesStart[5], t),
                CalcPath(verticesStart[7], verticesStart[15], verticesStart[14], verticesStart[6], t),
                CalcPath(verticesStart[15], verticesStart[14], verticesStart[6], verticesStart[7], t),
                CalcPath(verticesStart[9], verticesStart[1], verticesStart[0], verticesStart[8], t),
                CalcPath(verticesStart[1], verticesStart[0], verticesStart[8], verticesStart[9], t),
                CalcPath(verticesStart[2], verticesStart[3], verticesStart[11], verticesStart[10], t),
                CalcPath(verticesStart[10], verticesStart[2], verticesStart[3], verticesStart[11], t),
                CalcPath(verticesStart[13], verticesStart[5], verticesStart[4], verticesStart[12], t),
                CalcPath(verticesStart[5], verticesStart[4], verticesStart[12], verticesStart[13], t),
                CalcPath(verticesStart[6], verticesStart[7], verticesStart[15], verticesStart[14], t),
                CalcPath(verticesStart[14], verticesStart[6], verticesStart[7], verticesStart[15], t),
        }
	
        box := ln.Box{min, max}
	return &Hypercube{verticesNow, box}
}

func (c *Hypercube) Compile() {
}

func (c *Hypercube) BoundingBox() ln.Box {
	return c.Box
}

func (c *Hypercube) Contains(v ln.Vector, f float64) bool {
        min, max := ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}
	if v.X < min.X-f || v.X > max.X+f {
		return false
	}
	if v.Y < min.Y-f || v.Y > max.Y+f {
		return false
	}
	if v.Z < min.Z-f || v.Z > max.Z+f {
		return false
	}
	return true
}

func (c *Hypercube) Intersect(r ln.Ray) ln.Hit {
        /*min, max := c.Vertices[0], c.Vertices[6]
	n := min.Sub(r.Origin).Div(r.Direction)
	f := max.Sub(r.Origin).Div(r.Direction)
	n, f = n.Min(f), n.Max(f)
	t0 := math.Max(math.Max(n.X, n.Y), n.Z)
	t1 := math.Min(math.Min(f.X, f.Y), f.Z)
	if t0 < 1e-3 && t1 > 1e-3 {
		return ln.Hit{c, t1}
	}
	if t0 >= 1e-3 && t0 < t1 {
		return ln.Hit{c, t0}
	}*/
	return ln.NoHit
}

func (c *Hypercube) Paths() ln.Paths {
        
        vertexJoins := [32][2]int{
                {0, 1}, {1, 2}, {2, 3}, {3, 0},
                {0, 4}, {1, 5}, {2, 6}, {3, 7},
                {4, 5}, {5, 6}, {6, 7}, {7, 4},
 
                {0, 8}, {1, 9}, {2, 10}, {3, 11},
                {4, 12}, {5, 13}, {6, 14}, {7, 15},

                {8, 9}, {9, 10}, {10, 11}, {11, 8},
                {8, 12}, {9, 13}, {10, 14}, {11, 15},
                {12, 13}, {13, 14}, {14, 15}, {15, 12},
        }
        
        var paths ln.Paths
        
        for i := 0; i < 32; i++ {
                paths = append(paths, ln.Path{c.Vertices[vertexJoins[i][0]], c.Vertices[vertexJoins[i][1]]})
        }
        
	return paths
}

func main() {
        
        eye := ln.Vector{-4, 3, 2}        // camera position
        center := ln.Vector{0, 0, 0}      // camera looks at
        up := ln.Vector{0, 0, 1}          // up direction
        width := 1024.0                   // rendered width
        height := 1024.0                  // rendered height
        fovy := 50.0                      // vertical field of view, degrees
        znear := 0.1                      // near z plane
        zfar := 10.0                      // far z plane
        step := 0.01                      // how finely to chop the paths for visibility testing
        
        var scene ln.Scene
        var paths ln.Paths
        var t float64
        
        for i := 0; i < 125; i++ {
                t = float64(i) / 125.0
                
                scene = ln.Scene{}
                scene.Add(NewHypercube(t))
                paths = scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)

        	paths.WriteToSVG(fmt.Sprintf("hypercube_%03d.svg", i), width, height)
                
        }
}
