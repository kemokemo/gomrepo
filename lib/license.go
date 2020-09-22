package gomrepo

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetLicense returns the license name specified by args.
func GetLicense(name string) (string, error) {
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

// PrintLicenses prints all licenses of 'modules' to the 'w' writer.
func PrintLicenses(w io.Writer, modules []string) error {
	var err error
	for _, module := range modules {
		fields := strings.Fields(module)
		if len(fields) < 2 {
			continue
		}
		lic, e := GetLicense(fields[0])
		if e != nil {
			err = fmt.Errorf("%v: %v", err, e)
		}
		fmt.Fprintln(w, fmt.Sprintf("%s %s %s", fields[0], fields[1], lic))
	}
	return err
}
