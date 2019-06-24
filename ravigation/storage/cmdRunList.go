package storage

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
)

func getRunListField(agvName string) string {
	return CMD_NAMESPACE + ":" + RUN_GROUP + ":" + agvName
}

func SetCmdToRunList(conn redis.Conn, agvName string, ncg NodeCmdGroup) error {
	field := getRunListField(agvName)
	data, err := json.Marshal(ncg)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", field, data)
	return err
}

func GetCmdFromRunList(conn redis.Conn, agvName string, ncg *NodeCmdGroup) error {
	field := getRunListField(agvName)
	data, err := redis.Bytes(conn.Do("GET", field))
	if err != nil {
		return err
	}

	return json.Unmarshal(data, ncg)
}

func DeleteCmdFromRunList(conn redis.Conn, agvName string) (int, error) {
	field := getRunListField(agvName)
	return redis.Int(conn.Do("DEL", field))
}
