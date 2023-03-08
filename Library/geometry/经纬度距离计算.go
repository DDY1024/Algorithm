package main

import (
	"math"
)

// 参考: https://blog.csdn.net/gatieme/article/details/45599581

const (
	earthRadiusMi = 3958 // 地球半径（英里）
	earthRaidusKm = 6371 // 地球半径（千米）
)

type Coord struct {
	Lat float64 // 纬度坐标
	Lon float64 // 经度坐标
}

// 度数转弧度
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

// 计算任意两个坐标点的距离
// 劣弧：小于半圆的弧
// 优弧：大于半圆的弧
func Distance(p, q Coord) (float64, float64) {
	lat1 := degreesToRadians(p.Lat)
	lon1 := degreesToRadians(p.Lon)
	lat2 := degreesToRadians(q.Lat)
	lon2 := degreesToRadians(q.Lon)
	diffLat := lat2 - lat1
	diffLon := lon2 - lon1
	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return c * earthRadiusMi, c * earthRaidusKm
}
