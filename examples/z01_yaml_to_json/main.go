package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_yaml_to_json"
)

func main() {
	y := zdpgo_yaml_to_json.New()
	jsonStr, err := y.YamlFileToJSONStr("test/test.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonStr)
}
