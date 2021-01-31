
<p align="center">
  <p align="center">
    <img src="media/logo.svg" width="200" />
  </p>
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
![](media/screenshot.png)



## Installation

Shell (Mac):
```
curl -fsSL https://raw.githubusercontent.com/xjh22222228/wifi/main/install.sh | bash
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
