package utils

import "github.com/yanyiwu/gojieba"

func Jieba(str string) (split []string) {
	obj := gojieba.NewJieba()
	defer obj.Free()
	split = obj.Cut(str, true)
	return
}