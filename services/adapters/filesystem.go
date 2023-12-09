package adapters

import (
	"os"
	"os/exec"
)

type realFileHandler struct{}

func (handler *realFileHandler) CreateFile(name string, content string) error {
	return os.WriteFile(name, []byte(content), 0644)
}

func (handler *realFileHandler) CopyFile(from string, to string) error {
	data, err := os.ReadFile(from)
	if err != nil {
		return err
	}
	err = os.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Maybe refactor to OS layer, not file handler
func (handler *realFileHandler) RunCommand(command string) error {
	return exec.Command(command).Run()
}

func (handler *realFileHandler) GetCWD() (string, error) {
	return os.Getwd()
}

//go:generate mockgen -source=filesystem.go -destination=filesystem_mock.go -package=adapters
type FileSystemHandler interface {
	CreateFile(name string, content string) error
	CopyFile(from string, to string) error
	RunCommand(command string) error
	GetCWD() (string, error)
}

func NewRealFileSystemHandler() FileSystemHandler {
	return &realFileHandler{}
}
