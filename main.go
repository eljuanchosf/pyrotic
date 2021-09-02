package main

import (
	"log"
	"os"

	"github.com/code-gorilla-au/pyrotic/internal/chalk"
	"github.com/code-gorilla-au/pyrotic/internal/commands"
)

func main() {
	setLogger()
	if err := commands.Execute(); err != nil {
		os.Exit(1)
	}
}

func setLogger() {
	log.SetPrefix(chalk.Cyan("pyrotic "))

	if val, ok := os.LookupEnv("ENV"); ok {
		if val == "DEV" {
			log.Println(chalk.Cyan("DEVELOPMENT MODE"))
			log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.Lmsgprefix)
			return
		}
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lmsgprefix)

}
