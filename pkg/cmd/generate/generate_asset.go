package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type image struct {
	name  string
	asset string
}

func generateAsset() {
	inFs, err := filepath.Glob("./pkg/asset/image/*.png")
	if err != nil {
		log.Fatalln(err)
	}

	var images []image
	for _, f := range inFs {
		bs, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatalln(err)
		}

		varName := filepath.Base(f)
		varName = strings.Replace(varName, ".png", "", 1)

		images = append(images, image{varName, string(bs)})
	}

	outF, err := os.Create("./pkg/asset/image_generated.go")
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := outF.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := fmt.Fprintln(outF, "// This file was generated from `./pkg/cmd/generate/generate_asset.go`."); err != nil {
		log.Fatalln(err)
	}

	if _, err := fmt.Fprintln(outF, ""); err != nil {
		log.Fatalln(err)
	}

	if _, err := fmt.Fprintln(outF, "package asset"); err != nil {
		log.Fatalln(err)
	}

	if _, err := fmt.Fprintln(outF, ""); err != nil {
		log.Fatalln(err)
	}

	for _, image := range images {
		if _, err := fmt.Fprintf(outF, "var %s = []byte(%q)\n", image.name, image.asset); err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := fmt.Fprintln(outF, ""); err != nil {
		log.Fatalln(err)
	}

	if _, err := fmt.Fprintln(outF, "func init() {"); err != nil {
		log.Fatalln(err)
	}

	for _, image := range images {
		if _, err := fmt.Fprintf(outF, "	setImage(\"%s\", %s)\n", image.name, image.name); err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := fmt.Fprintln(outF, "}"); err != nil {
		log.Fatalln(err)
	}
}
