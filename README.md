# zdpgo_yaml_to_json
Go实现yaml格式转json格式

## 版本历史
- 版本0.1.0 2022年4月20日 读取yaml文件转换为JSON字符串

## 使用示例
### 读取yaml文件转换为JSON字符串
```go
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
```