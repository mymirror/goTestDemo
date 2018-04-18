package main

import (
	//"userHttp"
	//"userSqlite"
	//"userHttp"
	"fmt"
	"reflect"
)

func main (){

	//userHttp.WebService()

	//data := userSqlite.Usegr{
	//	Name:"100",
	//}
	//userSqlite.InsertData(data)

	//userSqlite.QueryData()

	u:= user{"hello","12"}

	v:= reflect.TypeOf(u)

	m:="21"
	x:= reflect.ValueOf(&m)
	x.Elem().SetString("12")

	fmt.Println(v)
	fmt.Println(x)

}

//反射

type user struct {
	Name string
	Age  string
}


