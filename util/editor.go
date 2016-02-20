package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// ReceiveInputFromEditor opens the text editor that corresponds to the current
// $EDITOR environment variable, and pipes the saved input back through a
// channel to the caller.
func ReceiveInputFromEditor(input chan string) {
	editor := os.Getenv("EDITOR")
	tempDir := os.TempDir()
	tempFile, tempFileErr := ioutil.TempFile(tempDir, "ATOMIC_INPUT")

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error receiving input from $EDITOR: %s", err)
			input <- ""
		}
	}()

	if tempFileErr != nil {
		panic(tempFileErr)
	}

	defer os.Remove(tempFile.Name())
	editorPath, editorPathError := exec.LookPath(editor)

	if editorPathError != nil {
		panic(editorPathError)
	}

	process := exec.Command(editorPath, tempFile.Name())
	process.Stdin = os.Stdin
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	processError := process.Run()

	if processError != nil {
		panic(processError)
	}

	fileContents, fileContentsError := ioutil.ReadFile(tempFile.Name())

	if fileContentsError != nil {
		panic(fileContentsError)
	}

	input <- string(fileContents)
}
