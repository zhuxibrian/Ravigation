package storage

import (
	Redis "github.com/gomodule/redigo/redis"
	"strings"
)

type ConnectInfo struct {
	Device string `json:"device"`
	Name string `json:"name"`
	Host string `json:"host"`
	Port uint32 `json:"port"`
}

const CONNECTINFO_NAMESPACE = "ConnectInfo"

func getConnectInfoField(device string, name string) string {
	return CONNECTINFO_NAMESPACE + ":" + device + ":" + name
}

/**
保存连接信息
 */
func CreateConnectInfo(conn Redis.Conn, info ConnectInfo) error {
	field := getConnectInfoField(info.Device, info.Name)
	_, err := conn.Do("HMSET", Redis.Args{field}.AddFlat(info)...)
	return err
}

/**
读取连接信息
 */
func GetConnectInfo(conn Redis.Conn, info *ConnectInfo, device string, name string) error {
	field := getConnectInfoField(device, name)
	data, err := Redis.Values(conn.Do("HGETALL", field))
	if err!= nil {
		return err
	}
	return Redis.ScanStruct(data, info)
}

/**
删除连接信息
 */
func DeleteConnectInfo(conn Redis.Conn, device string, name string) (int,error) {
	field := getConnectInfoField(device, name)
	return Redis.Int(conn.Do("DEL", field))
}

/**
获取所有连接信息
 */
func GetAllConnectInfo(conn Redis.Conn, infoList *[]ConnectInfo) error {
	data, err := Redis.Strings(conn.Do("KEYS", CONNECTINFO_NAMESPACE+"*"))
	if err != nil {
		return err
	}
	for _, v := range data {
		s := strings.Split(v, ":")
		var info ConnectInfo
		if err := GetConnectInfo(conn, &info, s[1], s[2]); err != nil {
			return err
		}

		*infoList = append(*infoList, info)
	}

	return nil
}
