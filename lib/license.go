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

// GomClient is the client to get module info from 'go.dev' host.
type GomClient struct {
	client *http.Client
}

// NewGomClient returns a new client for go module report.
func NewGomClient() *GomClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{ServerName: "go.dev"},
	}
	client := &http.Client{
		Transport: tr,
	}
	return &GomClient{client: client}
}

// GetLicense returns the license name specified by args.
func (g *GomClient) GetLicense(name string) (string, error) {
	u, err := url.Parse(fmt.Sprintf("https://pkg.go.dev/%s?tab=licenses", name))
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", err
	}

	resp, err := g.client.Do(req)
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

// pkginfo is the info of packages.
type pkginfo struct {
	ID      string
	Version string
	License string
	Error   error
}

const semaphore = 10

// GetLicenseList returns the formated license table with id, version and license.
func (g *GomClient) GetLicenseList(w io.Writer, modules []string, tf Formatter) error {
	pkgCn := make(chan pkginfo)
	tokens := make(chan struct{}, semaphore)
	var counter int

	for _, module := range modules {
		fields := strings.Fields(module)
		if len(fields) < 2 {
			continue
		}
		counter++
		go func(id, ver string) {
			tokens <- struct{}{}
			lic, err := g.GetLicense(id)
			<-tokens
			pkgCn <- pkginfo{id, ver, lic, err}
		}(fields[0], fields[1])
	}

	var pkgs []pkginfo
	var err error
	for counter > 0 {
		pkg := <-pkgCn
		if pkg.Error != nil {
			err = fmt.Errorf("failed to get licenses: %v", pkg.Error)
		} else {
			pkgs = append(pkgs, pkg)
		}
		counter--
	}
	if err != nil {
		return err
	}

	return tf.table(w, pkgs)
}
