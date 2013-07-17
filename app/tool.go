package app

import (
    "strings"
)


func ParseHttpQuery(q string) map[string]string {
    parts := strings.Split(string(q), "&")
    params := make(map[string]string, len(parts))

    for _, v := range parts {
        p := strings.Split(v, "=")

        if len(p) == 2 {
            params[p[0]] = p[1]
        }
    }

    return params
}
