package storage

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"strings"
)

type AgvState struct {
	Name string
	Timestamp int64
	Point
	ActionStatus int32
}


const (
	ActionStatusWaitingForStart = iota
	ActionStatusRunning
	ActionStatusFinished
	ActionStatusPaused
	ActionStatusStopped
	ActionStatusError
)

const (
	AGV_NAMESPACE= "Agv"

	STATE_GROUP = "State"
)

const AGVTIMEOUT = 20

func getAgvStateField(name string) string {
	return AGV_NAMESPACE + ":" + STATE_GROUP + ":" + name
}

/**
设置agv状态
 */
func SetAgvState(conn redis.Conn, state AgvState) error {
	field := getAgvStateField(state.Name)
	data, err := json.Marshal(&state)
	if err != nil {
		return err
	}
	_, err = conn.Do("SETEX", field, AGVTIMEOUT, string(data))
	return err
}

/**
获得指定agv状态
 */
func GetAgvState(conn redis.Conn, state *AgvState, name string) error {
	field := getAgvStateField(name)
	data, err := redis.String(conn.Do("GET", field))
	if err!= nil {
		return err
	}

	return json.Unmarshal([]byte(data), state)
}

/**
删除指定agv状态
 */
func DeleteAgvState(conn redis.Conn, name string) error {
	field := getAgvStateField(name)
	_, err := conn.Do("DEL", field)
	return err
}

/**
获得全部agv状态
 */
func GetAllAgvState(conn redis.Conn, stateList *[]AgvState) error {

	values, err := redis.Strings(conn.Do("KEYS", AGV_NAMESPACE + ":" + STATE_GROUP + "*"))
	if err != nil {
		return err
	}

	for _, v := range values {
		var agv AgvState
		s := strings.Split(v, ":")
		if err := GetAgvState(conn, &agv, s[2]); err != nil {
			return err
		}
		*stateList = append(*stateList, agv)
	}

	return nil
}

/**
获得全部指定状态的agv
 */
func GetAllAgvStateWithActionStatus(conn redis.Conn, actionStatus int32, stateList *[]AgvState) error {

	allStates := make([]AgvState, 10)
	err := GetAllAgvState(conn, &allStates)
	if err != nil {
		return err
	}

	for _, v := range allStates{
		if v.ActionStatus != actionStatus {
			continue
		}
		*stateList = append(*stateList, v)
	}

	return nil
}
