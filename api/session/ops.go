/*
* @Author: Nessaj
* @Date:   2019-10-09 14:40:02
* @Last Modified by:   Nessaj
* @Last Modified time: 2019-10-28 14:50:46
*/

package session

import (
    "time"
    "sync"
    "video_server/api/defs"
    "video_server/api/dbops"
    "video_server/api/utils"
)



var sessionMap *sync.Map

func init(){
    sessionMap = &sync.Map{}
}

func nowInMilli()int64 {
    return time.Now().UnixNano()/100000 //毫秒级别

}

func deleteExpiredSession(sid string) {
    sessionMap.Delete(sid)
    dbops.DeleteSession(sid)
}


func LoadSessionsFromDB() {
    r,err:=dbops.RetriveAllSessions()
    if err!=nil{
        return
    }
    r.Range(
        func(k,v interface{}) bool{
        ss:=v.(*defs.SimpleSession)
        sessionMap.Store(k,ss)
        return true
    })


}

func GenerateNewSessionId(un string) string {
    id,_:=utils.NewUUID()
    ct:=nowInMilli()
    ttl:=ct + 30*60*1000
    ss:=&defs.SimpleSession{Username:un, TTL:ttl}
    sessionMap.Store(id,ss)
    dbops.InserSession(id,ttl,un)
    return id
}

func IsSessionExpired(sid string)(string,bool) {
    ss,ok := sessionMap.Load(sid)
    if ok{
        ct:=nowInMilli()
        if ss.(*defs.SimpleSession).TTL<ct{
            deleteExpiredSession(sid)
            return "",true
        }
        return ss.(*defs.SimpleSession).Username, false
    }

    return "",true
}



