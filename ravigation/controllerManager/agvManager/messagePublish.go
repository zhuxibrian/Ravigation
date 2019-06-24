package agvManager

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"Ravigation/ravigation/storage"
	"Ravigation/ravigation/utils"
)

func ConnectAgv(conn redis.Conn, info storage.ConnectInfo)  {
	var agv storage.AgvState
	err := storage.GetAgvState(conn, &agv, info.Name)
	if err == nil {
		return
	}

	data, err :=json.Marshal(info);
	if err != nil {
		return
	}
	utils.Publish(TOPIC_RAVI_AGV_CONNECT+ "/" + info.Name, 0, false, string(data))
}

func CommandAgv(agv storage.AgvState, group storage.CmdGroup) {
	data, err := json.Marshal(group)
	if err != nil {
		return
	}

	utils.Publish(TOPIC_RAVI_AGV_COMMAND + "/" + agv.Name, 2, false, string(data))
}


