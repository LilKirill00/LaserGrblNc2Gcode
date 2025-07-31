package main

import (
	"flag"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/256dpi/gcode"
)

func main() {
	var (
		inputFile  = flag.String("i", "", "Usage: -i=<input_file>")
		outputFile = flag.String("of", "", "Usage: -of=<output_file>")
		outputDir  = flag.String("od", "", "Usage: -od=<output_dir>")
	)

	flag.Parse()

	mCode := gcode.GCode{Letter: "M", Value: 3}

	f, err := os.Open(*inputFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	file, err := gcode.ParseFile(f)
	if err != nil {
		panic(err)
	}

	ng := make([]gcode.Line, 0)

	for _, v := range file.Lines {
		g := gcode.Line{}
		for i, c := range v.Codes {
			// Фиксируем какой вид M3 или M4 используется
			if c.Letter == "M" && slices.Contains([]float64{3, 4}, c.Value) {
				mCode.Value = c.Value
			}

			// если строка начинается с X,Y,Z то добавить G1 в начало
			if i == 0 && slices.Contains([]string{"X", "Y", "Z"}, c.Letter) {
				g.Codes = append(g.Codes, gcode.GCode{Letter: "G", Value: 1})
			}

			// если в строке есть S и строка не начинается с M то добавить новую строку
			if c.Letter == "S" && v.Codes[0].Letter != "M" {
				ng = append(ng, gcode.Line{Codes: []gcode.GCode{mCode, c}})
				continue
			}

			g.Codes = append(g.Codes, c)
		}

		if len(g.Codes) != 0 {
			ng = append(ng, g)
		}
	}

	// Запись
	var fw *os.File
	if outputFile != nil && *outputFile != "" {
		fw, err = os.Create(*outputFile)
		if err != nil {
			panic(err)
		}
	} else {
		_, base, _ := ParseFile(*inputFile)
		fw, err = os.Create(filepath.Join(*outputDir, base+".gcode"))
		if err != nil {
			panic(err)
		}
	}
	defer fw.Close()

	gcode.WriteFile(fw, &gcode.File{Lines: ng})
}

func ParseFile(file string) (dir string, filename string, ext string) {
	dir = filepath.Dir(file)
	base := filepath.Base(file)
	ext = filepath.Ext(file)
	filename = strings.TrimSuffix(base, ext)
	return
}
