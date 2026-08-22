package main

import (
	"context"
	"fmt"
	simplesql "study/simpleSql"
	"study/simpleSql/sqlfunc"
	//"github.com/k0kubun/pp/v3"
)

func main() {
	ctx := context.Background()
	conn, err := simplesql.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = sqlfunc.CreateTable(ctx, conn)
	if err != nil {
		panic(err)
	}
	_, err = sqlfunc.ListPages(3, ctx, conn)
	if err != nil{
		panic(err)
	}
	fmt.Println("succed")
}
