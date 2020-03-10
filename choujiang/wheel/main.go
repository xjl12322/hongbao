package main
//大转盘抽奖
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)
var mu sync.Mutex
// 奖品中奖概率
type Prate struct {
	Rate int		// 万分之N的中奖概率
	Total int		// 总数量限制，0 表示无限数量
	CodeA int		// 中奖概率起始编码（包含）
	CodeB int		// 中奖概率终止编码（包含）
	Left *int32 		// 剩余数
}

var prizeList []string = []string{
	"一等奖，火星单程船票",
	"二等奖，凉飕飕南极之旅",
	"三等奖，iPhone一部",
	"50元优惠卷",
	"",							// 没有中奖
}

// 奖品的中奖概率设置，与上面的 prizeList 对应的设置
var left1 = int32(1)
var left2 = int32(3)
var left3 = int32(5)
var left4 = int32(3000)
var left5 = int32(5000)

var rateList []Prate = []Prate{
	Prate{1, 1, 0, 1, &left1},
	Prate{2, 2, 1, 2, &left2},
	Prate{5, 10, 3, 5, &left3},
	Prate{90,0, 0, 3000, &left4},
	Prate{100,0, 0, 9999, &left5},
}

// 奖品列表
func IndexHandle(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, fmt.Sprintf("大转盘奖品列表：<br/> %s", strings.Join(prizeList, "<br/>\n")))
	return
}
func GetDebug(c *gin.Context){
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK,fmt.Sprintf("获奖概率： %v", rateList))
	return
}

func GetPrize(c *gin.Context){
	c.Header("Content-Type", "text/html")
	// 第一步，抽奖，根据随机数匹配奖品
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	// 得到个人的抽奖编码
	code := r.Intn(10000)
	var myprize string
	var prizeRate *Prate
	// 从奖品列表中匹配，是否中奖
	for i,prize := range prizeList{
		rate := &rateList[i]
		if code>= rate.CodeA && code <=rate.CodeB{
			// 满足中奖条件
			myprize = prize
			prizeRate = rate
			break
		}

	}
	if myprize == ""{
		// 没有中奖
		myprize = "很遗憾，再来一次"
		c.String(http.StatusOK,myprize)
	}
	// 第二步，发奖，是否可以发奖
	if prizeRate.Total == 0{
		// 无限奖品
		fmt.Println("中奖： ", myprize)
		c.String(http.StatusOK,myprize)
		return
	}else if *prizeRate.Left>0{
		left := atomic.AddInt32(prizeRate.Left,-1)
		if left >=0{
			fmt.Println("中奖： ", myprize)
			c.String(http.StatusOK,myprize)
			return
		}


	}
	myprize = "很遗憾，再来一次"
	c.String(http.StatusOK,myprize)
	return




}



func main()  {

	router := gin.Default()
	//dns := "root:mysqlxjl12322@163.com@tcp(152.136.43.225:3306)/bloger?parseTime=true"
	//err := db.Init(dns)
	//if err != nil {
	//	panic(err)
	//}
	//加载静态文件
	//router.Static("/static/", "./static")
	//加载模板
	//router.LoadHTMLGlob("views/*")
	router.GET("/", IndexHandle)
	router.GET("/debug", GetDebug)
	router.GET("/prize",GetPrize)
	_ = router.Run(":8001")




}
//cloud.google.com/go v0.37.0 // indirect
