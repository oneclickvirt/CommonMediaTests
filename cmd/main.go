package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/oneclickvirt/CommonMediaTests/commonmediatests/netflix"
	"github.com/oneclickvirt/CommonMediaTests/commonmediatests/website"
)

func main() {
	go func() {
		http.Get("https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Foneclickvirt%2Fbacktrace&count_bg=%2323E01C&title_bg=%23555555&icon=sonarcloud.svg&icon_color=%23E7E7E7&title=hits&edge_flat=false")
	}()

	var (
		res0, res1, res2 string
		wg               sync.WaitGroup
	)

	languagePtr := flag.String("l", "", "Language parameter (en or zh)")
	flag.Parse()

	var language string

	if *languagePtr == "" {
		language = "zh"
	} else {
		language = *languagePtr
	}

	language = strings.ToLower(language)

	switch language {
	case "en":
		wg.Add(3)
		func() {
			defer wg.Done()
			res1, _ = website.YoutubeCheck("en")
		}()
		func() {
			defer wg.Done()
			res0, _ = netflix.Netflix("en")
		}()
		func() {
			defer wg.Done()
			res2, _ = website.Disneyplus("en")
		}()
	case "zh":
		wg.Add(3)
		func() {
			defer wg.Done()
			res1, _ = website.YoutubeCheck("zh")
		}()
		func() {
			defer wg.Done()
			res0, _ = netflix.Netflix("zh")
		}()
		func() {
			defer wg.Done()
			res2, _ = website.Disneyplus("zh")
		}()
	default:
		fmt.Println("不支持的语言参数")
		return
	}

	wg.Wait()

	fmt.Println("----------------Netflix-----------------")
	fmt.Printf(res0)
	fmt.Println("----------------Youtube-----------------")
	fmt.Printf(res1)
	fmt.Println("---------------DisneyPlus---------------")
	fmt.Printf(res2)
}
