/*
* @Author: Nessaj
* @Date:   2019-12-10 15:28:43
* @Last Modified by:   Nessaj
* @Last Modified time: 2019-12-10 15:35:24
*/

package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
    router :=httprouter.New()

    router.GET("/videos/:vid-id",streamHandler)

    router.POST("/upload/:vid-id",uploadHandler)
    return router
}


func main() {

    r:=RegisterHandlers()
    http.ListenAndServer(":9000",r)
}