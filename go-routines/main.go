package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://x.com",
	}

	//สร้างแชแนลสำหรับสื่อสารระหว่าง main routine กับ child routine
	c := make(chan string)

	for _, link := range links {
		//call checklink in go routine (multiverse)
		go checkLink(link, c)
	}

	for l := range c {
		//ใช้ go routine + function literals เพื่อสร้าง anonymous ฟังก์ชั่นขึ้นมาให้สามารถใช้ time.sleep ได้ ก่อน checkLink โดยไม่ไป sleep กระทบ main routine
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(link, " is down")
	} else {
		fmt.Println(link, " is up")
		//return message to channel
		c <- link
	}
}
