/*
* @Author: Nessaj
* @Date:   2019-12-10 15:28:43
* @Last Modified by:   Nessaj
* @Last Modified time: 2019-12-19 15:52:37
*/

package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
)

type middleWareHandler struct{
    r *httprouter.Router
    l *ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler{
    m:=middleWareHandler{}
    m.r = r
    m.l =NewConnLimiter(cc)
    return m
}

func (m middleWareHandler)ServeHTTP (w http.ResponseWriter, r *http.Request) {
    if !m.l.GetConn(){
        sendErrorResponse(w, http.StatusTooManyRequests, "Too many request")
        return
    }
    m.r.ServeHTTP(w,r)
    defer m.l.ReleaseConn()
}


func RegisterHandlers() *httprouter.Router {
    router :=httprouter.New()

    router.GET("/videos/:vid-id",streamHandler)

    router.POST("/upload/:vid-id",uploadHandler)
    return router
}


func main() {
    r:=RegisterHandlers()
    mh:=NewMiddleWareHandler(r,2)
    http.ListenAndServe(":9000",mh )
}