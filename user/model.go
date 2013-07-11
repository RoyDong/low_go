package user

import (
    "encoding/hex"
    "log"
    "fmt"
    "crypto/sha512"
    "time"
    "low/mysql"
)

const (
    Table = "user"
    SaltLen = 32
)

var entities = make(map[int64]*Entity)

func New() *Entity {
    return new(Entity).SetCreatedAt(time.Now().Unix())
}

func Find(id int64) (*Entity, bool) {
    user, has := entities[id]

    if has {
        return user, true
    }

    user = new(Entity)
    sql := fmt.Sprintf("select id,name,email,passwd,salt,created_at from `"+
            Table +"` where `id`=%d", id)
    e := mysql.DB().QueryRow(sql).Scan(&user.id, &user.name, &user.email,
            &user.passwd, &user.salt, &user.created_at)

    if e != nil {
        log.Fatal(e)
        return user, false
    }

    entities[user.id] = user
    return user, true
}

func Save(user *Entity) bool {
    data := user.dataForMysql()

    if user.id > 0 {
        if _, has := entities[user.id]; !has {
            entities[user.id] = user
        }
        return mysql.Update(Table, data, fmt.Sprintf("`id`=%d", user.id))
    }

    user.id = mysql.Insert(Table, data)

    if user.id == 0 {
        return false
    }

    entities[user.id] = user
    return true
}

func Delete(id int64) bool {
    if !mysql.Delete(Table, fmt.Sprintf("`id`=%d", id)){
        return false
    }

    delete(entities, id)
    return true
}

func TotalCount() int {
    return len(entities)
}

func HashPasswd(passwd string, salt string) string {
    hash := sha512.New()

    if _, e := hash.Write([]byte(passwd + salt)); e != nil {
        log.Fatal(e)
        return ""
    }

    return hex.EncodeToString(hash.Sum(nil))
}
