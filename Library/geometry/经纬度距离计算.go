package main

// 计算原理参考: https://blog.csdn.net/gatieme/article/details/45599581

import (
	"math"
)

const (
	// 地球半径
	earthRadiusMi = 3958 // 英里
	earthRaidusKm = 6371 // 千米
)

type Coord struct { // 经纬度坐标
	Lat float64
	Lon float64
}

// 度数 --> 弧度
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

// 计算任意两个坐标点的劣弧(小于半圆)、优弧（大于半圆）
func Distance(p, q Coord) (float64, float64) { // 英里、千米
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
