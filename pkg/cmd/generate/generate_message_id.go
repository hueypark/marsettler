package main

import (
	"bufio"
	"os"
	"strings"
	"text/template"
)

func generateMessageID(root string) error {
	protoFile, err := os.Open(root + "/pkg/message/message.proto")
	if err != nil {
		return err
	}
	defer protoFile.Close()

	type Message struct {
		ID      int32
		Message string
	}

	type Data struct {
		Messages []Message
	}

	var data Data

	scanner := bufio.NewScanner(protoFile)
	var id int32
	for scanner.Scan() {
		text := scanner.Text()

		if !strings.HasPrefix(text, "message") {
			continue
		}

		text = strings.ReplaceAll(text, "message", "")
		text = strings.ReplaceAll(text, "{", "")
		text = strings.TrimSpace(text)

		data.Messages = append(
			data.Messages,
			Message{
				Message: text,
				ID:      id,
			})

		id++
	}

	tmpl, err := template.ParseFiles(root + "/pkg/message/message_id.tmpl")
	if err != nil {
		return err
	}

	handlerFile, err := os.Create(root + "/pkg/message/message_id.go")
	if err != nil {
		return err
	}
	defer handlerFile.Close()

	return tmpl.Execute(handlerFile, data)
}
