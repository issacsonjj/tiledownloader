// utils.go
package downloader

import (
	"math"
)

//caculate the tile number of @lon in level @zoomLevel
func CalTileX(lon float64, zoomLevel uint32) int {
	m := (lon + 180) / 360.0
	scale := 1 << zoomLevel
	m = m * float64(scale)
	x := int(math.Floor(m))

	return x
}

//caculate the tile number of @lat in level @zoomLevel
func CalTileY(lat float64, zoomLevel uint32) int {
	scale := 1 << zoomLevel
	tmp := 1 - math.Log(math.Tan(lat*math.Pi/180)+1/math.Cos(lat*math.Pi/180))/math.Pi
	return int(math.Floor(tmp / 2.0 * float64(scale)))
}
