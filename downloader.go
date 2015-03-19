// utils.go
package downloader

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Tile struct {
	X int
	Y int
	Z int //zoom
}

type Downloader struct {
	Server string
	Client *http.Client
}

func NewDownloader(server string) *Downloader {
	ret := new(Downloader)

	ret.Server = server
	ret.Client = &http.Client{}

	return ret
}

func (downloader *Downloader) DownloadTile(t *Tile) (bytes []byte, err error) {
	bytes = nil
	err = nil
	reqUrl := fmt.Sprintf("%s/%d/%d/%d.png", downloader.Server, t.Z, t.X, t.Y)
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return
	}
	resp, err := downloader.Client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	bytes, err = ioutil.ReadAll(resp.Body)
	return
}
