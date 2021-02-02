#!/bin/bash

GOARCH=amd64

macBuildName=wifi_darwin_${GOARCH}
winBuildName=wifi_windows_${GOARCH}
linuxBuildName=wifi_linux_${GOARCH}

# clear build/
rm -rf wifi_build
mkdir wifi_build

build() {
  # Build
  GOOS=darwin GOARCH=${GOARCH} go build -o $macBuildName cmd/wifi.go
  GOOS=linux GOARCH=${GOARCH} go build -o $linuxBuildName cmd/wifi.go
  GOOS=windows GOARCH=${GOARCH} go build -o $winBuildName.exe cmd/wifi.go

  # Compress
  upx $macBuildName
  upx $linuxBuildName
  upx $winBuildName.exe

  # Move
  mv -f $macBuildName wifi_build/$macBuildName
  mv -f $linuxBuildName wifi_build/$linuxBuildName
  mv -f $winBuildName.exe wifi_build/$winBuildName.exe

  # gzip
  tar -cvf wifi_build/${macBuildName}.tar wifi_build/${macBuildName} && gzip wifi_build/${macBuildName}.tar
  tar -cvf wifi_build/${linuxBuildName}.tar wifi_build/${linuxBuildName} && gzip wifi_build/${linuxBuildName}.tar
  zip -j wifi_build/${winBuildName}.zip wifi_build/${winBuildName}.exe

  # Remove
  rm -f wifi_build/${macBuildName}
  rm -f wifi_build/${linuxBuildName}
  rm -f wifi_build/${winBuildName}.exe
}

build
