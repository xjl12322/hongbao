package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"hongbao/muke/awesomeProject/api/defs"
	"hongbao/muke/awesomeProject/api/utils"
	"log"
	"time"
)

//数据库用户users表的crud操作
func AddUserCredential(loginName,pwd string)error  {
    stmtIns,err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?)")
	if err != nil{
		return err
	}
    _,err = stmtIns.Exec(loginName,pwd)
	//fmt.Println("11111")
    if err != nil{
    	return err
	}
    defer stmtIns.Close()
    return nil

}
//根据用户名获取密码
func GetUserCredential(loginName string)(string,error)  {
	stmtOut,err := dbConn.Prepare("select pwd from users where login_name = ?")
	if err != nil{
		log.Printf("%s",err)
		return "",err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err!= sql.ErrNoRows{
		return "",err
	}
	defer stmtOut.Close()
	return pwd,nil

}


func DeleteUser(loginName,pwd string)(error)  {
	stmtDel,err := dbConn.Prepare("delete from users where login_name=? and pwd = ?")
	if err != nil{
		log.Printf("%s",err)
		return err
	}

	_,err = stmtDel.Exec(loginName,pwd)
	if err != nil{
		return err
	}
	defer stmtDel.Close()
	return nil


}
//数据库用户视屏资源的crud操作
func AddNewVideo(aid int,name string) (*defs.VideoInfo,error) {
	vid,err := utils.NewUUID()
	if err != nil{
		return nil,err
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns,err := dbConn.Prepare(`insert into video_info 
	(id,author_id,name,display_ctime) values(?,?,?,?)`)

	if err != nil{
		return nil,err
	}
	_,err =  stmtIns.Exec(vid,aid,name,ctime)
	res := &defs.VideoInfo{Id:vid,AuthorId:aid,Name:name,DisplayCtime:ctime}
	defer stmtIns.Close()
	return res,nil

}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id=?")

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}

	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

//用户评论表的cr操作

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(` SELECT comments.id, users.Login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil
}




