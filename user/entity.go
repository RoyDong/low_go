package user

import (
    "low/app"
    "encoding/json"
    "strconv"
)

type User struct {
    id int64
    name, email, passwd, salt string
    created_at int64
}

/*

setters

*/
func (user *User) SetName(name string) *User {
    user.name = name
    return user
}

func (user *User) SetEmail(email string) *User {
    user.email = email
    return user
}

func (user *User) SetPasswd(passwd string) *User {
    user.salt = app.RandString(SaltLen)
    user.passwd = HashPasswd(passwd, user.salt)
    return user
}

func (user *User) SetCreatedAt(sec int64) *User {
    user.created_at = sec
    return user
}

/*

getters

*/
func (user *User) Id() int64 {
    return user.id
}

func (user *User) Name() string {
    return user.name
}

func (user *User) Email() string {
    return user.email
}


type DataForJson struct {
    Id int64 `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
    CreatedAt int64 `json:"created_at"`
}

func (user *User) Data() *DataForJson {
    return &DataForJson{
        Id: user.id,
        Name: user.name,
        Email: user.email,
        CreatedAt: user.created_at,
    }
}

func (user *User) Json() []byte {
    stream, _ := json.Marshal(&DataForJson{
        Id: user.id,
        Name: user.name,
        Email: user.email,
        CreatedAt: user.created_at,
    })
    return stream
}

func (user *User) CheckPasswd(passwd string) bool {
    return HashPasswd(passwd, user.salt) == user.passwd
}

func (user *User) dataForMysql() map[string]string {
    return map[string]string{
        "name": user.name,
        "email": user.email,
        "passwd": user.passwd,
        "salt": user.salt,
        "created_at": strconv.FormatInt(user.created_at, 10),
    }
}
