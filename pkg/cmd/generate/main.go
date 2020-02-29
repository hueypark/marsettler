package main

import (
	"bytes"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	generateFlatbuffers()
	generateAsset()
}

func generateFlatbuffers() {
	files, err := filepath.Glob("./game/message/fbs/*.fbs")
	if err != nil {
		log.Fatalln(err)
	}

	for _, f := range files {
		cmd := exec.Command("flatc", "--go", "-o", "./game/message", f)
		cmdOutput := &bytes.Buffer{}
		cmd.Stdout = cmdOutput
		err := cmd.Run()
		if err != nil {
			log.Println(cmdOutput.String())
			log.Fatalln(err)
		}
	}
}

func generateAsset() {
	files, err := filepath.Glob("./pkg/asset/*.png")
	if err != nil {
		log.Fatalln(err)
	}

	for _, f := range files {
		cmd := exec.Command(
			"file2byteslice",
			"-input", f,
			"-output", strings.Replace(f, ".png", ".go", 1),
			"-package", "asset",
			"-var", strings.Replace(filepath.Base(f), ".png", "", 1))

		cmdOutput := &bytes.Buffer{}
		cmd.Stdout = cmdOutput
		err := cmd.Run()
		if err != nil {
			log.Println(cmdOutput.String())
			log.Fatalln(err)
		}
	}
}
