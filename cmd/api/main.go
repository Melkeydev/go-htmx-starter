package main

import (
	"fmt"
	appServer "fullGStack/internal/server"
)

func main() {

	server := appServer.NewServer()
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("%v", err)
		panic("cannot start server")
	}
}
