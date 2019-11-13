/*
* @Author: Nessaj
* @Date:   2019-07-15 16:53:02
* @Last Modified by:   Nessaj
* @Last Modified time: 2019-10-28 19:00:42
*/

package main

import (
    "io"
    "encoding/json"
    "net/http"
    "io/ioutil"
    "github.com/julienschmidt/httprouter"
    "video_server/api/dbops"
    "video_server/api/defs"
    "video_server/api/session"

)

func CreateUser(w http.ResponseWriter, r * http.Request, p httprouter.Params) {
    // io.WriteString(w, "Create User Handler...")
    res,_:=ioutil.ReadAll(r.Body)
    ubody:=&defs.UserCredential{}

    if err:=json.Unmarshal(res,ubody);err!=nil{
        sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
        return
    }

    if err:= dbops.AddUserCredential(ubody.Username, ubody.Pwd);err !=nil{
        sendErrorResponse(w,defs.ErrorDBError)
        return
    }

    id:= session.GenerateNewSessionId(ubody.Username)
    su := &defs.SignedUp{Success: true,SessionId:id}

    if resp,err:= json.Marshal(su); err!=nil{
        sendErrorResponse(w,defs.ErrorInternalFaults)
        return
    }else{
        sendNormalResponse(w,string(resp),201)
    }
}

func Login(w http.ResponseWriter, r * http.Request, p httprouter.Params) {
    uname :=p.ByName("user_name")
    io.WriteString(w,uname)
}