package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var tempvid string
//测试用户表crud操作
func clearTables()  {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")

}


func TestMain(m *testing.M)  {
	clearTables()
	m.Run()
	clearTables()
}
func TestUserWorkFlow(t *testing.T)  {
	t.Run("add",TestAddUser)
	t.Run("get",TestGetUser)
	t.Run("delete",TestDeleteUser)
	t.Run("reget",TestRegetUser)


}

func TestAddUser(t *testing.T)  {
	err := AddUserCredential("xinjialei","123")
	if err != nil{
		t.Errorf("Error of AddUser: %v", err)
	}
}
func TestGetUser(t *testing.T)  {
	pwd,err := GetUserCredential("xinjialei")
	if err != nil || pwd != "123"{
		t.Errorf("Error of GetUser:%v",err)
	}

}
func TestDeleteUser(t *testing.T)  {
	err := DeleteUser("xinjialei","123")
	if err != nil{
		t.Errorf("Error of DeleteUser: %v", err)
	}

}
func TestRegetUser(t *testing.T)  {
	pwd,err := GetUserCredential("xinjialei")
	if err != nil{
		t.Errorf("Error of RegetUser: %v", err)
	}
	if pwd != ""{
		t.Errorf("deleteing user test fail")
	}
}



//测试视频资源crud操作

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", TestAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil{
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}

//评论测试

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", TestAddUser)
	t.Run("AddCommnets", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"

	err := AddNewComments(vid, aid, content)

	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}

	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
}






