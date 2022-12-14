package tools

import "math"

// GetDistance 通过经纬度计算距离, 返回单位为：千米
func GetDistance(longitude1, latitude1, longitude2, latitude2 float64) float64 {
	radius := 6371000.0 //6378137.0
	rad := math.Pi / 180.0
	latitude1 = latitude1 * rad
	longitude1 = longitude1 * rad
	latitude2 = latitude2 * rad
	longitude2 = longitude2 * rad
	theta := longitude2 - longitude1
	dist := math.Acos(math.Sin(latitude1)*math.Sin(latitude2) + math.Cos(latitude1)*math.Cos(latitude2)*math.Cos(theta))
	return dist * radius / 1000
}
