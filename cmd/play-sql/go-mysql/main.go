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

	retrievingResultSet(db)
	preparingQueries(db)
	singleRowQuery(db)
	queryRow2(db)

	modifyData1(db)
}

func retrievingResultSet(db *sql.DB) {
	var (
		id   int
		name string
	)
	rows, err := db.Query("select id, username from tbl_user where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
func preparingQueries(db *sql.DB) {
	var (
		id   int
		name string
	)

	stmt, err := db.Prepare("select id ,username from tbl_user where id = ?  LIMIT 10")
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("user data :", id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

func singleRowQuery(db *sql.DB) {
	var name string
	err := db.QueryRow("select  username from tbl_user where id= ?", 1).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}
func queryRow2(db *sql.DB) {

	stmt, err := db.Prepare("select  username from tbl_user where  id= ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var name string
	// 当QueryRow 方法查询不到结果时 会出现sql.ErrNoRows错误 但此错误被推迟到了Scan方法上
	err = stmt.QueryRow(1).Scan(&name)
	if err != nil {
		//if err == sql.ErrNoRows {
		//	// there were no rows, but otherwise no error occurred
		//} else {
		//	log.Fatal(err)
		//}
		log.Fatal(err)
	}

	fmt.Println(name)
}

// http://go-database-sql.org/modifying.html
func modifyData1(db *sql.DB) {
	// INSERT INTO `my_blog`.`tbl_user` (
	//  `id`,
	//  `username`,
	//  `password`,
	//  `email`,
	//  `profile`
	//)
	//VALUES
	//  (
	//    'id',
	//    'username',
	//    'password',
	//    'email',
	//    'profile'
	//  );
	stmt, err := db.Prepare("INSERT INTO tbl_user(username) values (?)")
	if err != nil {
		log.Fatal(err)
	}
	rslt, err := stmt.Exec("Yiqing")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := rslt.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := rslt.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d , affected = %d", lastId, rowCnt)
}
