package texts

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	g_util "github.com/respectZ/glowstone/util"
)

type Lang map[string]string

type ILang interface {
	Add(string, string)
	Get(string) (string, bool)
	Remove(string)
	Clear()
	All() map[string]string
	IsEmpty() bool
	Size() int

	Save(string) error
}

func (m *Lang) Add(key string, value string) {
	(*m)[key] = value
}

func (m *Lang) Get(key string) (string, bool) {
	value, ok := (*m)[key]
	return value, ok
}

func (m *Lang) Remove(key string) {
	delete(*m, key)
}

func (m *Lang) Clear() {
	*m = make(map[string]string)
}

func (m *Lang) All() map[string]string {
	return *m
}

func (m *Lang) IsEmpty() bool {
	return len(*m) == 0
}

func (m *Lang) Size() int {
	return len(*m)
}

func (m *Lang) Load(pathToRP string) error {
	src := filepath.Join(pathToRP, "texts", "en_US.lang")
	if _, err := os.Stat(src); err == nil {
		// File exists
		lang, err := g_util.Loadlang(src)
		if err != nil {
			return err
		}
		*m = lang
		return nil
	}
	return fmt.Errorf("file does not exist: %s", src)
}

func (m *Lang) Save(pathToRP string) error {
	var s string
	var keys []string
	for k := range *m {
		keys = append(keys, k)
	}
	// Sort keys
	sort.Strings(keys)

	for _, k := range keys {
		s += k + "=" + (*m)[k] + "\n"
	}

	g_util.Writefile(pathToRP, []byte(s))

	return nil
}

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
