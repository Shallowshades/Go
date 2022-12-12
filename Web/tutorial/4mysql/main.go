package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" //调用init(), 但不使用api
)

func main() {

	//sql.Open用来打开一个数据库
	//第一个参数 string 			数据库驱动的名称
	//第二个参数 dataSourceName		数据库源
	//由多GO程确保并发安全、自己维持自己的线程池，仅调用一次即可，且没必要关闭
	db, err := sql.Open("mysql", "root:123456@(192.168.29.1:3306)/itcast?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	//Ping验证数据库连接是否有效，必要时建立连接
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	//删除一个数据表
	{
		//Exec 执行一条sql，不返回任何行数据（不能用于查询语句）
		//第一个参数 string 	SQL语句
		//第二个参数 ...any		替换sql语句中的占位符
		//返回SQL执行的结果
		//LastInsertId() 	最后一次插入数据的id（一般是auto increment字段）
		//RowsAffected()	受到影响的行数
		result, err := db.Exec("drop table if exists users")
		if err != nil {
			log.Fatal(err)
		}
		id, err := result.LastInsertId()
		rows, err := result.RowsAffected()

		fmt.Println("last insert id : ", id, " , rows affected : ", rows)
	}

	{ // Create a new table
		query := `
            CREATE TABLE if not exists users (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	/*
		One thing to notice is: By default Go uses prepared statements for inserting dynamic data into our SQL queries, which is a way to securely pass user supplied data to our database without the risk of any damage.
		tips: 使用已经准备好语句动态的提交SQL
		like this:INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)
	*/
	{ // Insert a new user
		username := "johndoe"
		password := "secret"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	{
		// Insert a new user
		username := "Cloud"
		password := "Zack"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	{ // Query a single user
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query1 := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		//查询一行,并且只能用Scan读取，没有其他接口
		if err := db.QueryRow(query1, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, username, password, createdAt)
	}

	{ // Query all users
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}
		//一般用来执行select，返回*rows
		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			log.Fatal(err)
		}
		//tips:Next被调用且没有更多的数据，将自动调用
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user
			//Scan接口
			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v\n", users)
	}

	{
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
