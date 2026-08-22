package main

import (
	"context"
	"fmt"
	simplesql "study/simpleSql"
	"study/simpleSql/sqlfunc"
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
	fmt.Println("succed")
}
