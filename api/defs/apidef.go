/*
* @Author: Nessaj
* @Date:   2019-09-23 10:00:06
* @Last Modified by:   Nessaj
* @Last Modified time: 2019-10-23 16:40:57
*/

package defs

//requests
type UserCredential struct{
    Username string `json:"user_name"`
    Pwd string `json:"Pwd"`
}

//response
type SignedUp struct{
    Success bool `json:"success"`
    SessionId string `json:"session_id"`
}


//data model
type VideoInfo struct{
    Id string
    AuthorId int
    Name string
    DisplayCtime string
}

type Comment struct{
    Id string
    VideoId string
    Author  string
    Content string
}

type SimpleSession struct{
    Username string
    TTL int64

}