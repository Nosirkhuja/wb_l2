package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

func main() {
	resp, err := ntp.Query("ntp1.stratum2.ru")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Time.Local())

	now, err := ntp.Time("ntp2.stratum2.ru")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(now)
}
