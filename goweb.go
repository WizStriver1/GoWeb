package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"os"
	"runtime"

	"app/common/database"
	"app/common/jsonconfig"
	"app/common/server"
	"app/common/session"
	"app/route"
)

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// Load the configuration file
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)

	// Configure the session cookie store
	session.Configure(config.Session)
	// Connect to database
	database.Connect(config.Database)

	// Start the listener
	server.Run(route.LoadHTTP(), route.LoadHTTPS(), config.Server)
}

var config = &configuration{}

// configuration contains default setting
type configuration struct {
	Database database.Info   `json:"Database"`
	Server   server.Server   `json:"Server"`
	Session  session.Session `json:"Session"`
}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
