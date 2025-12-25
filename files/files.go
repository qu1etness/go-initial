package files

import (
	"os"

	"github.com/fatih/color"
)

func ReadFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteFile(content []byte, fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		color.Red("Failed to create file:", err)
	}

	_, err = file.Write(content)

	defer func() {
		err := file.Close()
		if err != nil {
			color.Red("Failed to close file:", err)
		}
	}()

	if err != nil {
		color.Red("Failed to write to file:", err)
	}
	color.Green("Successfully written to file: %s", fileName)
}
