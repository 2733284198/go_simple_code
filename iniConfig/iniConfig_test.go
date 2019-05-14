package iniConfig

import (
	"fmt"
	"go_simple_code"
	"io/ioutil"
	"testing"
)

type ServerConfig struct {
	Ip string	`ini:"ip"`
	Port int	`ini:"port"`
}

type MysqlConfig struct {
	UserName string	`ini:"username"`
	Passwd string	`ini:"passwd"`
	DataBase string `ini:"database"`
	Host string	`ini:"host"`
	Port int	`ini:"port"`
}

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf MysqlConfig	`ini:"mysql"`
}

func TestIniConfig(t *testing.T) {
	fmt.Println("hello")
	data, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Error("open file error:",err)
	}
	var conf Config
	err = go_simple_code.Unmarshal(data, &conf)
	if err != nil {
		t.Fatalf("unmarshal failed, err:%v",err)
	}
	t.Log("unmarshall success")
	//t.Logf("unmarshal success, conf:%v\n",conf)
}

