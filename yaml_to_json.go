package zdpgo_yaml_to_json

import (
	"github.com/zhangdapeng520/zdpgo_json"
	"github.com/zhangdapeng520/zdpgo_yaml"
)

type YamlToJson struct {
	Yaml *zdpgo_yaml.Yaml
	Json *zdpgo_json.Json
}

func New() *YamlToJson {
	y := YamlToJson{}
	y.Yaml = zdpgo_yaml.New()
	y.Json = zdpgo_json.New()
	return &y
}
