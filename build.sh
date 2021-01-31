#!/bin/bash

GOARCH=amd64

macBuildName=wifi_darwin_${GOARCH}
winBuildName=wifi_windows_${GOARCH}
linuxBuildName=wifi_linux_${GOARCH}

# clear build/
rm -rf wifi_build
mkdir wifi_build

build() {
  # Mac os
  GOOS=darwin GOARCH=${GOARCH} go build cmd/wifi.go
  upx wifi
  mv wifi wifi_build/$macBuildName
  tar -cvf wifi_build/${macBuildName}.tar wifi_build/${macBuildName}
  gzip wifi_build/${macBuildName}.tar
  rm -f wifi_build/${macBuildName}


  # Linux
  GOOS=linux GOARCH=${GOARCH} go build cmd/wifi.go
  upx wifi
  mv wifi wifi_build/$linuxBuildName
  tar -cvf wifi_build/${linuxBuildName}.tar wifi_build/${linuxBuildName}
  gzip wifi_build/${linuxBuildName}.tar
  rm -f wifi_build/${linuxBuildName}


  # Win
  GOOS=windows GOARCH=${GOARCH} go build cmd/wifi.go
  upx wifi.exe
  mv wifi.exe wifi_build/${winBuildName}.exe
  zip -j ${winBuildName}.zip wifi_build/${winBuildName}.exe
  rm -f wifi_build/${winBuildName}.exe
  mv ${winBuildName}.zip wifi_build/${winBuildName}.zip
}

build
