package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func encode(b byte) byte {
	if b < 26 {
		return b + 65
	} else if b < 52 {
		return b + 71
	} else if b < 62 {
		return b - 4
	} else if b == 62 {
		return 43
	} else if b == 63 {
		return 47
	} else {
		return 63
	}
}

func main() {
	var InputFileName string
	if len(os.Args) != 2 {
		fmt.Println("The input in not correct")
	} else {
		match, err := regexp.MatchString(`\w+\.\w+`, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		if match {
			InputFileName = os.Args[1]
		} else {
			fmt.Println("The input in not correct")
		}
	}
	buff := make([]byte, 3)
	var n int
	var cont = true
	OutputFileName := regexp.MustCompile(`\.\w+`).ReplaceAllString(InputFileName, "_b64.txt")
	InputFile, err := os.Open(InputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer InputFile.Close()
	OutputFile, err := os.Create(OutputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	for cont {
		n, err = InputFile.Read(buff)
		if err == io.EOF {
			cont = false
		} else if err != nil {
			log.Fatal(err)
		}
		if n == 3 {
			OutputFile.Write([]byte{encode(buff[0] / 4), encode((buff[0]%4)*16 + buff[1]/16), encode((buff[1]%16)*4 + buff[2]/64), encode(buff[2] % 64)})
		} else if n == 2 {
			OutputFile.Write([]byte{encode(buff[0] / 4), encode((buff[0]%4)*16 + buff[1]/16), encode((buff[1] % 16) * 4), '='})
			cont = false
		} else if n == 1 {
			OutputFile.Write([]byte{encode(buff[0] / 4), encode((buff[0] % 4) * 16), '=', '='})
			cont = false
		}
	}

}
