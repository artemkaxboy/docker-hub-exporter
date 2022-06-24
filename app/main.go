/*
 * Copyright (c) 2022 Artem Kolin (https://github.com/artemkaxboy)
 */

package main

import (
	"docker-hub-exporter/cmd"
	"github.com/jessevdk/go-flags"
	"os"

	log "github.com/go-pkgz/lgr"
)

// Opts with all cli commands and flags
type Opts struct {
	ServerCmd cmd.ServerCommand `command:"server"`

	Dbg bool `long:"dbg" env:"DEBUG" description:"debug mode"`
}

func main() {

	var opts Opts
	p := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash)
	p.CommandHandler = func(command flags.Commander, args []string) error {
		setupLog(opts.Dbg)
		c := command.(flags.Commander)

		err := c.Execute(args)
		if err != nil {
			log.Printf("[ERROR] failed with %+v", err)
		}
		return err
	}

	if _, err := p.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}

func setupLog(dbg bool) {
	if dbg {
		log.Setup(log.Debug, log.CallerFile, log.CallerFunc, log.Msec, log.LevelBraces)
		log.Printf("[DEBUG] debug mode enabled")
		return
	}
	log.Setup(log.Msec, log.LevelBraces)
}
