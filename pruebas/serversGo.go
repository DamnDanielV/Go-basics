package main

import (
	"fmt"
	"net/http"
	"time"
)

// checkServer: revisa si un servidor está activo
// server: servidor a revisar
// ch: canal por el que se devolvera la data
func checkServer(server string, ch chan string) {
	_, err := http.Get(server)
	if err != nil {
		ch <- err.Error() //3. envia el mensaje de error por el canal
	} else {
		ch <- server + " activo\n" // 3. envia el mensaje por el canal
	}
}

func main() {
	now := time.Now()
	channel := make(chan string) // 1. inicializa un objeto de tipo chan string
	servers := []string{
		"http://google.com",
		"http://facebook.com",
	}

	for _, v := range servers {
		go checkServer(v, channel) // 2. indica una ejecucion asincrona
	}

	for i := 0; i < len(servers); i++ {
		// ciclo para esperar a leer toda la data de los canales
		fmt.Printf(<-channel)
	}

	timeExecution := time.Since(now)
	fmt.Printf("tiempo de ejecución %s\n", timeExecution)
}
