package main

import (
	appServer "fullGStack/internal/server"
)

func main() {

	server := appServer.NewServer()
	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
