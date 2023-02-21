package tests

import (
	"log"
	"time"
)

func One() {
	log.Println("ONE start")
	time.Sleep(15 * time.Second)
	log.Println("ONE finish")
}

func Two() {
	log.Println("TWO start")
	time.Sleep(10 * time.Second)
	log.Println("TWO finish")
}

func Three(text string) {
	log.Println("THREE -", text, "start")
	time.Sleep(15 * time.Second)
	log.Println("THREE -", text, "finish")
}
