package assemble

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	TmpFile     = "/tmp/%s"
	Ext         = ".list"
	NASMCommand = "nasm -f elf64 %s.asm"
	LdCommand   = "ld %s.o -o %s -lc -dynamic-linker /lib/ld-linux-x86-64.so.2"
)

type Output struct {
	OutFilePath string
	TmpFilePath string
	Asm         string
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func newOutput(filePath string, asm string) Output {
	if strings.HasSuffix(filePath, Ext) {
		filePath = strings.TrimSuffix(filePath, Ext)
	}

	out := Output{
		OutFilePath: filePath,
		TmpFilePath: fmt.Sprintf(TmpFile, randomString(7)),
		Asm:         asm,
	}
	return out
}

func (o *Output) createOutputFile() {
	fd, err := os.Create(o.OutFilePath + ".asm")
	HandleErr(err)

	_, err = fd.WriteString(o.Asm)
	fd.Close()
	HandleErr(err)
}

// fucked stupid pointless piece of shit, 100% not my fault, fuck you google
func (o *Output) assemble() {
	cmd1 := fmt.Sprintf(NASMCommand, o.OutFilePath)
	cmd2 := fmt.Sprintf(LdCommand, o.OutFilePath, o.OutFilePath)

	e1 := exec.Command(cmd1)
	e2 := exec.Command(cmd2)
	output1, err := e1.Output()
	HandleErr(err)
	output2, err := e2.Output()
	HandleErr(err)

	fmt.Println(output1)
	fmt.Println(output2)
}

func SaveAndAssemble(filePath string, asm string) {
	output := newOutput(filePath, asm)
	output.createOutputFile()
	//	output.assemble()
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
