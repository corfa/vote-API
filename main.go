package main

import "Voting-API/app"

func main() {

	application := &app.App{}
	application.Initialize()
	application.Run(":3000")

	//conn, err := redis.Dial("tcp", "localhost:5000", redis.DialDatabase(5))
	//if err != nil {
	//	// обработка ошибок
	//}
	//_, err = conn.Do("SET", "mykey", "ivan")
	//if err != nil {
	//	// обработка ошибок
	//}
	//res, err := conn.Do("GET", "mykey")
	//if err != nil {
	//	// обработка ошибок
	//}
	//
	//value, err := redis.String(res, err)
	//if err != nil {
	//	// обработка ошибок
	//}
	//
	//fmt.Println(value)

}
