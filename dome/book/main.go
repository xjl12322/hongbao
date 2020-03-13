package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	dome"hongbao/dome"

	"net/http"
)

func main()  {
	err := dome.initDB()
	if err != nil{
		panic(err)
	}

	r:= gin.Default()
	//加载页面
	r.LoadHTMLGlob("E:/goland/gowork/hongbao/book/templates/*")
	r.GET("/book/list",booklistHandler)
	errs := r.Run(":8088")
	if errs != nil{
		fmt.Println(errs)
	}
}

func booklistHandler(c *gin.Context)  {

	bookList,err := dome.queryAllBook()
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":1,
			"msg":err,
		})
		return
	}
	//返回数据
	c.HTML(http.StatusOK,"book_list.html",gin.H{
		"code":0,
		"data":bookList,
	})


}






