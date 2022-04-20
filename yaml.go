package zdpgo_yaml_to_json

// YamlFileToJSONStr 将yaml文件转换为JSON格式的字符串
func (y *YamlToJson) YamlFileToJSONStr(yamlFilePath string) (string, error) {
	var (
		data   map[string]interface{}
		err    error
		result string
	)
	err = y.Yaml.ReadConfig(yamlFilePath, &data)
	if err != nil {
		return "", err
	}
	result, err = y.Json.Dumps(data)
	if err != nil {
		return "", err
	}
	return result, nil
}
