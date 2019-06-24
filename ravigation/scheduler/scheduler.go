package scheduler

import (
	"errors"
	"Ravigation/ravigation/controllerManager/agvManager"
	"Ravigation/ravigation/storage"
	"runtime"
	"time"
)

var conn = storage.GetConn()

func Start() {
	go scheduler()
}

func scheduler() {
	for {
		var ncg storage.NodeCmdGroup
		if err := storage.LPopCmdGroupFromWaitList(conn, &ncg); err != nil {
			//logrus.Error("scheduler lpop cmdgroup from wait list error:", err)
			continue
		}

		var cg storage.CmdGroup;
		if err := storage.GetCmdGroup(conn, ncg.CmdGroupName, &cg); err != nil {
			continue
		}

		var agv storage.AgvState
		if err := GetIdleAgv(&agv); err != nil {
			continue
		}

		agvPoint := storage.Point{
			X: agv.X,
			Y: agv.Y,
			Z: agv.Z,
		}

		index, err := storage.GetPointIndex(conn, agvPoint)
		if err != nil {
			continue
		}

		Navigation(index, &cg)

		agvManager.CommandAgv(agv, cg)
		storage.SetCmdToRunList(conn, agv.Name, ncg)
		agvName := agv.Name
		for agv.ActionStatus == storage.ActionStatusFinished {
			storage.GetAgvState(conn, &agv, agvName)
			time.Sleep(time.Duration(1) * time.Second)
			runtime.Gosched()
		}
	}
}

func GetIdleAgv(agv *storage.AgvState) error {

	var agvList []storage.AgvState

	if err := storage.GetAllAgvStateWithActionStatus(conn, storage.ActionStatusFinished, &agvList); err != nil {
		return err
	}

	//TODO 根据条件筛选空闲agv执行命令

	if len(agvList) == 0 {
		return errors.New("there is no idle agv!")
	}

	*agv = agvList[0]

	return nil
}





