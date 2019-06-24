package storage

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"strings"
)

/**
坐标信息
 */
type Point struct {
	X float64
	Y float64
	Z float64
}

const (
	POINT_NAMESPACE= "Point"
)

func getPointField(index int) string {
	return POINT_NAMESPACE + ":" + strconv.Itoa(index)
}

func NewPoint(x float64, y float64) Point {
	return Point{
		X: x,
		Y: y,
		Z: 0,
	}
}


/**
添加点坐标
 */
func AddPoint(conn redis.Conn, index int, point Point) error {
	field := getPointField(index)
	data, err := json.Marshal(point)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", field, string(data))
	return err
}

/**
获得指定点坐标
 */
func GetPoint(conn redis.Conn, point *Point, index int) error {
	field := getPointField(index)
	str, err := redis.String(conn.Do("GET", field))
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(str), point)
}

func DeletePoint(conn redis.Conn, index int) (int, error) {
	field := getPointField(index)
	return redis.Int(conn.Do("Del", field))
}

/**
根据坐标获得坐标index
 */
func GetPointIndex(conn redis.Conn, point Point) (int, error) {
	field := POINT_NAMESPACE + "*"
	var index = 0
	keys, err := redis.Strings(conn.Do("KEYS", field))
	if err != nil {
		return 0, err
	}

	for _, key := range keys {
		data, err := redis.Bytes(conn.Do("GET", key))
		if err != nil {
			return 0, err
		}
		var p Point
		json.Unmarshal(data, &p)
		if p.X == point.X && p.Y == point.Y {
			strs := strings.Split(key, ":")
			index, err = strconv.Atoi(strs[len(strs)-1])
			break
		}
	}

	return index, nil

}