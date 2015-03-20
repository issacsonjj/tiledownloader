# tiledownloader
tiledownloader is a library used to download the tiles from TMS service.

## How to install
go get -u github.com/issacsonjj/tiledownloader

## How to use it

```Go
package main

import (
	"fmt"
	"github.com/issacsonjj/tiledownloader"
)

func main() {
	//use openstreet map TMS service
	dl := downloader.NewDownloader("http://a.tile.openstreetmap.org")

	//specify an area in longitude and latitude
	minLon := 121.1
	maxLon := 122.1

	minLat := 31.1
	maxLat := 32.1

	//specify the zoom
	zoom := uint32(2)

	minTileX := downloader.CalTileX(minLon, zoom)
	maxTileX := downloader.CalTileX(maxLon, zoom)

	minTileY := downloader.CalTileY(minLat, zoom)
	maxTileY := downloader.CalTileY(maxLat, zoom)

	fmt.Printf("minTileX: %d, maxTileX: %d, minTileY: %d, maxTileY: %d\n", minTileX, maxTileX, minTileY, maxTileY)

	for x := minTileX; x <= maxTileX; x++ {
		for y := minTileY; y <= maxTileY; y++ {
			tile := new(downloader.Tile)
			tile.Z = int(zoom)
			tile.X = x
			tile.Y = y
			bytes, err := dl.DownloadTile(tile)
			if err != nil || bytes == nil {
				fmt.Printf("download %d/%d/%d fail\n", zoom, x, y)
			} else {
				fmt.Printf("download %d/%d/%d success\n", zoom, x, y)
			}
		}
	}
}
