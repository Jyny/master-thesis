package main

import (
	"os"
	"server/pkg/routes"
	"server/pkg/sql"
)

func main() {
	gorm := sql.GromInit(os.Getenv("DB"))
	routes.Run(gorm, "localhost:8080")
}
