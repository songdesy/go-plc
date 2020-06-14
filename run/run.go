package main

import (
	"../service"
	"../tcp/myhttp"
	"../tcp/server"
	"fmt"
)
func main() {
	//02100000000A140B5F7FC1001ECF954498ED7142F6DC29439BB70AF225
	//9714
	src := []string{"190808001", "2019221", "1223.42", "123.43", "311.43"}
	data := service.SendDataHandler(src, "02100000000A14", 2, true)
	fmt.Println(data)
	go server.Start()
	myhttp.Start()

}
