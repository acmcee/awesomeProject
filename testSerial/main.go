package main

import (
	"encoding/json"
	"fmt"
)

/*
结构体序列化
map序列化
*/

type Server struct {
	ServerName string `json:"name"`
	ServerIp   string `json:"ip"`
	ServerPort int    `json:"port"`
}

func SerializeStruct() {
	server := new(Server)
	server.ServerName = "MySQL"
	server.ServerIp = "1.1.1.1"
	server.ServerPort = 3306

	b, err := json.Marshal(server) // 序列化成json 字节的数组
	if err != nil {
		fmt.Println("Marshal error, ", err.Error())
		return
	}
	fmt.Println("JSON", string(b)) // 将json 字节数组转成string
}

func SerializeMap() {
	server := make(map[string]interface{}) // value 可以是各种类型，直接用接口
	server["ServerName"] = "MySQL"
	server["ServerIp"] = "1.1.1.1"
	server["ServerPort"] = 3306

	b, err := json.Marshal(server) // 序列化成json 字节的数组
	if err != nil {
		fmt.Println("Marshal error, ", err.Error())
		return
	}
	fmt.Println("JSON", string(b)) // 将json 字节数组转成string
}

func DeSerializeStruct() {
	jsonString := `{"name":"MySQL","ip":"1.1.1.1","port":3306}`
	server := new(Server)
	err := json.Unmarshal([]byte(jsonString), server)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(*server)
}

func DeSerializeMap() {
	jsonString := `{"name":"MySQL","ip":"1.1.1.1","port":3306}`
	server := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &server)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(server)

}

func main() {
	// 序列化struct
	//SerializeStruct()

	// 序列化map
	//SerializeMap()

	// 反序列化

	//DeSerializeStruct()

	DeSerializeMap()

}
