/*
* @Author: Nessaj
* @Date:   2019-10-23 16:13:51
* @Last Modified by:   Nessaj
* @Last Modified time: 2019-10-28 19:00:10
*/

package main

import (
    "net/http"
    "video_server/api/defs"
    "video_server/api/session"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Nname"

func ValidateUserSession(r *http.Request) bool {
    sid :=r.Header.Get(HEADER_FIELD_SESSION)
    if len(sid)==0{
        return false
    }

    uname,ok := session.IsSessionExpired(sid)
    if ok{
        return false
    }

    r.Header.Add(HEADER_FIELD_UNAME,uname)
    return true
}



func ValidateUser(w http.ResponseWriter, r *http.Request)bool {
    uname:=r.Header.Get(HEADER_FIELD_UNAME)
    if len(uname)==0{
        sendErrorResponse(w,defs.ErrorNotAuthUser)
        return false
    }

    return true

}

