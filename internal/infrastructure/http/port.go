package http

import (
	"fmt"
	"net"
	"os"
)

func GetAvailablePort(basePort int) (int, net.Listener) {
	var ln net.Listener
	var err error

	for i := 0; i < 3; i++ {
		tryPort := basePort + i
		ln, err = net.Listen("tcp", fmt.Sprintf(":%d", tryPort))
		if err == nil {
			return tryPort, ln
		}
	}

	fmt.Println("Error: All three ports are in use. The app cannot be started.")
	os.Exit(1)
	return 0, nil
}
