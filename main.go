package main

import (
	"fmt"
	"os"
)

const JD_PATH_KEY = "JD_PATH"
const JD_CONF_PATH = "/usr/local/etc/jd"

// template must be .txt. .md does not work
const JD_TEMPLATE_NAME = "journalTemplate.txt"

// TODO exit with the name of the file which can be piped into another command
// TODO allow tags to be added with the -t --tag flag (can be multiple), which get put into a section
// TODO allow pages to be titled with whatever the person gives as an arg, if they don't, then title it date format: 'Saturday 9th April 2022'

func main() {
	settings := getSettings()

	pf := generatePageFile(settings.filenameFullPath())
	defer pf.Close()

	err := writeTemplateToFile(pf, settings.toTemplateArgs())
	if err != nil {
		os.Remove(settings.filenameFullPath())
		exitFatal("failed to write template. error: %v", err)
	}

	// output the fully qualified name of the file
	if !settings.quiet {
		fmt.Print(pf.Name())
	}
}

func generatePageFile(fileName string) *os.File {
	if fileName == "" {
		fileName = generateFilename()
	}
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		exitFatal("could not create file: %s. error: %v", fileName, err)
	}
	return file
}
