package main

import (
	"database/sql"
	"fmt"
	"ginfirst/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

//構造体で管理を行う
type Todo struct {
	ID         int    `form:"id"`
	Human      string `form:"human"`
	Content    string `form:"content"`
	Status     int    `form:"status"`
	CreatedAt  time.Time
	CreatedAtS string
}

var todo []Todo
var idMax = 1

func Saiban() int {
	idMax = idMax + 1
	return idMax
}

func GetDataTodo(c *gin.Context) {
	var b Todo
	c.Bind(&b)

	b.ID = Saiban()
	b.Status = 0
	b.CreatedAtS = time.Now().Format("2006-01-02 15:04:05")

	todo = append(todo, b)

	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todo": todo,
	})
}

func GetDoneTodo(c *gin.Context) {
	var b Todo
	c.Bind(&b)
	if b.Status == 0 {
		b.Status = 1
	} else {
		b.Status = 0
	}

	for idx, t := range todo {
		if t.ID == b.ID {
			todo[idx] = b
		}
	}
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todo": todo,
	})

}

func main() {
	fmt.Println("これはテストです")

	//データベース接続
	db,err :=sql.Open("mysql","root:@tcp(localhost:8889)/todos?parseTime=true")
	if err !=nil{
		log.Fatalf("Cannot connect database: %v",err)
	}

	//イニシャライズ
	r := gin.Default()
	//読み込み
	r.LoadHTMLFiles("./template/index.html")

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/todo", func(c *gin.Context) {

		// todo なんかここでエラーが発生している
		todos,_ :=models.Todos().All(c,db)

		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			//ここで上で作ったパラメータを渡す
			"todo": todos,
		})
	})

	r.GET("/yaru", GetDataTodo)
	r.GET("/done", GetDoneTodo)
	//サーバーを開く
	r.Run()
}
