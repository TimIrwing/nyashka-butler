package messages

import (
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"syscall"
)

const dataDirPath = "data"
const messagesPath = dataDirPath + "/messages.json"

type Messages map[string]string

var defaultMessages = &Messages{
	"welcome": "Приветствую!",
	"goodbye": "Прощай!",
}

func (m Messages) ReadMessageFile() {
	b, err := os.ReadFile(messagesPath)
	if v, ok := err.(*fs.PathError); ok && v.Err != syscall.ENOENT {
		log.Printf("Error reading message file: %s", err)
	}

	err = json.Unmarshal(b, &m)
	if err != nil {
		for k, v := range *defaultMessages {
			if m[k] == "" {
				m[k] = v
			}
		}
	}
}

func (m Messages) WriteMessageFile() bool {
	b, err := json.Marshal(m)
	if err != nil {
		log.Printf("Error encoding message file: %s", err)
		return false
	}

	err = createDataDir()
	if err != nil {
		log.Printf("Error creating data dir: %s", err)
		return false
	}
	err = os.WriteFile(messagesPath, b, 0644)
	if err != nil {
		log.Printf("Error writing message file: %s", err)
		return false
	}
	return true
}

func createDataDir() error {
	return os.MkdirAll(dataDirPath, 0644)
}
