package main

import (
	"flag"
	"fmt"
	"github.com/MartialBE/openccgo/utils"
	"github.com/go-ego/gse"
	"github.com/liuzl/gocc"
	"log"
	"os"
	"path/filepath"
)

var inputFile = flag.String("i", "", "input file path")
var outputFile = flag.String("o", "", "output file path")
var mode = flag.String("m", "s2hk", "mode ")
var Seg gse.Segmenter

func init() {
	*gocc.Dir = "./resource"
	flag.CommandLine.Usage = func() {
		fmt.Println("-i input file path")
		fmt.Println("-o output file path")
		fmt.Println("-m mode:s2t,t2s,s2tw,tw2s,s2hk,hk2s,s2twp,tw2sp,t2tw,t2hk")
	}
	if err := RestoreAssets("./", "resource"); err != nil {
		os.RemoveAll(filepath.Join("./", "resource"))
	}
	flag.Parse()
	Seg.LoadDict("./resource/gse/dictionary.txt")
}

func main() {
	if *outputFile == ""{
		*outputFile = "out1_" + *inputFile
	}
	file, err := utils.LoadFile(*inputFile, *outputFile)
	if err != nil {
		log.Fatal(err)
	}
	strArr := Seg.Cut(file.Content, true)
	newStr := ""
	cc, err := gocc.New(*mode)
	for i := 0; i < len(strArr); i++ {
		language, err := cc.Convert(strArr[i])
		if err != nil {
			log.Fatal(err)
		}
		newStr += language
	}
	file.Content = newStr
	file.WriteFile()
	log.Print("转化成功~~~")
}
