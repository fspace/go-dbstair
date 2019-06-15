package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	fmt.Println(sql.Drivers())

	// 建议db 作为全局对象被共享使用 而且不要经常open和close 这样影响性能 应该作为长存活的对象使用 生命周期和程序生命周期一致为好
	db, err := sql.Open("mysql",
		//"user:password@tcp(127.0.0.1:3306)/hello")
		"root:@tcp(127.0.0.1:3306)/my_blog")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		// do something here
	}

}
