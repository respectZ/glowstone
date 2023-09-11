package lang

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func Create(LangPath string) error {
	// Check if file exists
	if _, err := os.Stat(LangPath); err == nil {
		// File exists
		return errors.New("file already exists")
	} else if os.IsNotExist(err) {
		// File does not exist
		// Create file
		file, err := os.Create(LangPath)
		if err != nil {
			return err
		}
		defer file.Close()
	} else {
		// File may or may not exist
		return err
	}
	return nil
}

func Read(LangPath string) (map[string]string, error) {
	if _, err := os.Stat(LangPath); err == nil {
		// File exists
		return nil, errors.New("file already exists")
	}
	// Read file
	file, err := os.Open(LangPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			split := strings.Split(line, "=")
			key := split[0]
			value := split[1]
			data[key] = value
		}
	}

	return data, nil
}
