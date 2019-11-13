/*
* @Author: Nessaj
* @Date:   2019-09-23 10:14:39
* @Last Modified by:   Nessaj
* @Last Modified time: 2019-10-28 14:51:09
*/

package main

import (
    "io"
    "encoding/json"
    "net/http"
    "video_server/api/defs"
)

func sendErrorResponse(w http.ResponseWriter,errResp defs.ErrorResponse) {
    w.WriteHeader(errResp.HttpSC)

    resStr,_ := json.Marshal(&errResp.Error)
    io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, resp string,sc int) {
    w.WriteHeader(sc)
    io.WriteString(w,resp)
}