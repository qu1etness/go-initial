package files

import (
	"os"

	"github.com/fatih/color"
)

func ReadFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func WriteFile(content []byte, fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		color.Red("Failed to create file:", err)
	}

	_, err = file.Write(content)

	defer file.Close()
	if err != nil {
		color.Red("Failed to write to file:", err)
	}
	color.Green("Successfully written to file: %s", fileName)
}
