package main

import(
import (
    "io"
    "encoding/json"
    "net/http"
    "video_server/api/defs"

)

func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
    w.WriterHeader(sc)
    io.WriterString(w, errMsg)
}