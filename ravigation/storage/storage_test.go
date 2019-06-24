package storage

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"testing"
)


func TestSetAgvState(t *testing.T) {
	point := Point{
		X:1.1,
		Y:2.2,
		Z:3.3,
	}
	state := AgvState{
		Name: "agv1",
		Timestamp: 123456,
		Point: point,
		ActionStatus:ActionStatusFinished,
	}

	conn := GetConn()
	defer conn.Close()
	err := SetAgvState(conn, state)
	if err != nil {
		t.Error("set agv state error", err)
	}
}

func TestGetAgvState(t *testing.T) {
	var state AgvState
	conn := GetConn()
	defer conn.Close()

	err := GetAgvState(conn, &state, "agv1")
	if err != nil || state.Name != "agv1" {
		t.Error("get agv state error", err)
	}
}

func TestDeleteAgvState(t *testing.T) {
	conn := GetConn()
	defer conn.Close()
	err := DeleteAgvState(conn, "agv1")
	if err != nil {
		t.Error("delete agv state error", err)
	}
	res, err := redis.Int(conn.Do("EXISTS", getAgvStateField("agv1")))
	if err != nil || res != 0 {
		t.Error("delete agv state error: EXISTS is not 0")
	}
}

func TestGetAllAgvState(t *testing.T) {
	point := Point{
		X:1.1,
		Y:2.2,
		Z:3.3,
	}
	state := AgvState{
		Name: "agv1",
		Timestamp: 123456,
		Point: point,
		ActionStatus:ActionStatusFinished,
	}

	conn := GetConn()
	defer conn.Close()
	err := SetAgvState(conn, state)
	if err != nil {
		t.Error("set agv state error")
	}

	stateList := []AgvState{}
	err = GetAllAgvState(conn, &stateList)
	if cap(stateList) == 0 {
		t.Error("TestGetAllAgvState has error")
	}
}

func TestGetAllAgvStateWithActionStatus(t *testing.T) {
	point := Point{
		X:1.1,
		Y:2.2,
		Z:3.3,
	}
	state := AgvState{
		Name: "agv1",
		Timestamp: 123456,
		Point: point,
		ActionStatus:ActionStatusStopped,
	}

	conn := GetConn()
	defer conn.Close()
	err := SetAgvState(conn, state)
	if err != nil {
		t.Error("set agv state error")
	}

	agvList := []AgvState{}
	err = GetAllAgvStateWithActionStatus(conn, ActionStatusStopped, &agvList)
	if cap(agvList) == 0 {
		t.Error("TestGetAllAgvStateWithActionStatus has error")
	}
}

func TestAddPoint(t *testing.T) {
	point := NewPoint(1.1,3.4)
	conn := GetConn()
	defer conn.Close()
	err := AddPoint(conn, 1, point)
	if err != nil {
		t.Error("Add point error")
	}
}

func TestGetPoint(t *testing.T) {
	conn := GetConn()
	defer conn.Close()
	point := NewPoint(0,0)
	if err := GetPoint(conn, &point, 1); err != nil {
		t.Error("get point error")
	}
	if point.X != 1.1 {
		t.Error("get point error")
	}
}

func TestDeletePoint(t *testing.T) {
	conn := GetConn()
	defer conn.Close()
	count, err := DeletePoint(conn, 1)
	if err != nil || count == 0 {
		t.Error("delete point error", err)
	}
}

func TestSetNodeButton(t *testing.T) {
	conn := GetConn()
	defer conn.Close()
	button := Button{
		ButtonName:"btn1",
		Status:"down",
	}
	node := Node{
		NodeName:"node1",
		ButtonList:[]Button{button},
	}
	err := SetNodeButton(conn, node)
	if err != nil {
		t.Error("set node error")
	}
}

func TestGetNodeButton(t *testing.T) {
	conn := GetConn()
	defer conn.Close()
	var node Node
	err := GetNodeButton(conn, &node, "node1")
	if err != nil || cap(node.ButtonList) == 0 ||node.ButtonList[0].ButtonName != "btn1" {
		t.Error("get node error", err)
	}
}

func TestGetAllNodeButton(t *testing.T) {
	conn := GetConn()
	defer conn.Close()
	var nodeList []Node
	err := GetAllNodeButton(conn, &nodeList)
	if err != nil || cap(nodeList) == 0 || nodeList[0].ButtonList[0].ButtonName != "btn1" {
		t.Error("get node error", err)
	}
}

func TestCreateConnectInfo(t *testing.T) {
	conn := GetConn()
	defer conn.Close()
	connectInfo := ConnectInfo{
		Device:"agv",
		Name:"agv1",
		Host: "122.122.122.122",
		Port: 0,
	}

	err := CreateConnectInfo(conn, connectInfo)
	if err != nil {
		t.Error("create connect info error")
	}
}

func TestGetConnectInfo(t *testing.T) {
	conn := GetConn()
	defer conn.Close()
	var info ConnectInfo
	GetConnectInfo(conn, &info, "agv", "agv1")
	if info.Host != "122.122.122.122" {
		t.Error("get connect info error")
	}
}

func TestGetAllConnectInfo(t *testing.T) {
	conn := GetConn()
	defer conn.Close()
	var info []ConnectInfo
	GetAllConnectInfo(conn, &info)
	if info[0].Host != "122.122.122.122" {
		t.Error("get connect info error")
	}
}

func TestDeleteConnectInfo(t *testing.T) {
	conn := GetConn()
	defer conn.Close()
	count, err := DeleteConnectInfo(conn, "agv", "agv1")
	if err != nil || count == 0 {
		t.Error("delete connect info error")
	}
}

func TestAddCmdGroup(t *testing.T) {
	conn := GetConn()
	defer conn.Close()

	cmdstr := &MoveToCmd{
		Type:CMDTYPE_MOVETO,
		Index: 1,
	}
	cmdstr2 := &MoveToCmd{
		Type:CMDTYPE_MOVETO,
		Index: 2,
	}

	cmdList := []BaseCmd{
		cmdstr,
		cmdstr2,
	}

	cg := CmdGroup{
		CgName:"cgname1",
		CmdList:cmdList,
	}
	err := AddCmdGroup(conn, cg)
	if err != nil {
		t.Error("add cmdgroup error", err)
	}
}

func TestGetCmdGroup(t *testing.T) {
	conn := GetConn()
	defer conn.Close()
	var cg CmdGroup
	GetCmdGroup(conn, "cgname1", &cg)
	if cg.CmdList[0].GetType() != CMDTYPE_MOVETO {
		t.Error("get cmd group error")
	}
}

func TestRPushCmdGroupToWaitList(t *testing.T) {
	node := NodeCmdGroup{
		NodeName:"node1",
		BtnName:"btn1",
		Timestamp:123456,
		CmdGroupName:"cg1",
	}
	conn := GetConn()
	defer conn.Close()

	if err := RPushCmdGroupToWaitList(conn, &node); err!= nil {
		t.Error("rpush cmdgroup error")
	}
}

func TestLPopCmdGroupFromWaitList(t *testing.T) {
	conn := GetConn()
	defer conn.Close()

	var node NodeCmdGroup
	err := LPopCmdGroupFromWaitList(conn, &node)
	if err != nil || node.CmdGroupName != "cg1" {
		t.Error("TestLPopCmdGroupFromWaitList error")
	}
}

func TestLPushCmdGroupToWaitList(t *testing.T) {
	node := NodeCmdGroup{
		NodeName:"node1",
		BtnName:"btn1",
		Timestamp:123456,
		CmdGroupName:"cg1",
	}
	conn := GetConn()
	defer conn.Close()

	if err := LPushCmdGroupToWaitList(conn, &node); err!= nil {
		t.Error("rpush cmdgroup error")
	}
}

func TestGetAllCmdGroupFromWaitList(t *testing.T) {
	conn := GetConn()
	defer conn.Close()
	ncgList := []NodeCmdGroup{}
	err := GetAllCmdGroupFromWaitList(conn, &ncgList)
	if err != nil || ncgList[0].CmdGroupName != "cg1" {
		t.Error("TestGetAllCmdGroupFromWaitList error")
	}
}

func TestDeleteCmdGroupFromWaitList(t *testing.T) {
	conn := GetConn()
	defer conn.Close()
	err := DeleteCmdGroupFromWaitList(conn, 0)
	if err != nil {
		t.Error("DeleteCmdGroupFromWaitList error")
	}
}

func TestSetCmdToRunList(t *testing.T) {
	node := NodeCmdGroup{
		NodeName:"node1",
		BtnName:"btn1",
		Timestamp:123456,
		CmdGroupName:"cg1",
	}
	conn := GetConn()
	defer conn.Close()
	err := SetCmdToRunList(conn, "agv1", node)
	if err != nil {
		t.Error("TestSetCmdToRunList error ")
	}
}

func TestGetCmdFromRunList(t *testing.T) {
	var node NodeCmdGroup
	conn := GetConn()
	defer conn.Close()
	err := GetCmdFromRunList(conn, "agv1", &node)
	if err != nil || node.CmdGroupName != "cg1" {
		t.Error("TestGetCmdFromRunList", err)
	}
}

func TestDeleteCmdFromRunList(t *testing.T) {
	conn := GetConn()
	defer conn.Close()
	res, err:= DeleteCmdFromRunList(conn, "agv1")
	if res != 1 || err !=nil {
		t.Error("TestDeleteCmdFromRunList error", err)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Before ====================")

	code := m.Run()
	fmt.Println("End ====================")
	os.Exit(code)
}







