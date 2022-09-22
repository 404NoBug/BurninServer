package sensitive

import (
	"fmt"
	"testing"
)

func TestNewSensitiveTrie(t *testing.T) {
	sensitiveWords := []string{
		"傻逼",
		"傻逼东西",
		"傻叉",
		"垃圾",
		"垃圾玩意",
	}
	sensitiveTrie.AddWords(sensitiveWords)
	match, text := sensitiveTrie.Match("什么垃圾打野，傻逼一样，叫你来开龙不来，sb")
	fmt.Println("match", match)
	fmt.Println("text", text)
}

func TestHansCovertPinyin(t *testing.T) {
	sensitiveWords := []string{
		"傻逼",
		"傻叉",
		"垃圾",
		"妈的",
		"sb",
	}
	// 汉字转拼音
	pinyinContents := HansCovertPinyin(sensitiveWords)
	fmt.Println(pinyinContents)
}
