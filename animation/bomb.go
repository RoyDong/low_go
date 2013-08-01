package animation

import (
    "math"
)

type Bomb struct {
    Sprite
    speed, timelimit, timeSinceFly float64
}


func (b *Bomb) SetSpeed(s float64) {
    b.speed = s
}

func (b *Bomb) SetTimelimit(t float64) {
    b.timelimit = t
}

func (b *Bomb) Fly(angle float64) {
    b.moveVelocity = Vector{X: b.speed * math.Sin(angle), Y: b.speed * math.Cos(angle)}
    b.timeSinceFly = 0
}

func (b *Bomb) Update(dt float64) {
    b.timeSinceFly += dt

    if b.timeSinceFly > b.timelimit {
        b.disappear = true
    }else{
        b.Sprite.Update(dt)
    }
}
