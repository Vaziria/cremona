package main

import (
	"changeme/backend"
	"changeme/backend/apis"
	"log"
)

func main() {
	log.Println("test ")

	repo := backend.AccountRepo{}
	repo.Initialize()

	akun, _ := repo.Find("bima")

	api := apis.NewChatApi(akun)

	api.UserInit()
}
