package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var config *viper.Viper

func loadConfigFromYaml () (err error)  {
	config = viper.New()

	//设置配置文件的名字
	config.SetConfigName("ravigation")

	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH
	config.AddConfigPath("./config/")
	config.AddConfigPath("../config/")



	//设置配置文件类型
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		return err
	}

	//if err := config.Unmarshal(&raviConfig); err!=nil{
	//	return err
	//}


	return nil
}

func Config() *viper.Viper  {
	return config
}
func init() {
	if err := loadConfigFromYaml(); err!=nil {
		logrus.Error("load config from yaml error: %s", err.Error())
	}
}


