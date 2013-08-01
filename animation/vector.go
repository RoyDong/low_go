package animation

import (
    "math"
)

type Vector struct {
    X, Y, Z float64
}

func (v Vector) Distance(t Vector) float64 {
    return math.Sqrt(math.Pow(t.X - v.X, 2) +
        math.Pow(t.Y - v.Y, 2) +
        math.Pow(t.Z - v.Z, 2))
}

func (v Vector) Add(t Vector) Vector {
    return Vector{
        X: v.X + t.X,
        Y: v.Y + t.Y,
        Z: v.Z + t.Z,
    }
}

func (v Vector) Minus(t Vector) Vector {
    return Vector{
        X: v.X - t.X,
        Y: v.Y - t.Y,
        Z: v.Z - t.Z,
    }
}

func (v Vector) Scale(s float64) Vector {
    return Vector{
        X: v.X * s,
        Y: v.Y * s,
        Z: v.Z * s,
    }
}
