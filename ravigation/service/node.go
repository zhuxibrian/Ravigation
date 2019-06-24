package service

import "Ravigation/ravigation/storage"

func AddBtnCmdMap(group storage.ButtonCmdGroup) error {
	conn := storage.GetConn()
	defer conn.Close()
	return storage.SetBtnCmdMap(conn, group)
}