package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/oneclickvirt/CommonMediaTests/commediatests"
	. "github.com/oneclickvirt/defaultset"
)

func main() {
	go func() {
		http.Get("https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Foneclickvirt%2FCommonMediaTests&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false")
	}()
	fmt.Println(Green("项目地址:"), Yellow("https://github.com/oneclickvirt/CommonMediaTests"))
	var showVersion, help bool
	var language string
	cmtFlag := flag.NewFlagSet("cmt", flag.ContinueOnError)
	cmtFlag.BoolVar(&help, "h", false, "Show help information")
	cmtFlag.BoolVar(&showVersion, "v", false, "Show version")
	cmtFlag.StringVar(&language, "l", "", "Language parameter (en or zh)")
	cmtFlag.BoolVar(&commediatests.EnableLoger, "e", false, "Enable logging")
	cmtFlag.Parse(os.Args[1:])
	if help {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		cmtFlag.PrintDefaults()
		return
	}
	if showVersion {
		fmt.Println(commediatests.ComMediaTestsVersion)
		return
	}
	if language == "" {
		language = "zh"
	}
	language = strings.ToLower(language)
	res := commediatests.MediaTests(language)
	fmt.Printf(res)
}
