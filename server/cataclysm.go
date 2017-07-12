package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/miton18/cataclysm/cmd"
)

// Cataclysm aims to be the simpliest Torrent client with an API
// Usage:
//     cataclysm [flags]
//     cataclysm [command]

// Available Commands:
//     init        Initialise Cataclysm configuration
//     version     Print the version number
// -c, --config
// -v, --verbose
func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Panicf(err.Error())
	}
}
