package main

import (
	. "projects/log_files_obfuscator/database"
	. "projects/log_files_obfuscator/routes"
)

func main() {
	path := "logs/data.json"
	Save(path)
	RunRoutes()
}
