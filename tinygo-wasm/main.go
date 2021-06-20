package main

import (
	"log"
	"net/http"
	"strings"
)

const dir = "./html"

func main() {
	fs := http.FileServer(http.Dir(dir))
	log.Print("Serving " + dir + " on http://localhost:8080")

}
