package main

import "fmt"

func funcMap() {
	var countryCapitalMap map[string]string
	/* 创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* 使用 key 输出 map 值 */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* 查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if ok {
		fmt.Println("Capital of United States is", capital)
	} else {
		fmt.Println("Capital of United States is not present")
	}

	/* 删除元素 */
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后 map:")
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* map 的长度 */
	fmt.Println("Length of map is:", len(countryCapitalMap))

	/* map 是引用类型 */
	var newMap map[string]string = countryCapitalMap
	newMap["United Kingdom"] = "London"
	fmt.Println("New map after adding UK:", newMap)
	fmt.Println("Original map:", countryCapitalMap)
	/* 注意：对 newMap 的操作并不影响 countryCapitalMap */
	fmt.Println("Length of new map is:", len(newMap))
	fmt.Println("Is newMap the same as countryCapitalMap?", &newMap == &countryCapitalMap)
	/* 注意：newMap 和 countryCapitalMap 并不是同一个 map */
	fmt.Println("Country capital of United Kingdom is:", newMap["United Kingdom"])
	fmt.Println("Country capital of United Kingdom is:", countryCapitalMap["United Kingdom"])
}

// 同一个目录下只能有一个 main 函数
// func main() {
// 	funcMap()
// }
