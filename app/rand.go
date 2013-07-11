package app

import ( 
    "math/rand"
)

var chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func RandString(length int) string {
    rndChars := make([]byte, length)
    max := len(chars) - 1

    for i := 0; i < length; i++ {
        rndChars[i] = chars[rand.Intn(max)]
    }

    return string(rndChars)
}

func RandInt(min int, max int) int {
    return min + rand.Intn(max - min)
}
