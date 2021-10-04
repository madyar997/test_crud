package main

import (
	"github.com/labstack/echo/v4"
	"test_crud/delivery/http"
	"test_crud/repository/postgres"
	"test_crud/usecase"

	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

var conn *sql.DB

func init() {
	tmpConn, err := sql.Open("pgx", "host=localhost dbname=test_crud user=postgres password=mdamdamda port=5432")
	if err != nil {
		log.Fatal(err)
	}
	conn = tmpConn

	err = conn.Ping()
	if err != nil {
		log.Fatal("Error occures while connecting to db: ", err)
	}
	fmt.Println("Connected to db successfully")
}

func main() {
	fmt.Println("Welcome to the server ")

	e := echo.New()
	ar := postgres.NewPostgresArticleRepository(conn)
	au := usecase.NewArticleUsecase(ar)
	http.NewArticleHandler(e, au)

	fmt.Println(e.Start(":8080"))
}
