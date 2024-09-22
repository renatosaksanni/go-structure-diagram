package renderer

import (
	"fmt"
	"go-structure-diagram/pkg/diagram"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Render takes the generated diagrams and outputs them in the desired format.
func Render(diagrams diagram.Diagrams, format, outputPath string) error {
	// Combine all diagrams into one Mermaid file
	content := diagrams.ClassDiagram + "\n" + diagrams.SequenceDiagram + "\n" + diagrams.StateDiagram

	if strings.ToLower(format) == "mermaid" {
		return ioutil.WriteFile(outputPath, []byte(content), 0644)
	} else if strings.ToLower(format) == "png" {
		// Use mermaid-cli to convert Mermaid to PNG
		tmpFile, err := ioutil.TempFile("", "diagram-*.mmd")
		if err != nil {
			return err
		}
		defer os.Remove(tmpFile.Name())

		if _, err := tmpFile.WriteString(content); err != nil {
			return err
		}
		tmpFile.Close()

		cmd := exec.Command("mmdc", "-i", tmpFile.Name(), "-o", outputPath)
		return cmd.Run()
	} else {
		return fmt.Errorf("unsupported format: %s", format)
	}
}
