package scheduler

import (
	"Ravigation/ravigation/storage"
	"errors"
	"github.com/RyanCarrier/dijkstra"
)

//func Navigation()  {
//	graph, err := dijkstra.Import("graph/g1")
//	best, err := graph.Shortest(0,6)
//	if err!=nil{
//		logrus.Debug(err)
//	}
//	logrus.Info("Shortest distance ", best.Distance, " following path ", best.Path)
//}

func Navigation(index int, cg *storage.CmdGroup) error {
	graph, err := dijkstra.Import("graph/g1")
	if err != nil {
		return err
	}
	startIndex := index
	destIndex := 0

	for cmdIndex, cmd := range cg.CmdList {
		if cmd.GetType() != storage.CMDTYPE_MOVETO {
			continue
		}
		destIndex = cmd.(*storage.MoveToCmd).Index
		best, err := graph.Shortest(startIndex, destIndex)
		if err != nil {
			return errors.New("graph shortest error")
		}

		for _, i := range best.Path {
			var p storage.Point
			storage.GetPoint(conn, &p, i)
			cg.CmdList[cmdIndex].(*storage.MoveToCmd).PointList = append(cg.CmdList[cmdIndex].(*storage.MoveToCmd).PointList, p)
		}

		startIndex = destIndex
	}

	return nil
}
