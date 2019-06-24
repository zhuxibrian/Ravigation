package service

import "Ravigation/ravigation/storage"

func AddPoint(index int, point storage.Point) error {
	conn := storage.GetConn()
	defer conn.Close()
	return storage.AddPoint(conn, index, point)
}