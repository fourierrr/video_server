/*
* @Author: Nessaj
* @Date:   2019-12-10 15:29:04
* @Last Modified by:   Nessaj
* @Last Modified time: 2019-12-28 17:51:54
*/

package main

import (
    // "io"
    "os"
    "net/http"
    "io/ioutil"
    "time"
    "log"
    "github.com/julienschmidt/httprouter"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    vid:= p.ByName("vid-id")
    vl := VIDEO_DIR+vid
    video,err:=os.Open(vl)
    if err!=nil{
        log.Printf("error when we try to open file")
        sendErrorResponse(w, http.StatusInternalServerError,"internal error")
        return
    }

    w.Header().Set("Content-Type","video/mp4")
    http.ServeContent(w,r,"",time.Now(),video)
    defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    r.Body = http.MaxBytesReader(w, r.body, MAX_UPLOAD_SIZE)
    if err:=r.ParseMultipartForm(MAX_UPLOAD_SIZE); err!=nil{
        sendErrorResponse(w, http.StatusBadRequest,"File is too Big")
        return
    }

    file,_,err:=r.FormFile("file")
    if err != nil{
        sendErrorResponse(w,http.StatusInternalServerError,"Internal Error")
        return
    }

    data,err:=ioutil.ReadAll(file)
    if err!=nil{
        log.Printf("Read file error :%v",err)
        sendErrorResponse(w,http.StatusInternalServerError,"Internal Error")
    }

    fn := p.ByName("vid-id")
    err = ioutil.WriteFile(VIDEO_DIR+fn,data,0666)
    if err!=nil{
        log.Printf("Write file error:%v",err)
        sendErrorResponse(w,http.StatusInternalServerError,"Internal Error")
        return
    }

    w.WriteHeader(http.StatusCreated)
    io.WriteString(w,"uploaded success")
}