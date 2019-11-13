/*
* @Author: Nessaj
* @Date:   2019-07-15 16:01:36
* @Last Modified by:   Nessaj
* @Last Modified time: 2019-10-28 19:01:21
*/

package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
)



/*  http接口如下
type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) error
*/



//自己定义中间件，用于补充额外功能
type middleWareHandler struct {
    r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
    m:=middleWareHandler{}
    m.r=r
    return m
}

//中间件实现http的ServeHTTP接口
func (m middleWareHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
    //check session
    ValidateUserSession(r)
    m.r.ServeHTTP(w,r)
}





func RegisterHandlers() *httprouter.Router{
    router := httprouter.New()

    router.POST("/", CreateUser)

    router.POST("/user/:user_name", Login  )

    return router
}

func main() {
    r:= RegisterHandlers()
    mh:=NewMiddleWareHandler(r)
    http.ListenAndServe(":8000", mh)
}


