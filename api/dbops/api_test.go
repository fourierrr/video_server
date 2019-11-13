/*
* @Author: Nessaj
* @Date:   2019-09-25 17:29:50
* @Last Modified by:   Nessaj
* @Last Modified time: 2019-10-22 10:44:48
*/

package dbops

import(
    "testing"
    "strconv"
    "time"
    "fmt"
)

// init(dblogin, truncate tables)

var tempvid string

func clearTables() {
    dbConn.Exec("truncate users")
    dbConn.Exec("truncate video_info")
    dbConn.Exec("truncate comments")
    dbConn.Exec("truncate sessions")

}

func TestMain(m *testing.M) {
    clearTables()
    m.Run()
    clearTables()
}

func TestUserWorkFlow(t *testing.T) {
    t.Run("Add",testAddUser)
    t.Run("Get",testGetUser)
    t.Run("Delete",testDeleteUser)
    t.Run("Reget",testRegetUser)
}

func testAddUser(t *testing.T) {
    err:=AddUserCredential("nessaj","123")
    if err!=nil{
        t.Errorf(" error of AddUser:%v",err)
    }
}
func testGetUser(t *testing.T) {
    pwd,err:=GetUserCredential("nessaj")
    if pwd!="123" ||err!=nil{
        t.Errorf("error of GetUser")
    }
}
func testDeleteUser(t *testing.T) {
    err:=Deleteuser("nessaj","123")
    if err!=nil{
        t.Errorf(" error of DeleteUser:%v",err)
    }
}

func testRegetUser(t *testing.T) {
    pwd,err :=GetUserCredential("nessaj")
    if err!=nil{
        t.Errorf("error of regetuser: %v",err)
    }

    if pwd!=""{
        t.Errorf("Delete user failed ")
    }
}


func TestVideoInfoWorkFlow(t *testing.T) {
    clearTables()
    t.Run("PrepareUser",testAddUser)
    t.Run("AddVideo",testAddVideoInfo)
    t.Run("GetVideo",testGetVideoInfo)
    t.Run("DeleteVideo",testDeleteVideoInfo)
    t.Run("RegetVideo",testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
    vi,err:=AddNewVideo(1,"my-video22")
    if err!=nil{
        t.Errorf("Error of AddVideoInfo:%v",err)
    }

    tempvid=vi.Id
}

func testGetVideoInfo(t *testing.T) {
    _,err:=GetVideoInfo(tempvid)
    if err!=nil{
        t.Errorf("Error of GetVideoInfo:%v",err)
    }

}

func testDeleteVideoInfo(t *testing.T) {
    err:=DeleteVideoInfo(tempvid)
    if err!=nil{
        t.Errorf("Error of DeleteVideoInfo:%v",err)
    }
}

func testRegetVideoInfo(t *testing.T) {
    vi,err:=GetVideoInfo(tempvid)
    if err!=nil || vi!=nil{
        t.Errorf("Error of RegetVideoInfo:%v",err)
    }
}


func TestComments(t *testing.T) {
    clearTables()
    t.Run("PrepareUser", testAddUser)
    t.Run("AddComments", testAddComments)
    t.Run("ListComments", testListComments)

}

func testAddComments(t *testing.T) {
    vid:="12345"
    aid:=1
    content:="niubi111"

    err:= AddNewComments(vid,aid,content)

    if err!=nil{
        t.Errorf("Error of AddNewComments: %v",err)

    }
}

func testListComments(t *testing.T) {
    vid:="12345"
    from:=1514764800
    to,_ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
    res,err:= ListComments(vid ,from, to)
    print(res)
    if err!=nil{
        t.Errorf("Error of ListComments: %v",err)
    }

    for i,ele:= range res{
        fmt.Printf("comment: %d,%v \n",i,ele)
        // print("123\n")
    }



}