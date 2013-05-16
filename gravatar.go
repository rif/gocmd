package main

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strconv"
)

func Gravatar(email, def string, size int) string {
	m := md5.New()
	m.Write([]byte(email))
	bs := m.Sum(nil)
	gravatar_url := "http://www.gravatar.com/avatar/" + fmt.Sprintf("%x", bs)
	u, _ := url.Parse(gravatar_url)
	q := u.Query()
	q.Set("d", def)
	q.Set("s", strconv.Itoa(size))
	u.RawQuery = q.Encode()
	return u.String()
}

func main() {
	fmt.Println(Gravatar("someone@somewhere.com", "http://www.example.com/default.jpg", 40))
}
