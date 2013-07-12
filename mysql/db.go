package mysql

import (
    "log"
    "low/app"
    "strings"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DB() *sql.DB {
    if db == nil {
        config := app.Config.Mysql
        db, _ = sql.Open("mysql", config.Dsn + "/" + config.Dbname)
    }

    return db
}

func Insert(table string, data map[string]string) int64 {
    i, l := 0, len(data)
    keys, values := make([]string, l), make([]string, l)

    for k, v := range data {
        keys[i] = "`" + k + "`"
        values[i] = "'" + v + "'"
        i++
    }

    sql := "insert into `" + table + "` (" +
            strings.Join(keys, ",") + ") values (" +
            strings.Join(values, ",") + ")"

    result, e := DB().Exec(sql)

    if e != nil {
        log.Fatal(e)
    }

    id, e := result.LastInsertId()

    if e!= nil {
        log.Fatal(e)
    }

    return id
}

func Update(table string, data map[string]string, condition string) bool {
    i, parts := 0, make([]string, len(data))

    for k, v := range data {
        parts[i]= "`" + k + "`='" + v + "'"
        i++
    }

    sql := "update `" + table + "` set " + strings.Join(parts, ",") +
            " where " + condition

    _, e:= DB().Exec(sql)

    if e != nil {
        log.Fatal(e)
    }

    return true
}

func Delete(table string, condition string) bool {
    _, e := DB().Exec("delete from `" + table + "` where " + condition)

    if e != nil {
        log.Fatal(e)
    }

    return true
}
