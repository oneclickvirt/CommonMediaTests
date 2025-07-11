# CommonMediaTests

[![Hits](https://hits.spiritlhl.net/CommonMediaTests.svg?action=hit&title=Hits&title_bg=%23555555&count_bg=%230eecf8&edge_flat=false)](https://hits.spiritlhl.net)

[![Build and Release](https://github.com/oneclickvirt/CommonMediaTests/actions/workflows/main.yaml/badge.svg)](https://github.com/oneclickvirt/CommonMediaTests/actions/workflows/main.yaml)

御三家流媒体解锁测试

基于 [netflix-verify](https://github.com/sjlleo/netflix-verify) [VerifyDisneyPlus](https://github.com/sjlleo/VerifyDisneyPlus) [TubeCheck](https://github.com/sjlleo/TubeCheck) 整合代码，优化测试速度

## 功能

- [x] 双栈测试
- [x] 并发测试netflix、youtube、disneyplus是否解锁以及解锁的地区
- [x] 支持双语输出，以```-l```指定```zh```或```en```可指定输出的语言，未指定时默认使用中文输出
- [x] 全平台编译支持

## TODO

- [ ] 整合框架
- [ ] 修复指定英文输出的条件下依然输出中文国家名

## 使用

下载、安装、升级

```shell
curl https://raw.githubusercontent.com/oneclickvirt/CommonMediaTests/main/cmt_install.sh -sSf | bash
```

或

```
curl https://cdn.spiritlhl.net/https://raw.githubusercontent.com/oneclickvirt/CommonMediaTests/main/cmt_install.sh -sSf | bash
```

使用

```
cmt
```

或

```
./cmt
```

进行测试

```
Usage: cmt [options]
  -e    Enable logging
  -h    Show help information
  -l string
        Language parameter (en or zh)
  -v    Show version
```

无环境依赖，理论上适配所有系统和主流架构，更多架构请查看 https://github.com/oneclickvirt/CommonMediaTests/releases/tag/output

![图片](https://github.com/oneclickvirt/CommonMediaTests/assets/103393591/8d4e5aa9-1ab6-4452-af6b-ef3665a902d8)


## 卸载

```
rm -rf /root/cmt
rm -rf /usr/bin/cmt
```

## 在Golang中使用

```
go get github.com/oneclickvirt/CommonMediaTests@v0.0.4-20250629044730
```
