package main

import (
	"bufio"
	"os"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

func generateMessageHandler(root string) error {
	protoFile, err := os.Open(root + "/pkg/message/message.proto")
	if err != nil {
		return err
	}
	defer protoFile.Close()

	type Message struct {
		Message string
		Handler string
	}

	type Data struct {
		Messages []Message
	}

	var data Data

	scanner := bufio.NewScanner(protoFile)
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
				Handler: strcase.ToLowerCamel(text) + "Handler",
			})
	}

	tmpl, err := template.ParseFiles(root + "/pkg/internal/net/handler.tmpl")
	if err != nil {
		return err
	}

	handlerFile, err := os.Create(root + "/pkg/internal/net/handler_generated.go")
	if err != nil {
		return err
	}
	defer handlerFile.Close()

	return tmpl.Execute(handlerFile, data)
}
