package main

import "fmt"

func main() {
	var mapa map[string]string
	mapa = make(map[string]string)
	mapa["name"] = "葛新"
	fmt.Println(mapa)

	var mapb map[string]string = make(map[string]string)
	mapb["age"] = "18"
	fmt.Println(mapb)

	mapc := map[string]interface{}{"name": "葛新", "age": 29, "gender": "male", "hobbies": []string{"羽毛球", "跑步", "骑行"}}
	fmt.Println(mapc)

	birthday, ok := mapc["birthday"]
	if ok {
		fmt.Println(birthday)
	} else {
		fmt.Println("birthday不存在")
	}
}
