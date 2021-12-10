package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	pb "github.com/leslesnoa/chatbox-grpc/pb"
)

// type Post struct {
// 	Id   int
// 	Name string
// 	Text string
// }

func Connect() *sql.DB {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	log.Println("start dbConnectFunc.")
	user := "test"
	password := "test"
	host := "localhost"
	port := "3306"
	dbName := "testdb"

	conn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4"

	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Println("DB接続時エラー")
		panic(err.Error())
	}
	return db
}

func GetRows(db *sql.DB) []*pb.ChatData {
	log.Println("start GetRowsFunc.")
	cmd := "SELECT * FROM posts"
	rows, err := db.Query(cmd)
	if err != nil {
		log.Println("Query実行時エラー")
		log.Println(err.Error())
	}
	defer rows.Close()
	log.Println("query success.")
	var result []*pb.ChatData
	for rows.Next() {
		var p *pb.ChatData
		p = new(pb.ChatData)
		err := rows.Scan(&p.Id, &p.Name, &p.Text)
		if err != nil {
			log.Println("クエリ実行時エラー")
			log.Fatal(err.Error())
		}
		result = append(result, p)
	}
	return result
}

func CreateRow(db *sql.DB, r *pb.ChatData) {
	log.Println("Starting CreateRowFunc.")
	// db.Prepare()
	stmtInsert, err := db.Prepare("INSERT INTO posts(name, text) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtInsert.Close()

	post := pb.ChatData{Name: r.GetName(), Text: r.GetText()}
	result, err := stmtInsert.Exec(post.Name, post.Text)
	if err != nil {
		panic(err.Error())
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(lastInsertID)

	// return &pb.User{Id: int32(lastInsertID), Name: r.Name, Email: r.Email}, nil
}
