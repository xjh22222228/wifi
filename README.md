
<p align="center">
  <p align="center">
    <b>wifi</b>
  </p>
  <p align="center">Quickly fetch your WiFi password and if needed, generate a QR code of your WiFi to allow phones to easily connect.</p>

  <p align="center">
    <img src="https://img.shields.io/github/go-mod/go-version/xjh22222228/wifi" />
    <img src="https://img.shields.io/github/v/release/xjh22222228/wifi" />
    <img src="https://img.shields.io/github/license/xjh22222228/wifi" />
  </p>
</p>


## Demo
![Preview](https://raw.githubusercontent.com/xjh22222228/public/gh-pages/wifi/screenshot.gif)



## Installation

Shell (Mac):
```bash
curl -fsSL https://raw.githubusercontent.com/xjh22222228/wifi/main/install.sh | bash

# or
curl -fsSL https://raw.sevencdn.com/xjh22222228/wifi/main/install.sh | bash
```

Windows:

[Download](https://github.com/xjh22222228/wifi/releases/latest/download/wifi_windows_amd64.zip)
```
./wifi.exe
```





## Options

```
-q, --qrcode           Print Qrcode
-i, --image            Generate QR code image
-v, --version          Output the version number
-p, --password         Print Password
-s, --ssid             Specify a SSID that you have previously connected to
-h, --help             Display help for command
```

## Example
```
$ wifi

$ wifi -q
```
