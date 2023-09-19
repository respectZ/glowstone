package glowstone

import (
	"bufio"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"sort"
	"strings"

	q "github.com/ericpauley/go-quantize/quantize"
	"github.com/respectZ/glowstone/util/jsonc"
)

func Makedir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func MarshalJSON(e interface{}) ([]byte, error) {
	return json.MarshalIndent(e, "", "  ")
}

func Copyfile(source, destination string) error {
	bytesRead, err := os.ReadFile(source)

	if err != nil {
		return err
	}

	// Strip directory
	dir := filepath.Dir(destination)
	// Create directory
	Makedir(dir)

	err = os.WriteFile(destination, bytesRead, os.ModePerm)

	if err != nil {
		return err
	}
	return nil
}

func Writefile(path string, b []byte) error {
	// Strip directory
	dir := filepath.Dir(path)
	// Create directory
	Makedir(dir)
	file, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	defer file.Close()

	_, err := file.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func Writejson(path string, data interface{}) error {
	// Strip directory
	dir := filepath.Dir(path)
	// Create directory
	Makedir(dir)

	file, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(data)
	return err
}

func Writelang(path string, data map[string]string) error {
	// Load lang first
	lang, err := Loadlang(path)
	// Update lang if no error
	if err == nil {
		for k, v := range data {
			lang[k] = v
		}
		data = lang
	}
	// Get keys
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	// Sort keys
	sort.Strings(keys)

	var s string
	for _, k := range keys {
		s += k + "=" + data[k] + "\n"
	}
	return Writefile(path, []byte(s))
}

func Loadlang(path string) (map[string]string, error) {
	file, err := os.Open(path)
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

func LoadJSON(path string, e interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonc.ToJSON(data), e)
	return err
}

// Walk walks the file tree rooted at rootDir, calling walkFn for each file or directory in the tree, including rootDir.
// All errors that arise visiting files and directories are filtered by walkFn.
// The files are walked in lexical order, which makes the output deterministic but means that for very large directories Walk can be inefficient.
// Walk does not follow symbolic links.
func Walk(rootDir string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(rootDir, func(path string, f os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !f.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

/******************* Minecraft Utils *******************/

// ConvertImageOpacity converts an image to specified alpha.
//
// Example:
//
//	// Convert image to 10 alpha
//	ConvertImageOpacity("image.png", "image_new.png" 10)
func ConvertImageOpacity(filePath string, outPath string, alphaValue uint8) error {
	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode the PNG image
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// Create a new RGBA image with the same size as the original image
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	// Apply the alpha value to each pixel in the new image
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Get the original pixel color
			originalColor := img.At(x, y)
			r, g, b, _ := originalColor.RGBA()

			// Create a new color with the specified alpha value
			newColor := color.NRGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: alphaValue,
			}

			// Set the pixel in the new image with the new color
			newImg.Set(x, y, newColor)
		}
	}

	// Create a new output PNG image file
	outputFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Encode and save the new image to the output file in PNG format
	err = png.Encode(outputFile, newImg)
	if err != nil {
		return err
	}
	return nil
}

func getPallete(filePath string) ([]color.Color, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	q := q.MedianCutQuantizer{}
	p := q.Quantize(make([]color.Color, 0, 256), img)
	return p, nil
}

// GetEggColor returns the primary and secondary color for a spawn egg.
//
// Example:
//
//	// Get egg color
//	primary, secondary, err := GetEggColor("image.png")
func GetEggColor(filePath string) (string, string, error) {
	p, err := getPallete(filePath)
	if err != nil {
		return "", "", err
	}
	// Convert p[0] and p1[1] to hex
	r, g, b, _ := p[0].RGBA()
	r1, g1, b1, _ := p[1].RGBA()
	h1 := fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8)
	h2 := fmt.Sprintf("#%02x%02x%02x", r1>>8, g1>>8, b1>>8)
	return h1, h2, nil
}
