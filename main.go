package main

import (
	"flag"
	"fmt"
	"github.com/MartialBE/opencc_jieba_conversion/utils"
	"github.com/liuzl/gocc"
	"log"
	"os"
	"path/filepath"
)

var inputFile = flag.String("i", "", "input file path")
var outputFile = flag.String("o", "", "output file path")
var mode = flag.String("m", "s2hk", "mode ")

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
}

func main() {
	if *outputFile == ""{
		*outputFile = "out_" + *inputFile
	}
	file, err := utils.LoadFile(*inputFile, *outputFile)
	if err != nil {
		log.Fatal(err)
	}
	jieba := utils.Jieba(file.Content)
	newStr := ""
	cc, err := gocc.New(*mode)
	for i := 0; i < len(jieba); i++ {
		language, err := cc.Convert(jieba[i])
		if err != nil {
			log.Fatal(err)
		}
		newStr += language
	}
	file.Content = newStr
	file.WriteFile()
	log.Print("转化成功~~~")
}
