package main

import (
	"github.com/michaelanony/go-spider/initSpider"
	"log"
)

func main() {
	if err := initSpider.Init(); err != nil {
		log.Fatal(err)
	}

}
