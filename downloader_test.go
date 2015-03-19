// downloader_test.go
package downloader

import (
	"testing"
)

func TestDownloadering(t *testing.T) {
	targetUrl := "http://a.tile.openstreetmap.org"
	dl := NewDownloader(targetUrl)
	tile := new(Tile)
	tile.Z = 0
	tile.X = 0
	tile.Y = 0

	bytes, err := dl.DownloadTile(tile)
	if err != nil || bytes == nil {
		t.Fatalf("err: %s\n", err.Error())
	}
}

func TestTileCaculating(t *testing.T) {
	//position of Shanghai
	lon := 121.2
	lat := 31.1
	zoom := uint32(2)

	tileX := CalTileX(lon, zoom)
	tileY := CalTileY(lat, zoom)

	if tileX != 3 || tileY != 1 {
		t.Fatalf("caculate tile index fail, x = %d, y = %d\n", tileX, tileY)
	}
}
