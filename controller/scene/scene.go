package scene


import (
    "low/app"
    "low/animation"
    "strconv"
    "log"
)

var s *animation.Sprite

func Enter(m app.Message) {

    scene := animation.CreateScene()

    s = new(animation.Sprite)


    scene.AddSprite(s)

    go scene.Start()

    m.ReplySuccess(nil)
}

func Move(m app.Message) {
    x, _ := m.Get("x")
    y, _ := m.Get("y")

    fx, _ := strconv.ParseFloat(x, 64)
    fy, _ := strconv.ParseFloat(y, 64)
    log.Println(x, fx) 
    s.SetPosition(animation.Vector{X: fx, Y: fy,})
    m.ReplySuccess(nil)
}
