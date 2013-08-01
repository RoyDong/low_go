package animation

import (
    "time"
    "log"
)

const (
    FrameInterval = 500
)

type Scene struct {
    lastFrameTime float64
    sprites []ISprite
    added []ISprite
}

func CreateScene() *Scene {
    return &Scene{
        sprites: make([]ISprite, 0),
        added: make([]ISprite, 0),
    }
}

func (s *Scene) AddSprite(sprite ISprite) {
    sprite.SetDisappear(false)
    s.added = append(s.added, sprite)
}

func (s *Scene) perform(dt float64) {
    var sprite ISprite
    var i,j int
    l := len(s.sprites)

    for i = 0; i < l; i++ {
        sprite = s.sprites[i]
        sprite.ClearIntersectors()

        for j = i + 1; j < l; j++ {
            sprite.CheckCollision(s.sprites[j])
        }
    }

    for _, sprite = range s.sprites {
        sprite.Update(dt)
    }

    sprites := make([]ISprite, 0, len(s.sprites) + len(s.added))

    for _, sprite = range s.sprites {
        if !sprite.Disappear() {
            sprites = append(sprites, sprite)
        }
    }


    for _, sprite = range s.added {
        sprites = append(sprites, sprite)
    }

    s.added = make([]ISprite, 0, len(s.added))
    s.sprites = sprites
}

func (s *Scene) Start() {
    s.lastFrameTime = float64(time.Now().UnixNano()) / 1000000000

    for t := range time.Tick(FrameInterval * time.Millisecond) {
        ft := float64(t.UnixNano()) / 1000000000
        s.perform(ft - s.lastFrameTime)
        s.lastFrameTime = ft
    }

    log.Println("Scene close")
}
