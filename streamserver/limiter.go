/*
* @Author: Nessaj
* @Date:   2019-12-10 15:31:27
* @Last Modified by:   Nessaj
* @Last Modified time: 2019-12-10 17:08:32
*/

package main

import (
    "log"

)

type ConnLimiter struct {
    concurrentConn int
    bucket chan int
}


func NewConnLimiter(cc int) *ConnLimiter{
    return &ConnLimiter: {
        concurrentConn: cc,
        bucket: make(chan int,cc)
    }
}

func (cl *ConnLimiter) GetConn() bool {
    if len(cl.bucket)>= cl.concurrentConn{
        log.Printf("Reached the rate limitation.")
        return false
    }

    cl.bucket<-1
    return true
}

func (cl *ConnLimiter)ReleaseConn() {
    c:=<- cl.bucket
    log.Printf("New connction coming: %d",c)

}