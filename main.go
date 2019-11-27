package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/MartialBE/openccgo/utils"
	"github.com/go-ego/gse"
	"github.com/liuzl/gocc"
)

var inputFile = flag.String("i", "", "input file path")
var outputFile = flag.String("o", "", "output file path")
var mode = flag.String("m", "s2hk", "mode ")
var dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))

// gse init
var Seg gse.Segmenter

func init() {
	*gocc.Dir = dir + "/resource"
	flag.CommandLine.Usage = func() {
		fmt.Println("-i input file path")
		fmt.Println("-o output file path")
		fmt.Println("-m mode:s2t,t2s,s2tw,tw2s,s2hk,hk2s,s2twp,tw2sp,t2tw,t2hk")
	}
	if err := RestoreAssets(dir+"/", "resource"); err != nil {
		os.RemoveAll(filepath.Join(dir+"/", "resource"))
	}
	flag.Parse()
	Seg.LoadDict(dir + "/resource/gse/dictionary.txt")
}

func main() {

	log.Print(dir)
	if *outputFile == "" {
		*outputFile = "out1_" + *inputFile
	}
	file, err := utils.LoadFile(*inputFile, *outputFile)
	if err != nil {
		log.Fatal(err)
	}

	strArr := Seg.Cut(file.Content, true)
	newStr := ""
	cc, err := gocc.New(*mode)
	if err != nil {
		log.Fatal(err)
	}

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
