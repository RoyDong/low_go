package animation

import (
    "low/app"
)

type Action interface {
    Sprite() ISprite
    Params() interface{}
}

type Action struct {
    sprite ISprite
    params interface{}
}
