package service

import "Ravigation/ravigation/storage"

func AddConnectInfo(info storage.ConnectInfo) error {
	conn := storage.GetConn()
	defer conn.Close()
	return storage.CreateConnectInfo(conn, info)
}

func DeleteConnectInfo(device string, name string) (int, error) {
	conn := storage.GetConn()
	defer conn.Close()
	return storage.DeleteConnectInfo(conn, device, name)
}

func GetAllConnectInfo(infoList *[]storage.ConnectInfo) error {
	conn := storage.GetConn()
	defer conn.Close()
	return storage.GetAllConnectInfo(conn, infoList)
}

func GetConnectInfo(info *storage.ConnectInfo, device string, name string) error {
	conn := storage.GetConn()
	defer conn.Close()
	return storage.GetConnectInfo(conn, info, device, name)
}