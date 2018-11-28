package main

import "fmt"

func main() {

	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["China"] = "Beijing"
	countryCapitalMap["India"] = "新德里"
	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是：", countryCapitalMap[country])
	}

	captial, hasValue := countryCapitalMap["China"] /*如果确定是真实的,则存在,否则不存在 */

	fmt.Println(captial)
	fmt.Println(hasValue)
	if hasValue {
		fmt.Println("中国首都是", captial)
	} else {
		fmt.Println("中国首都不存在")
	}

}
