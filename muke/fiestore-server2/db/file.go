package db

import (
	"database/sql"
	mydb"fiestore-server/db/mysql"
	"fmt"
)
// OnFileUploadFinished : 文件上传完成，保存mete 到mysql
func OnFileUploadFinished(filehash string,filename string,filesize int64,fileaddr string)bool  {
	stmt,err := mydb.DBConn().Prepare(
		"insert ignore into tbl_file (`file_sha1`,`file_name`,`file_size`," +
			"`file_addr`,`status`) values (?,?,?,?,1)")
	if err !=nil{
		fmt.Println("Failed to prepare statement,err:"+err.Error())
		return false
	}
	defer stmt.Close()
	ret,err :=stmt.Exec(filehash,filename,filesize,fileaddr)
	if err != nil{
		fmt.Println(err.Error())
		return false
	}

	if rf,err := ret.RowsAffected();nil == err{
		if rf<=0{
			fmt.Printf("File with hash:%s has been uploaded before", filehash)
		}
		return true
	}
	return false
}
// TableFile : 文件表结构体
type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}
// GetFileMeta : 从mysql获取文件信息
func GetFileMeta(filehash string) (*TableFile, error) {
	stmt, err := mydb.DBConn().Prepare(
		"select file_sha1,file_addr,file_name,file_size from tbl_file " +
			"where file_sha1=? and status=1 limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer stmt.Close()
	tfile := TableFile{}
	err = stmt.QueryRow(filehash).Scan(
		&tfile.FileHash, &tfile.FileAddr, &tfile.FileName, &tfile.FileSize)
	if err != nil{
		fmt.Println(err.Error())
		return nil,err
	}
	return &tfile,nil
}




















