package services

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var (
	FileSystemService filesystemServiceInterface = &filesystemService{}
)

const (
	filePath = "/tmp/downtime-notifier-db"
)

type filesystemService struct{}

type filesystemServiceInterface interface {
	IsFailing(serviceName string) bool
	Save(serviceName string, isFailing bool)
}

func doesFileExist() bool {
	// Check if the file exists. If it does not exist there are no failing services.
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func readLines() []string {
	// Open the file.
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Unable to read file: %v", err)
		return nil
	}
	defer file.Close()

	var result []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

// Check if the service failed in the past without a successful recovery.
func (f filesystemService) IsFailing(serviceName string) bool {
	if !doesFileExist() {
		return false
	}

	for _, line := range readLines() {
		if line == serviceName {
			return true
		}
	}

	return false
}

func (f filesystemService) Save(serviceName string, isFailing bool) {
	// If the service is not failing and the file does not exist we can return.
	if !isFailing && !doesFileExist() {
		return
	}

	// Read contents.
	content := readLines()

	// Open the file.
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Unable to save. Error whiler reading file: %v", err)
		return
	}
	defer file.Close()

	// Create an array for the new contents.
	var newContents []string

	for _, line := range content {
		// If we're not failing anymore and this line equals the service name, skip.
		if !isFailing && line == serviceName {
			continue
		}
		newContents = append(newContents, line)
	}

	if isFailing {
		newContents = append(newContents, serviceName)
	}

	arrAsString := strings.Join(newContents, "\n")
	_, err = file.WriteString(arrAsString)
	if err != nil {
		log.Printf("Error while writing to file: %v", err)
	}
}
