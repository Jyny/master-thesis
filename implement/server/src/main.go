package main

import (
	"os"
	"server/pkg/routes"
	"server/pkg/sql"
	"server/pkg/worker"
)

func main() {
	gorm := sql.GromInit(os.Getenv("DB"))
	worker := worker.New(gorm)
	worker.Start()
	routes.Run("localhost:8080", gorm, worker)
}
