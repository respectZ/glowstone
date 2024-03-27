package texts

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	g_util "github.com/respectZ/glowstone/util"
)

type Lang [][]string

type ILang interface {
	Add(string, string)
	Get(string) (string, bool)
	Remove(string)
	Clear()
	All() map[string]string
	IsEmpty() bool
	Size() int
	Has(string) bool

	// Checks if the lang file is tidy (doesn't have any empty newlines or comments).
	IsTidy() bool

	Load(string) error
	Save(string) error

	// Removes all comments and empty lines from the lang file.
	// Also sorts the keys alphabetically.
	Tidy()
}

func (m *Lang) Add(key string, value string) {
	*m = append(*m, []string{key, value})
}

func (m *Lang) Get(key string) (string, bool) {
	for _, kv := range *m {
		if kv[0] == key {
			return kv[1], true
		}
	}
	return "", false
}

func (m *Lang) Remove(key string) {
	for i, kv := range *m {
		if kv[0] == key {
			*m = append((*m)[:i], (*m)[i+1:]...)
			return
		}
	}
}

func (m *Lang) Clear() {
	*m = [][]string{}
}

func (m *Lang) All() map[string]string {
	mm := make(map[string]string)
	for _, kv := range *m {
		mm[kv[0]] = kv[1]
	}
	return mm
}

func (m *Lang) IsEmpty() bool {
	return len(*m) == 0
}

func (m *Lang) Size() int {
	return len(*m)
}

func (m *Lang) Has(key string) bool {
	for _, kv := range *m {
		if kv[0] == key {
			return true
		}
	}
	return false
}

func (m *Lang) IsTidy() bool {
	// Check if it doesn't have empty keys
	for _, kv := range *m {
		if kv[0] == "" {
			return false
		}
	}
	return true
}

func (m *Lang) Load(pathToRP string) error {
	src := filepath.Join(pathToRP, "texts", "en_US.lang")
	if _, err := os.Stat(src); err != nil {
		// File does not exist
		return err
	}

	file, err := os.Open(src)
	if err != nil {
		return err
	}

	defer file.Close()

	data := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			split := strings.Split(line, "=")
			key := split[0]
			value := strings.Join(split[1:], "=")
			data = append(data, []string{key, value})
		} else {
			// Add to data with empty key
			data = append(data, []string{"", line})
		}
	}

	*m = data

	return nil
}

func (m *Lang) Save(pathToRP string) error {
	// To prevent overwriting, we will read the file first.
	// If the file exists, we will read it and append the new data.

	dst := filepath.Join(pathToRP, "texts", "en_US.lang")

	tempLang := &Lang{}
	if _, err := os.Stat(dst); err == nil {
		// File exists
		if err := tempLang.Load(pathToRP); err != nil {
			return err
		}
	}

	// Append new data
	for _, kv := range *m {
		if tempLang.Has(kv[0]) {
			continue
		}
		tempLang.Add(kv[0], kv[1])
	}

	// Check for tidyness
	if tempLang.IsTidy() && m.IsTidy() {
		tempLang.Tidy()
	}

	// Write to file
	s := ""

	for _, kv := range *tempLang {
		if kv[0] == "" {
			s += kv[1] + "\n"
		} else {
			s += kv[0] + "=" + kv[1] + "\n"
		}
	}
	// Trim the last newline
	s = strings.TrimSuffix(s, "\n")

	g_util.WriteFile(dst, []byte(s))
	if !tempLang.IsTidy() {
		return fmt.Errorf("Lang file contains empty newline and/or comments. Lang will not be sorted.")
	}

	return nil
}

func (m *Lang) Tidy() {
	data := make([][]string, 0)
	for _, kv := range *m {
		if kv[0] != "" {
			data = append(data, kv)
		}
	}
	// Sort
	sort.Slice(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})

	*m = data
}
