package animation

import (
    "math"
    "unsafe"
)

const (
    CollisionTypeRect = 0
    CollisionTypeCircle = 1
)


func CircleIntersectsRect(pc, pr Vector, radius, width, height float64) bool {
    w := math.Abs(pc.X - pr.X)
    h := math.Abs(pc.Y - pr.Y)
    hw, hh := width / 2, height / 2

    if w >= radius + hw {
        return false
    }
    if h >= radius + hh {
        return false
    }

    w, h = w - hw, h - hh

    if w * w + h * h >= radius * radius {
        return false
    }

    return true
}


type Handler func(s ISprite)

type ISprite interface {
    CollisionType() int
    Position() Vector
    Width() float64
    Height() float64
    Radius() float64
    SetMoveVelocity(v Vector)
    SetPosition(v Vector)
    ClearIntersectors()
    CheckCollision(c ISprite)
    Disappear() bool
    SetDisappear(v bool)
    IsIntersects(c ISprite) bool
    Update(dt float64)
    AddEventHandler(name string, handler Handler)
    ClearEventHandlers(name string)
    TriggerEvent(name string)
    Pointer() unsafe.Pointer
}

type Sprite struct {
    intersectors []ISprite
    collisionType int
    position, moveVelocity Vector
    radius, width, height, rotation, rotateVelocity float64
    disappear bool
    events map[string][]Handler
}

func (s *Sprite) CollisionType() int {
    return s.collisionType
}

func (s *Sprite) Position() Vector {
    return s.position
}

func (s *Sprite) Width() float64 {
    return s.width
}

func (s *Sprite) Height() float64 {
    return s.height
}

func (s *Sprite) Radius() float64 {
    return s.radius
}

func (s *Sprite) SetMoveVelocity(v Vector) {
    s.moveVelocity = v
}

func (s *Sprite) SetPosition(v Vector) {
    s.position = v
}

func (s *Sprite) ClearIntersectors() {
    s.intersectors = make([]ISprite, len(s.intersectors))
}

func (s *Sprite) CheckCollision(c ISprite) {
    if s.Pointer() != c.Pointer() && s.IsIntersects(c) {
        s.intersectors = append(s.intersectors, c)
    }
}

func (s *Sprite) Pointer() unsafe.Pointer {
    return unsafe.Pointer(s)
}

func (s *Sprite) Disappear() bool {
    return s.disappear
}

func (s *Sprite) SetDisappear(v bool) {
    s.disappear = v
}

func (s *Sprite) IsIntersects(c ISprite) bool {
    if s.collisionType == CollisionTypeRect {
        if c.CollisionType() == CollisionTypeRect {
            return false
        }

        if c.CollisionType() == CollisionTypeCircle {
            return CircleIntersectsRect(
                c.Position(), s.position, c.Radius(), s.width, s.height)
        }
    }

    if s.collisionType == CollisionTypeCircle {
        if c.CollisionType() == CollisionTypeRect {
            return CircleIntersectsRect(
                s.position, c.Position(), s.radius, c.Width(), c.Height())
        }

        if c.CollisionType() == CollisionTypeCircle {
            return s.position.Distance(c.Position()) < s.radius + c.Radius()
        }
    }

    return false
}

func (s *Sprite) Update(dt float64) {
    s.position = s.position.Add(s.moveVelocity.Scale(dt))

    if s.rotateVelocity > 0 {
        s.rotation += s.rotateVelocity * dt
    }
}

func (s *Sprite) AddEventHandler(name string, handler Handler) {
    handlers, has := s.events[name]

    if !has {
        handlers = make([]Handler, 0)
    }

    s.events[name] = append(handlers, handler)
}

func (s *Sprite) ClearEventHandlers(name string) {
    s.events[name] = make([]Handler, 0)
}

func (s *Sprite) TriggerEvent(name string) {
    if handlers, has := s.events[name]; has && len(handlers) > 0 {
        for _, handler := range handlers {
            handler(s)
        }
    }
}
