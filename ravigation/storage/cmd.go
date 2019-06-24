package storage

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
)

const (
	CMDTYPE_MOVETO = "MoveToCmd"
)
const (
	CMD_NAMESPACE= "CMD"

	CMD_GROUP = "CmdGroup"
	WAIT_GROUP = "WaitGroup"
	RUN_GROUP = "RunGroup"
)

type CmdGroup struct {
	CgName string
	CmdList []BaseCmd
}

type BaseCmd interface {
	GetType() string
}

type MoveToCmd struct {
	Type string
	Index int
	PointList []Point
}

func (cmd *MoveToCmd) GetType() string  {
	return cmd.Type
}

func getCmdGroupField() string {
	return CMD_NAMESPACE + ":" + CMD_GROUP
}

func AddCmdGroup(conn redis.Conn, group CmdGroup) error {
	field := getCmdGroupField()
	data, err := json.Marshal(group)
	if err != nil {
		return err
	}
	_, err = conn.Do("HSET", field, group.CgName, data)
	return err
}

func DeleteCmdGroup(conn redis.Conn, cgName string) error {
	field := getCmdGroupField()
	_, err := conn.Do("HDEL", field, cgName)
	return err
}

func GetCmdGroup(conn redis.Conn, cgName string, cg *CmdGroup) error {
	field := getCmdGroupField()
	data, err := redis.Bytes(conn.Do("HGET", field, cgName))
	if err != nil {
		return err
	}

	return cg.UnmarshalJSON(data)

}

func (cg *CmdGroup) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	cg.CgName = string(bytes.Trim(*objMap["CgName"], "\""))
	var rawMessagesForColoredThings []*json.RawMessage
	err = json.Unmarshal(*objMap["CmdList"], &rawMessagesForColoredThings)
	if err != nil {
		return err
	}

	// Let's add a place to store our de-serialized Plant and Animal structs
	cg.CmdList = make([]BaseCmd, len(rawMessagesForColoredThings))

	var m map[string]interface{}
	for index, rawMessage := range rawMessagesForColoredThings {
		err = json.Unmarshal(*rawMessage, &m)
		if err != nil {
			return err
		}

		// Depending on the type, we can run json.Unmarshal again on the same byte slice
		// But this time, we'll pass in the appropriate struct instead of a map
		if m["Type"].(string) == CMDTYPE_MOVETO {
			var p MoveToCmd
			err := json.Unmarshal(*rawMessage, &p)
			if err != nil {
				return err
			}
			// After creating our struct, we should save it
			cg.CmdList[index] = &p
		} else {
			return errors.New("Unsupported type found!")
		}
	}

	// That's it!  We made it the whole way with no errors, so we can return `nil`
	return nil

}





