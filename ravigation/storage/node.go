package storage

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
)

type Node struct {
	NodeName string
	ButtonList []Button
}

type Button struct {
	ButtonName string
	Status string
}

type ButtonCmdGroup struct {
	NodeName string
	ButtonName string
	CmdGroupName string
}

const (
	NODE_NAMESPACE = "Node"

	NODE_GROUP = "Nodes"
	BUTTONCMD_GROUP = "BTN_CMD"
)

const NODETIMEOUT = 20

func getNodeField(nodeName string) string {
	return NODE_NAMESPACE + ":" + NODE_GROUP + ":" + nodeName
}

func getBtnCmdField(node string) string {
	return NODE_NAMESPACE + ":" + BUTTONCMD_GROUP + ":" + node
}

func SetNodeButton(conn redis.Conn, node Node) error {
	field := getNodeField(node.NodeName)
	data, err := json.Marshal(node)
	if err != nil {
		return err
	}

	_, err = conn.Do("SETEX", field, NODETIMEOUT, data)

	return err
}

func GetNodeButton(conn redis.Conn, node *Node, name string) error {
	field := getNodeField(name)
	data, err := redis.Bytes(conn.Do("GET", field))
	if err != nil {
		return err
	}

	return json.Unmarshal(data, node)
}

func GetAllNodeButton(conn redis.Conn, nodeList *[]Node) error {
	field := NODE_NAMESPACE + ":" + NODE_GROUP + "*"
	values, err := redis.Strings(conn.Do("KEYS", field))
	if err != nil {
		return err
	}

	for _, v := range values {
		var node Node
		data, err := redis.Bytes(conn.Do("GET", v))
		if err != nil {
			continue
		}
		json.Unmarshal(data, &node)
		*nodeList = append(*nodeList, node)
	}

	return nil
}

func SetBtnCmdMap(conn redis.Conn, group ButtonCmdGroup) error {
	field := getBtnCmdField(group.NodeName)
	_, err := conn.Do("HSET", field, group.ButtonName, group.CmdGroupName)
	return err
}

func GetBtnCmdMap(conn redis.Conn, node string, btn string) (string, error) {
	field := getBtnCmdField(node)
	return redis.String(conn.Do("HGET", field, btn))
}


