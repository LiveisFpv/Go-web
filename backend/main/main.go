package main

import (
	"homework8/internal/adrepo"
	"homework8/internal/app"
	"homework8/internal/ports/httpgin"
	"homework8/internal/usrepo"
)

func main() {
	repo := adrepo.New()
	usrepo := usrepo.New()
	usecase := app.NewApp(repo, usrepo)
	server := httpgin.NewHTTPServer(":18080", usecase)
	err := server.Listen()
	if err != nil {
		panic(err)
	}
}
