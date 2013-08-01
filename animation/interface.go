package animation

import (
    "unsafe"
)

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
