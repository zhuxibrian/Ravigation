package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"Ravigation/ravigation/storage"
	"strconv"
)

const url  = "http://127.0.0.1:13579"

func PostConnectInfo(info storage.ConnectInfo)  {
	bytesData, err := json.Marshal(info)
	if err != nil {
		fmt.Println(err.Error() )
		return
	}
	respBytes, _ := post(bytesData, "/device/connectInfo")

	fmt.Printf("%s\n", respBytes)

}

func GetConnectInfo(device string, name string) {
	path := "/device/" + device + "/name/" + name
	respBytes, _ := http.Get(url + path)

	fmt.Printf("%s\n", respBytes)
}

func PostPoint(point storage.Point, index int) {
	bytesData, err := json.Marshal(point)
	if err != nil {
		fmt.Println(err.Error() )
		return
	}
	respBytes, _ := post(bytesData, "/point/" + strconv.Itoa(index))

	fmt.Printf("%s\n", respBytes)
}

func PostButtonCmdGroup(group storage.ButtonCmdGroup) {
	bytesData, err := json.Marshal(group)
	if err != nil {
		fmt.Println(err.Error() )
		return
	}
	respBytes, _ := post(bytesData, "/node/btncmdmap")

	fmt.Printf("%s\n", respBytes)
}

func PostCmdGroup(cmdGroup string) {
	respBytes, _ := post([]byte(cmdGroup), "/cmdgroup/cmdgroup")

	fmt.Printf("%s\n", respBytes)
}



func post(bytesData []byte, path string) ([]byte, error) {

	resp, err := http.Post(url+path,
		"application/json;charset=UTF-8",
		bytes.NewReader(bytesData))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return respBytes, nil
}

