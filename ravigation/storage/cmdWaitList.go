package storage

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

type NodeCmdGroup struct {
	NodeName string
	BtnName string
	Timestamp int64
	CmdGroupName string
}

func getWaitCmdField() string {
	return CMD_NAMESPACE + ":" + WAIT_GROUP
}

func RPushCmdGroupToWaitList(conn redis.Conn, ncg *NodeCmdGroup) error {
	field := getWaitCmdField()
	data, err := json.Marshal(*ncg)
	if err != nil {
		return err
	}
	_, err = conn.Do("RPUSH", field, string(data))
	return err
}

func LPopCmdGroupFromWaitList(conn redis.Conn, ncg *NodeCmdGroup) error {
	field := getWaitCmdField()
	data, err := redis.Bytes(conn.Do("LPOP", field))
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, ncg)
	if err != nil {
		return err
	}
	return nil
}

func LPushCmdGroupToWaitList(conn redis.Conn, ncg *NodeCmdGroup) error {
	field := getWaitCmdField()
	data, err := json.Marshal(*ncg)
	if err != nil {
		return err
	}
	_, err = conn.Do("LPUSH", field, string(data))
	return err
}

func DeleteCmdGroupFromWaitList(conn redis.Conn, index int) error {
	field := getWaitCmdField()
	conn.Send("MULTI")
	conn.Send("LSET", field, index, "del")
	conn.Send("LREM", field, 0, "del")
	_, err := conn.Do("EXEC")
	return err

}

func GetAllCmdGroupFromWaitList(conn redis.Conn, ncgList *[]NodeCmdGroup) error {
	field := getWaitCmdField()
	replys, err := redis.Values(conn.Do("LRANGE", field, 0, -1))
	if err != nil {
		return err
	}
	for _, r := range replys {
		var ncg NodeCmdGroup
		if err := json.Unmarshal(r.([]byte), &ncg); err != nil {
			logrus.Error("GetAllCmdGroup error")
			continue
		}
		*ncgList = append(*ncgList, ncg)
	}

	return nil
}
