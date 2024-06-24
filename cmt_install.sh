#!/bin/bash
#From https://github.com/oneclickvirt/CommonMediaTests
#2024.06.24

rm -rf /usr/bin/cmt
rm -rf cmt
os=$(uname -s)
arch=$(uname -m)

case $os in
  Linux)
    case $arch in
      "x86_64" | "x86" | "amd64" | "x64")
        wget -O cmt https://github.com/oneclickvirt/CommonMediaTests/releases/download/output/CommonMediaTests-linux-amd64
        ;;
      "i386" | "i686")
        wget -O cmt https://github.com/oneclickvirt/CommonMediaTests/releases/download/output/CommonMediaTests-linux-386
        ;;
      "armv7l" | "armv8" | "armv8l" | "aarch64" | "arm64")
        wget -O cmt https://github.com/oneclickvirt/CommonMediaTests/releases/download/output/CommonMediaTests-linux-arm64
        ;;
      *)
        echo "Unsupported architecture: $arch"
        exit 1
        ;;
    esac
    ;;
  Darwin)
    case $arch in
      "x86_64" | "x86" | "amd64" | "x64")
        wget -O cmt https://github.com/oneclickvirt/CommonMediaTests/releases/download/output/CommonMediaTests-darwin-amd64
        ;;
      "i386" | "i686")
        wget -O cmt https://github.com/oneclickvirt/CommonMediaTests/releases/download/output/CommonMediaTests-darwin-386
        ;;
      "armv7l" | "armv8" | "armv8l" | "aarch64" | "arm64")
        wget -O cmt https://github.com/oneclickvirt/CommonMediaTests/releases/download/output/CommonMediaTests-darwin-arm64
        ;;
      *)
        echo "Unsupported architecture: $arch"
        exit 1
        ;;
    esac
    ;;
  FreeBSD)
    case $arch in
      amd64)
        wget -O cmt https://github.com/oneclickvirt/CommonMediaTests/releases/download/output/CommonMediaTests-freebsd-amd64
        ;;
      "i386" | "i686")
        wget -O cmt https://github.com/oneclickvirt/CommonMediaTests/releases/download/output/CommonMediaTests-freebsd-386
        ;;
      "armv7l" | "armv8" | "armv8l" | "aarch64" | "arm64")
        wget -O cmt https://github.com/oneclickvirt/CommonMediaTests/releases/download/output/CommonMediaTests-freebsd-arm64
        ;;
      *)
        echo "Unsupported architecture: $arch"
        exit 1
        ;;
    esac
    ;;
  OpenBSD)
    case $arch in
      amd64)
        wget -O cmt https://github.com/oneclickvirt/CommonMediaTests/releases/download/output/CommonMediaTests-openbsd-amd64
        ;;
      "i386" | "i686")
        wget -O cmt https://github.com/oneclickvirt/CommonMediaTests/releases/download/output/CommonMediaTests-openbsd-386
        ;;
      "armv7l" | "armv8" | "armv8l" | "aarch64" | "arm64")
        wget -O cmt https://github.com/oneclickvirt/CommonMediaTests/releases/download/output/CommonMediaTests-openbsd-arm64
        ;;
      *)
        echo "Unsupported architecture: $arch"
        exit 1
        ;;
    esac
    ;;
  *)
    echo "Unsupported operating system: $os"
    exit 1
    ;;
esac

chmod 777 cmt
cp cmt /usr/bin/cmt
