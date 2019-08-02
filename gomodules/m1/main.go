package main

import (
	"github.com/labstack/echo"
	//"hello/pb"

	"hello/pb"
	"hi/m2"
)

//func main() {
//	pb.Main()
//}

func main() {
	e := echo.New()
	e.GET("/", pb.Hi)
	e.Logger.Fatal(e.Start(":1323"))

	m2.T1()
}
