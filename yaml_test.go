package zdpgo_yaml_to_json

import "testing"

func getYamlToJSON() *YamlToJson {
	y := New()
	return y
}

func TestYamlToJson_YamlFileToJSONStr(t *testing.T) {
	y := getYamlToJSON()
	jsonStr, err := y.YamlFileToJSONStr("test/test.yaml")
	if err != nil {
		t.Error(err)
	}
	t.Log(jsonStr)
}
