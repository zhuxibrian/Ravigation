package cmd

//import (
//	"encoding/json"
//	"ravicommon"
//)
//
//func cmdGroupCheck(command string) (string, error) {
//	cmdGroup := ravicommon.CmdGroup{}
//	err := json.Unmarshal([]byte(command), &cmdGroup)
//	if err != nil {
//		return "", err
//	}
//	return cmdGroup.Name, nil
//
//}
//
//func ctlCmdCheck(command string) (error) {
//	ctlcmd := ravicommon.CtlCmd{}
//	err := json.Unmarshal([]byte(command), &ctlcmd)
//	if err != nil {
//		return err
//	}
//	return nil
//}
