package main

import (
	"backend/internal/app"
	"backend/internal/ports/httpgin"
)

func main() {
	usecase := app.NewApp(repo, usrepo)
	server := httpgin.NewHTTPServer(":18080", usecase)
	err := server.Listen()
	if err != nil {
		panic(err)
	}
}
