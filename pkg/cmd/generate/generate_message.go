package main

import (
	"bytes"
	"log"
	"os/exec"
)

func generateMessage(root string) {

	cmd := exec.Command(
		"protoc",
		"--gofast_out="+root+"/pkg/message",
		root+"/pkg/message/message.proto",
		"--proto_path="+root+"/pkg/message",
	)

	buffer := &bytes.Buffer{}
	cmd.Stdin = buffer
	cmd.Stdout = buffer
	cmd.Stderr = buffer

	err := cmd.Run()
	if err != nil {
		log.Println(buffer.String())
		log.Fatalln(err)
	}
}
