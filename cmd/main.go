package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/oneclickvirt/CommonMediaTests/commediatests"
	. "github.com/oneclickvirt/defaultset"
)

func main() {
	go func() {
		http.Get("https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Foneclickvirt%2FCommonMediaTests&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false")
	}()
	fmt.Println(Green("项目地址:"), Yellow("https://github.com/oneclickvirt/CommonMediaTests"))
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	languagePtr := flag.String("l", "", "Language parameter (en or zh)")
	flag.BoolVar(&commediatests.EnableLoger, "e", false, "Enable logging")
	flag.Parse()
	if showVersion {
		fmt.Println(commediatests.ComMediaTestsVersion)
		return
	}
	var language string
	if *languagePtr == "" {
		language = "zh"
	} else {
		language = *languagePtr
	}
	language = strings.ToLower(language)
	res := commediatests.MediaTests(language)
	fmt.Printf(res)
}
