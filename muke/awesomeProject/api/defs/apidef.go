package defs

//model 表结构
type UserCreadential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`


}
//response
type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string

}

type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}

//session 相关
type SimpleSession struct {
	Username string //login name
	TTL int64
}
