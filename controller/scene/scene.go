package scene


import (
    "low/app"
    "low/animation"
    "log"
)

var s *animation.Sprite

func Enter(m app.Message) {

    scene := new(animation.Scene)

    s = new(animation.Sprite)


    scene.AddSprite(s)

    go scene.Start()

    m.ReplySuccess(nil)
}

func Move(m app.Message) {
    x, _ := m.GetFloat("x")
    y, _ := m.GetFloat("y")

    log.Println(x)
    s.SetPosition(animation.Vector{X: x, Y: y,})
    m.ReplySuccess(nil)
}
