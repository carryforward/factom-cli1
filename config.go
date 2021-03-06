// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"gopkg.in/gcfg.v1"
	"os"
)

type CliConf struct {
	Main struct {
		Server string
		Wallet string
	}
	Entry struct {
		Chainid string
		Extid   string
	}
}

const defaultConf = `
[main]
Server	= localhost:8088
Wallet	= "$HOME/.factom/ecwallet"
[entry]
Chainid	= ""
Extid	= ""
`

func ReadConfig() *CliConf {
	cfg := new(CliConf)
	filename := os.Getenv("HOME") + "/.factom/factom-cli.conf"
	if err := gcfg.ReadFileInto(cfg, filename); err != nil {
		gcfg.ReadStringInto(cfg, defaultConf)
	}
	return cfg
}
