// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package main

import (
	"flag"
	"fmt"
	"github.com/elliotchance/pie/v2"
	"os"
	"strings"
)

var (
	path         = flag.String("path", os.Getenv("PATH"), "PATH to add onto")
	pathFile     = flag.String("file", "~/.paths", "The location of your paths file")
	allowMissing = flag.Bool("allowMissing", true, "Add non-existent directories to PATH")
	allowInvalid = flag.Bool("allowInvalid", false, "Add invalid directories to PATH")
)

func main() {
	flag.Parse()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	if s, ok := strings.CutPrefix(*pathFile, "~"); ok {
		*pathFile = homeDir + s
	}

	b, err := os.ReadFile(*pathFile)
	if err != nil {
		panic(err)
	}

	paths := make([]string, 0)
	validatePath := func(pathsNew []string) {
		for _, p := range pathsNew {
			p = os.ExpandEnv(p)

			if _, err := os.Stat(p); err != nil {
				if os.IsNotExist(err) { // file does not exist
					if !*allowMissing {
						continue
					}
				} else { // other error
					if !*allowInvalid {
						continue
					}
				}
			}

			if p == " " || len(p) == 0 {
				continue
			}

			if pie.Contains(paths, p) {
				continue
			}

			paths = pie.Insert(paths, 0, p)
		}
	}

	validatePath(strings.Split(string(b), "\n"))
	validatePath(strings.Split(*path, ":"))

	fmt.Printf(`export PATH="%s"`, strings.Join(paths, ":"))
}
