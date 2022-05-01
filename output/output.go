package output

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Output struct {
	InPath  string
	OutPath string
}

func NewOutput(path string) *Output {
	out := &Output{
		InPath:  path,
		OutPath: strings.TrimSuffix(path, ".list") + ".asm",
	}
	return out
}

func (o *Output) LoadFile() string {
	file, err := ioutil.ReadFile(o.InPath)
	HandleErr(err)
	return string(file)
}

func (o *Output) WriteAsm(asm string) {
	fd, err := os.Create(o.OutPath)
	HandleErr(err)
	_, err = fd.WriteString(asm)
	HandleErr(err)
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
