package main

import "fmt"

type serverState int

const (
	idle serverState = iota
	connected
	error
	retry
)

var serverStateMap = map[serverState]string{
	idle:      "Idle",
	connected: "Connected",
	error:     "Error",
	retry:     "Retry",
}

func (estadoServidor serverState) String() string {
	return serverStateMap[estadoServidor]
}

func verificadorDeRed(servidor serverState) serverState {
	switch servidor {
	case idle:
		return connected
	case connected, retry:
		return idle
	case error:
		return error
	default:
		panic(fmt.Errorf("Estado del servidor desconocido: %v", servidor))

	}
}

func main() {
	redServidor := verificadorDeRed(idle)
	fmt.Println(redServidor)

	segundaRed := verificadorDeRed(connected)
	fmt.Println(segundaRed)

	terceraRed := verificadorDeRed(error)
	fmt.Println(terceraRed)
}
