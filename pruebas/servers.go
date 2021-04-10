package main

import (
	"fmt"
	"net/http"
	"time"
)

// checkServer: revisa el estado de un servidor secuencialmente
func checkServer(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("servidor activo")
	}
}

func main() {
	now := time.Now()
	servers := []string{
		"http://google.com",
		"http://facebook.com",
	}

	for _, v := range servers {
		checkServer(v)
	}

	timeExecution := time.Since(now)
	fmt.Printf("tiempo de ejecución %s\n", timeExecution)
}
