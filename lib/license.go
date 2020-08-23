package gomrepo

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func getLicense(name string) (string, error) {
	u, err := url.Parse(fmt.Sprintf("https://pkg.go.dev/%s?tab=licenses", name))
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{ServerName: "go.dev"},
	}
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get contents with status code %v", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	license := doc.Find(`#\#lic-0`).Text()

	return license, nil
}
