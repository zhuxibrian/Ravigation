package service

import "Ravigation/ravigation/storage"

func AddCmdGroup(group storage.CmdGroup) error {
	conn := storage.GetConn()
	defer conn.Close()

	return storage.AddCmdGroup(conn, group)
}