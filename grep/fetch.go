package grep

import (
	"log"
	"net/http"
	"io/ioutil"
	"strings"
	"os"
)


func Fetch(options Options) {
	if options.Verbose {
		log.Printf("[ok] Found URL %s \n", options.URL)
	}

	if IsImage(options.URL) {
		strategySimpleNumeric(options)
	} else {
		fetchHtml(options)
	}
}

func getFileName(url string) string {
	defaultFileName := "unknown.png"
	l := strings.LastIndex(url, "/")
	if l > 0 && (l + 1 < len(url)) {
		return url[l+1:len(url)]
	}
	log.Printf("[warn] unable to determine filename for url %s. Filename will be %s", url,defaultFileName)
	return defaultFileName
}
func fetchAndStoreImage(url string, folderName string, fileName string) bool {
	var file = fileName
	if (folderName != "") {
		if _, err := os.Stat(folderName); os.IsNotExist(err) {

			os.Mkdir(folderName, 0777)
		}
		file = folderName + "/" + fileName
	}

	resp, err := http.Get(url)


	if err != nil {
		log.Printf("[warn] unable to fetch url %s got error %s", url, err.Error())
		return false;
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("[fatal unable to read body.", err)
		return false
	}
	defer resp.Body.Close()
	err = ioutil.WriteFile(file, contents, 0664)
	if err != nil {
		log.Fatal("[fatal] unable to create file -- ", err)
		return false
	}
	return true
}

func strategySimpleNumeric(options Options) bool {
	success := fetchAndStoreImage(options.URL, options.Folder, getFileName(options.URL))
	if success {
		newUrl := IncreaseUrl(options.URL)
		return fetchAndStoreImage(newUrl, options.Folder, getFileName(newUrl))
	}
	return false
}




