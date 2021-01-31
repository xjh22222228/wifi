package main

import (
    "errors"
    "flag"
    "fmt"
    "github.com/skip2/go-qrcode"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "regexp"
    "runtime"
    "strings"
)

var (
    platform = runtime.GOOS
    flagQrcode bool
    flagVersion bool
    flagPassword bool
    flagSSID bool
    flagImage bool
)

const (
    Version = "v1.0.0"
)

func GetSSID() (string, error)  {
    if platform == "darwin" {
        path := "/System/Library/PrivateFrameworks/Apple80211." +
            "framework/Versions/Current/Resources/airport"
        _, err := os.Stat(path)

        if err != nil {
            return "", err
        }

        c := exec.Command(path, "-I")

        v, err := c.CombinedOutput()

        if err != nil {
            fmt.Print(err)
            return "", err
        }

        str := string(v)
        regex, _ := regexp.Compile(" SSID: (.*)")
        f := regex.FindStringSubmatch(str)

        if len(f) < 2 {
            return "", errors.New("SSID: not found")
        }

        return f[1], nil
    }

    return "", errors.New("SSID: not found")
}

func GetPass(ssid string) (string, error) {
    if ssid == "" {
        return "", errors.New("SSID: not found")
    }

    if platform == "darwin" {
        cmd := exec.Command("security", "find-generic-password", "-l", ssid,
            "-D", "AirPort network password", "-w")
        o, err := cmd.CombinedOutput()

        if err != nil {
            return "", errors.New("permission denied")
        }
        return strings.TrimSpace(string(o)), nil
    }

    return "", errors.New("password: Not found")
}

func Qrcode(ssid, pwd string, isOut bool)  {
    text := "WIFI:T:WPA;S:" + ssid + ";P:" + pwd + ";;"
    qr, _ := qrcode.New(text, qrcode.Medium)

    if isOut {
        cwd, _ := os.Getwd()
        path := filepath.Join(cwd, "wifi.png")

        err := qr.WriteFile(256, path)
        if err != nil {
            log.Panicf("%T err: %T", path, err)
        }
        fmt.Println(path)
    } else {
        fmt.Print(qr.ToSmallString(false))
    }
}

func main()  {
    flag.BoolVar(&flagQrcode, "q", false, "Print Qrcode")
    flag.BoolVar(&flagQrcode, "qrcode", false, "Print Qrcode")
    flag.BoolVar(&flagImage, "i", false, "Print Qrcode")
    flag.BoolVar(&flagImage, "image", false, "Print Qrcode")
    flag.BoolVar(&flagVersion, "v", false, "Current Version")
    flag.BoolVar(&flagVersion, "version", false, "Current Version")
    flag.BoolVar(&flagPassword, "p", false, "Print Password")
    flag.BoolVar(&flagPassword, "password", false, "Print Password")
    flag.BoolVar(&flagSSID, "s", false, "Specify a SSID that you have previously connected to")
    flag.BoolVar(&flagSSID, "ssid", false,
        "Specify a SSID that you have previously connected to")
    flag.Parse()

    if flagVersion {
        fmt.Println(Version)
        return
    }

    ssid, err := GetSSID()

    if err != nil {
        fmt.Printf("SSID: %T", err)
        return
    }

    if flagSSID {
        fmt.Println("SSID: " + ssid)
        return
    }

    password, err := GetPass(ssid)

    if err != nil {
        fmt.Println(err)
        return
    }

    if flagQrcode {
        Qrcode(ssid, password, false)
        return
    }

    if flagImage {
        Qrcode(ssid, password, true)
        return
    }

    fmt.Println("password: " + password)
}
