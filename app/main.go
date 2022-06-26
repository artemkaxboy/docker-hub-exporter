/*
 * Copyright (c) 2022 Artem Kolin (https://github.com/artemkaxboy)
 */

package main

import (
	"docker-hub-exporter/cmd"
	"github.com/jessevdk/go-flags"
	"os"
	"strings"

	log "github.com/go-pkgz/lgr"
)

// Opts with all cli commands and flags
type Opts struct {
	ServerCmd cmd.ServerCommand `command:"server"`
}

func main() {

	var opts Opts
	p := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash|flags.IgnoreUnknown)
	p.CommandHandler = func(command flags.Commander, args []string) error {
		setupLog(true)
		c := command.(flags.Commander)

		err := c.Execute(args)
		if err != nil {
			log.Printf("[ERROR] failed with %+v", err)
		}
		return err
	}

	if _, err := p.ParseArgs(addServerCommandToArgs(addExtraDashForLegacyOptions(os.Args[1:]))); err != nil {
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

func addServerCommandToArgs(args []string) []string {
	return append([]string{"server"}, args...)
}

func addExtraDashForLegacyOptions(args []string) []string {

	dashedArgs := make([]string, len(args))
	for i, arg := range args {
		if strings.HasPrefix(arg, "--") {
			dashedArgs[i] = arg
		} else if strings.HasPrefix(arg, "-") {
			dashedArgs[i] = "-" + arg
		} else {
			dashedArgs[i] = arg
		}
	}
	return dashedArgs
}
