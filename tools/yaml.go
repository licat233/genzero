package tools

import (
	"os"
	"sort"

	"gopkg.in/yaml.v3"
)

func MergeYamlContent(yamlPath string, newContent []byte) error {
	if TrimSpace(string(newContent)) == "" {
		return nil
	}

	// 读取第一个YAML文件
	yaml1, err := os.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	// 将两个 YAML 文档解析成 map[string]interface{} 类型
	var data1 map[string]interface{}
	var data2 map[string]interface{}
	if err := yaml.Unmarshal(yaml1, &data1); err != nil {
		return err
	}
	if err := yaml.Unmarshal(newContent, &data2); err != nil {
		return err
	}

	// 合并两个 map
	result := merge(data1, data2)

	result = SortYaml(result)

	// 将合并后的结果转换成 YAML 格式
	b, err := yaml.Marshal(&result)
	if err != nil {
		return err
	}

	err = os.WriteFile(yamlPath, b, 0644)
	if err != nil {
		return err
	}

	// // 输出合并后的 YAML 文档
	// fmt.Println(string(b))
	return nil
}

// 合并两个 map
func merge(data1, data2 map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	// 将 data1 的数据复制到 result 中
	for k, v := range data1 {
		result[k] = v
	}

	// 将 data2 的数据合并到 result 中
	for k, v := range data2 {
		if _, ok := result[k]; ok {
			// 如果 key 已存在，执行深度合并
			result[k] = mergeMaps(result[k], v)
		} else {
			// 如果 key 不存在，直接复制
			result[k] = v
		}
	}

	return result
}

// 深度合并两个 map
func mergeMaps(map1, map2 interface{}) interface{} {
	switch m1 := map1.(type) {
	case map[interface{}]interface{}:
		m2 := map2.(map[interface{}]interface{})
		for k, v := range m2 {
			if mv, ok := m1[k]; ok {
				m1[k] = mergeMaps(mv, v)
			} else {
				m1[k] = v
			}
		}
		return m1
	case map[string]interface{}:
		m2 := map2.(map[string]interface{})
		for k, v := range m2 {
			if mv, ok := m1[k]; ok {
				m1[k] = mergeMaps(mv, v)
			} else {
				m1[k] = v
			}
		}
		return m1
	default:
		//如果map1已经存在，则不更改其值，依旧使用原来的值
		return map1
	}
}

func SortYaml(yamlContent map[string]interface{}) map[string]interface{} {
	// 对 map 的 key 进行排序
	keys := make([]string, 0, len(yamlContent))
	for k := range yamlContent {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 创建新的有序的 map
	orderedMap := make(map[string]interface{})
	for _, k := range keys {
		orderedMap[k] = yamlContent[k]
	}
	return orderedMap
}

func SortYamlContent(yamlContent string) (string, error) {
	// 解析 YAML 内容成 map 类型
	var data map[string]interface{}
	err := yaml.Unmarshal([]byte(yamlContent), &data)
	if err != nil {
		return "", err
	}

	// 对 map 的 key 进行排序
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 创建新的有序的 map
	orderedMap := make(map[string]interface{})
	for _, k := range keys {
		orderedMap[k] = data[k]
	}

	// 转换成 YAML 格式输出
	out, err := yaml.Marshal(orderedMap)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
