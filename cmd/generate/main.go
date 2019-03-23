package main

import (
	"bytes"
	"log"
	"os/exec"
	"path/filepath"
)

func main() {
	generateFlatbuffers()
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
