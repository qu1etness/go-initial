package files

import (
	"os"

	"github.com/fatih/color"
)

func ReadFile() {

}

func WriteFile(content string, fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		color.Red("Failed to create file:", err)
	}

	_, err = file.WriteString(content)

	defer file.Close()
	if err != nil {
		color.Red("Failed to write to file:", err)
	}
	color.Green("Successfully written to file: %s", fileName)
}
