package util

import(
	"strings"
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

