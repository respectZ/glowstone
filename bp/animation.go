package bp

import (
	"os"
	"path/filepath"
	"strings"

	bp "github.com/respectZ/glowstone/bp/animation"
	g_util "github.com/respectZ/glowstone/util"
)

type AnimationBP map[string]*AnimationFile

type IAnimationBP interface {
	Add(string, *bp.Animation)
	Get(string) (*bp.Animation, bool)
	Remove(string)
	Clear()
	All() map[string]*AnimationFile
	IsEmpty() bool
	Size() int
	UnmarshalJSON([]byte) error

	// New creates a new BPAnimation. Dest doesn't need to start with "animation/"
	//
	// Example:
	//
	//	a := New("player.animation.json")
	//	b := New("player/effect.animation.json")
	New(string) *bp.Animation

	Save(string) error

	LoadAll(string) error
}

func (m *AnimationBP) Add(key string, value *bp.Animation) {
	p, ok := (*m)[key]
	if ok {
		p.Data = value
	} else {
		(*m)[key] = &AnimationFile{
			Data: value,
		}
	}
}

func (m *AnimationBP) Get(key string) (*bp.Animation, bool) {
	value, ok := (*m)[key]
	if !ok {
		return nil, false
	}
	return value.Data, true
}

func (m *AnimationBP) Remove(key string) {
	delete(*m, key)
}

func (m *AnimationBP) Clear() {
	*m = make(map[string]*AnimationFile)
}

func (m *AnimationBP) All() map[string]*AnimationFile {
	return *m
}

func (m *AnimationBP) IsEmpty() bool {
	return len(*m) == 0
}

func (m *AnimationBP) Size() int {
	return len(*m)
}

func (m *AnimationBP) UnmarshalJSON(data []byte) error {
	return g_util.UnmarshalJSON(data, m)
}

func (m *AnimationBP) New(dest string) *bp.Animation {
	a := bp.New()
	m.Add(dest, a)
	return a
}

func (m *AnimationBP) Save(pathToBP string) error {
	for dest, v := range m.All() {
		data, err := v.Encode()
		if err != nil {
			return err
		}
		err = g_util.WriteFile(filepath.Join(pathToBP, destDirectory.Animation, dest), data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *AnimationBP) LoadAll(pathToBP string) error {
	files, err := g_util.Walk(filepath.Join(pathToBP, destDirectory.Animation))
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	for _, file := range files {
		a := bp.New()
		err := g_util.LoadJSON(file, &a)
		if err != nil {
			return err
		}
		// Remove the pathToBP from the file path
		file := strings.TrimPrefix(file, pathToBP+string(filepath.Separator)+destDirectory.Animation+string(filepath.Separator))
		m.Add(file, a)
	}
	return nil
}
