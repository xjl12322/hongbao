package controllers

import (
	"github.com/gin-gonic/gin"
	"hongbao/wblog/models"
	"net/http"
	"strconv"
)

func IndexGet(c *gin.Context) {
	var (
		pageIndex int
		//pageSize  int
		//total     int
		page      string
		err       error
		posts     []*models.Post
		//policy    *bluemonday.Policy
	)
	page = c.DefaultQuery("page","1")
	pageIndex, _ = strconv.Atoi(page)
	if pageIndex <= 0 {
		pageIndex = 1
	}
	posts, err = models.ListPublishedPost("", 1, 20)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	//total, err = models.CountPostByTag("")
	//if err != nil {
	//	c.AbortWithStatus(http.StatusInternalServerError)
	//	return
	//}
	//policy = bluemonday.StrictPolicy()
	//for _, post := range posts {
	//	post.Tags, _ = models.ListTagByPostId(strconv.FormatUint(uint64(post.ID), 10))
	//	post.Body = policy.Sanitize(string(blackfriday.MarkdownCommon([]byte(post.Body))))
	//}
	//user, _ := c.Get(CONTEXT_USER_KEY)
	//c.HTML(http.StatusOK, "index/index.html", gin.H{
	//	"posts":           posts,
	//	"tags":            models.MustListTag(),
	//	"archives":        models.MustListPostArchives(),
	//	"links":           models.MustListLinks(),
	//	"user":            user,
	//	"pageIndex":       pageIndex,
	//	"totalPage":       int(math.Ceil(float64(total) / float64(pageSize))),
	//	"path":            c.Request.URL.Path,
	//	"maxReadPosts":    models.MustListMaxReadPost(),
	//	"maxCommentPosts": models.MustListMaxCommentPost(),
	//})
	//total = 2
	c.HTML(http.StatusOK, "index/index.html", gin.H{
			"posts":           posts,
			//"tags":            models.MustListTag(),
			//"archives":        models.MustListPostArchives(),
			//"links":           models.MustListLinks(),
			//"user":            user,
			//"pageIndex":       pageIndex,
			//"totalPage":       int(math.Ceil(float64(total) / float64(pageSize))),
			//"path":            c.Request.URL.Path,
			//"maxReadPosts":    models.MustListMaxReadPost(),
			//"maxCommentPosts": models.MustListMaxCommentPost(),
	})








}