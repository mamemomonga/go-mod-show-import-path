package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	if len(os.Args) <= 1 {
		usage()
	}

	fmt.Println(rel2abs(os.Args[1]))
}

func usage() {
	fmt.Println("USAGE: " + os.Args[0] + " target_package_path")
	os.Exit(1)
}

func rel2abs(rel string) string {
	dst, err := filepath.Abs(rel)
	if err != nil {
		log.Fatal(err)
	}

	mod := ""
	depth := 0
	p := dst + "/"
	for p != "/" {
		gm := filepath.Join(p, "go.mod")
		if _, err := os.Stat(gm); !os.IsNotExist(err) {
			b, err := ioutil.ReadFile(gm)
			if err != nil {
				log.Fatal(err)
			}
			{
				s := string(b)
				n := strings.Index(s, "module ") + 7
				e := strings.Index(s[n:], "\n")
				mod = s[n:][:e]
			}
			pmo := strings.Replace(dst, p, "", 1)[1:]
			if depth != 0 {
				mod = filepath.Clean(filepath.Join(mod, pmo))
			}
			break
		}
		p, err = filepath.Abs(filepath.Join(p, ".."))
		if err != nil {
			log.Fatal(err)
		}
		depth++
	}

	return mod
}
