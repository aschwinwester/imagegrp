package grep

import (
	"golang.org/x/net/html"
	"net/http"
	"log"
)

// visit appends to links each link found in n and returns the result.
func visit(images []string, node *html.Node, options *Options) []string {
	if node.Type == html.ElementNode && node.Data == "img" {
		for _, attribuut := range node.Attr {
			if attribuut.Key == "src" {
				// image src is relative to host
				if (options.Verbose) {
					log.Printf("find image src %s", attribuut.Val)
				}
				images = append(images, attribuut.Val)
			}
		}
	}
	for childNode := node.FirstChild; childNode != nil; childNode = childNode.NextSibling {
		images = visit(images, childNode, options)
	}
	return images
}
/**
func filterImageNode(node *html.Node, options *Options) string {
	var attValue string
	filter := options.CSSClass != ""
	for _, attribuut := range node.Attr {
		if attribuut.Key == "src" {
			// image src is relative to host
			if (options.Verbose) {
				log.Printf("find image src %s", attribuut.Val)
			}
			images = append(images, attribuut.Val)
		}
	}
	if (options.CSSClass == nil) {

	}
}
*/
func findImages(node *html.Node, options *Options) ([]string, error){

	return visit(nil, node, options), nil
}

func fetchHtml(options Options) bool {
	resp, err := http.Get(options.URL)
	if (err != nil) {
		log.Printf("error fetching url %s", options.URL)
		return false
	}
	if (resp.StatusCode != http.StatusOK) {
		log.Printf("error fetching url %s because http status is %d", options.URL, resp.StatusCode)
		return false
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if (err != nil) {
		return false
	}

	imageURLS, err := findImages(doc, &options)
	if (err != nil) {
		log.Printf("finding images failed for url %s", options.URL)
	}
	for _, imageURL:= range imageURLS {
		absImageURL, err := createURL(imageURL, resp)
		if (err == nil) {
			log.Printf("find image url %s", absImageURL)
			fetchAndStoreImage(absImageURL, options.Folder, getFileName(imageURL))
		}

	}
	return true
}
