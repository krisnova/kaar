// Copyright © 2022 Kris Nóva <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//  Kubernetes Application Archive
//
//  ██╗  ██╗ █████╗  █████╗ ██████╗
//  ██║ ██╔╝██╔══██╗██╔══██╗██╔══██╗
//  █████╔╝ ███████║███████║██████╔╝
//  ██╔═██╗ ██╔══██║██╔══██║██╔══██╗
//  ██║  ██╗██║  ██║██║  ██║██║  ██║
//  ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝
//

package main

import (
	"fmt"
	"os"

	"github.com/kris-nova/kaar"
	"github.com/kris-nova/logger"
	"github.com/urfave/cli"
)

func main() {
	err, i := run()
	if err != nil {
		logger.Critical("runtime error: %v", err)
		os.Exit(i)
	}
}

func run() (error, int) {
	var extract bool
	var create bool
	var verbose bool
	var file bool

	app := &cli.App{
		Name:      "Kubernetes Application Archive",
		Author:    "Kris Nóva <kris@nivenly.com>",
		Version:   kaar.Version,
		UsageText: `kaar [flags] [archive.kaar] [path/]`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "x",
				Usage:       "Extract. Used to extract a kaarball.",
				Destination: &extract,
			},
			&cli.BoolFlag{
				Name:        "c",
				Usage:       "Create. Used to create a kaarball.",
				Destination: &create,
			},
			&cli.BoolFlag{
				Name:        "f",
				Usage:       "File. Used to pass a file path.",
				Destination: &file,
			},
			&cli.BoolFlag{
				Name:        "v",
				Usage:       "Verbose. Toggle verbose logs.",
				Destination: &verbose,
			},
		},
		Action: func(c *cli.Context) error {
			logger.BitwiseLevel = logger.LogAlways
			if verbose {
				logger.BitwiseLevel = logger.LogEverything
			}
			if create && extract {
				return fmt.Errorf("invalid usage: unable to create and extract")
			}
			var kfile, kdir string
			if !file {
				// Read from stdin
				// Write to stdout
				return fmt.Errorf("STDIN and STDOUT not supported")
			}
			if c.NArg() != 2 {
				return fmt.Errorf("invalid usage: kaarball, archive")
			}
			kfile = c.Args().Get(0)
			kdir = c.Args().Get(1)
			//fmt.Printf("kaar kaarball(%s) archive(%s)\n", kfile, kdir)

			// Create
			if create {
				_, err := kaar.Create(kdir, kfile)
				if err != nil {
					return fmt.Errorf("unable to create kaarball: %v", err)
				}
			}

			// Extract
			if extract {
				_, err := kaar.Extract(kdir, kfile)
				if err != nil {
					return fmt.Errorf("unable extract kaarball: %v", err)
				}
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return err, 1
	}
	return nil, 0
}
