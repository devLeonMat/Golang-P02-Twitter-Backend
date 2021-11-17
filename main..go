package main

import (
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/bd"
	"github.com/leonmatias2015/Golang-P02-Twitter-Backend/handlers"
	"log"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Without connection to database")
		return
	}
	handlers.Managers()
}
