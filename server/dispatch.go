package server

import (
    "low/controller/user"
    "low/controller/scene"
)

func (m *Message)Dispatch() {
    path := m.TitlePath()

    switch path[0] {
    case "user":

        switch path[1] {
        case "show":user.Show(m)
        case "signin":user.Signin(m)

        }

    case "scene":
        switch path[1] {
        case "enter":scene.Enter(m)
        case "move":scene.Move(m)

        }

    }
}
