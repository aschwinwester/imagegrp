package grep

import (
	"strings"
	"log"
	"fmt"
	"net/http"
)

var imageExtensions = [...]string{"jpg", "jpeg", "gif", "png"}
func IncreaseUrl(url string) string {
	lastSlashIndex := strings.LastIndex(url, "/")
	if (lastSlashIndex <= 0 || lastSlashIndex >= len(url)) {
		return ""
	}
	fileNameWithExtension := url[lastSlashIndex + 1 : len(url)]
	splitArray := strings.Split(fileNameWithExtension, ".")
	count := len(splitArray)
	if (count != 2) {
		log.Printf("filename %s splitted gave array of %d", fileNameWithExtension, count)
	}

	increasedFileName := increaseString(splitArray[0]) + "." + splitArray[1]
	return increasedFileName

}

func increaseString(name string) string {
	last := name[len(name)-1:]
	fmt.Println(last)
	return ""
}

func IsImage(url string) bool {
	for _, ext := range imageExtensions {
		if (strings.HasSuffix(url, ext)) {
			return true;
		}
	}
	return false;
}

func createURL(imageUrlPart string, response *http.Response) (string, error) {
	https := "https://"
	http := "http://"

	if strings.Contains(imageUrlPart, https) || strings.Contains(imageUrlPart, http) {
		return imageUrlPart, nil
	}

	url, err := response.Request.URL.Parse(imageUrlPart);

	if (err != nil) {
		return "", err
	}
	return url.String(), nil
}
