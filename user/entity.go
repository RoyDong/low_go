package user

import (
    "low/app"
    "encoding/json"
    "strconv"
)

type Entity struct {
    id int64
    name, email, passwd, salt string
    created_at int64
}

/*

setters

*/
func (user *Entity) SetName(name string) *Entity {
    user.name = name
    return user
}

func (user *Entity) SetEmail(email string) *Entity {
    user.email = email
    return user
}

func (user *Entity) SetPasswd(passwd string) *Entity {
    user.salt = app.RandString(SaltLen)
    user.passwd = HashPasswd(passwd, user.salt)
    return user
}

func (user *Entity) SetCreatedAt(sec int64) *Entity {
    user.created_at = sec
    return user
}

/*

getters

*/
func (user *Entity) Id() int64 {
    return user.id
}

func (user *Entity) Name() string {
    return user.name
}

func (user *Entity) Email() string {
    return user.email
}


type DataForJson struct {
    Id int64 `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
    CreatedAt int64 `json:"created_at"`
}

func (user *Entity) Json() string {
    stream, _ := json.Marshal(&DataForJson{
        Id: user.id,
        Name: user.name,
        Email: user.email,
        CreatedAt: user.created_at,
    })
    return string(stream)
}

func (user *Entity) CheckPasswd(passwd string) bool {
    return HashPasswd(passwd, user.salt) == user.passwd
}

func (user *Entity) dataForMysql() map[string]string {
    return map[string]string{
        "name": user.name,
        "email": user.email,
        "passwd": user.passwd,
        "salt": user.salt,
        "created_at": strconv.FormatInt(user.created_at, 10),
    }
}
