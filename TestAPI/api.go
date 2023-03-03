package main

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"log"
)

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
		fmt.Println(c)

        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
	r.POST("/somePost", posting)
    r.Run()
}

func posting(c *gin.Context) {
    var form map[string]interface{}
	c.BindJSON(&form)
	log.Print(form["Id"])
	log.Print(form["Name"])

    c.JSON(200, gin.H{
        "status": "ok",
    })
}


/* 
参考
- [GoのJSON操作【プログラミング初心者向け教材】](https://tokitsubaki.com/go-json-manipulation/411/)
- [golang gin でのJson　受け取り](https://teratail.com/questions/322074)
- [【Golang/gin】いつも使ってるgin.Contextの中身、覗いていきませんか？](https://qiita.com/SDTakeuchi/items/7f6314d166580a06d36c)
- [Go Gin爆速入門 (REST API)](https://qiita.com/ozora/items/0597e52b3f9c1759e292)
- [GolangでAPIを作るレシピ集](https://qiita.com/2san/items/7fd4fc50e01dd5b994d2)
- [フロントエンドエンジニアがGo言語でToDoリストAPIを作ってみた](https://liginc.co.jp/584227)
*/