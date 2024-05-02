package netflix

import (
	"flag"
	"github.com/oneclickvirt/CommonMediaTests/commonmediatests/netflix/printer"
	"github.com/oneclickvirt/CommonMediaTests/commonmediatests/netflix/verify"
)

var custom = flag.String("custom", "", "自定义测试NF影片ID\n绝命毒师的ID是70143836")
var address = flag.String("address", "", "本机网卡的IP")
var proxy = flag.String("proxy", "", "代理地址")

func Netflix(language string) (string, error) {
	flag.Parse()
	r := verify.NewVerify(verify.Config{
		LocalAddr: *address,
		Custom:    *custom,
		Proxy:     *proxy,
	})
	res, _ := printer.Print(*r, language)
	return res, nil
}
