package main

import (
	"strings"
	"bytes"
	"log"
	"bufio"
)

const (
	READ_HEADER int    = 0
	READ_DATA   int    = 1
	KEY_PAYLOAD string = "content"
)

// Core
func ParseMessage(buffer string) map[string]string {
	status := READ_HEADER

	br := bufio.NewReader(strings.NewReader(buffer))

	model := make(map[string]string)

	contentBuffer := bytes.NewBufferString("")
	for {
		temp, _, err := br.ReadLine()
		if err != nil {
			break
		}
		line := string(temp)

		if status == READ_DATA {
			contentBuffer.WriteString(line + "\n")
			continue
		}
		// Empty line. Next read data
		if strings.TrimSpace(line) == "" {
			status = READ_DATA
			continue
		}
		// Handle head
		head := strings.TrimSpace(line)
		index := strings.Index(head, ":")
		if index > 0 && index < len(head) {
			key := strings.TrimSpace(head[0:index])
			value := strings.TrimSpace(head[index+1:])
			model[key] = value
			log.Println(key + " = " + value)
			continue
		}
	}
	model[KEY_PAYLOAD] = contentBuffer.String()
	return model
}
