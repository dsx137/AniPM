package main

import (
	"embed"

	"github.com/dsx137/anipm/anipm-server/cmd"
)

//go:embed all:anipm-client/dist/*
var clientFS embed.FS

func main() {
	cmd.Main(clientFS)
}
