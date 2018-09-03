package util

import(
	"strings"
	"image/png"
    "os"
    "github.com/boombuler/barcode"
    "github.com/boombuler/barcode/qr"
)


/*
requestUrl: http://192.168.121.2:8080/index.html
requestHome: index.html

return: http://192.168.121.2:8080
*/
func GetIpPort(requestUrl string, requestHome string) (string) {
	index := strings.Index(requestUrl, requestHome)
	return Substr2(requestUrl, 0, index -1)
}


func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("... start is wrong")
	}

	if end < 0 || end > length {
		panic("... end is wrong")
	}

	return string(rs[start:end])
}

func GenerateQrImg(url string, path string, imgName string) {
	CreatePathIfNotExist(path)

    qrCode, _ := qr.Encode(url, qr.M, qr.Auto)

    qrCode, _ = barcode.Scale(qrCode, 256, 256)

    file, _ := os.Create(path + imgName)
    defer file.Close()

    png.Encode(file, qrCode)
}