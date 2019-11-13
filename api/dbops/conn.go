/*
* @Author: Nessaj
* @Date:   2019-09-25 15:39:38
* @Last Modified by:   Nessaj
* @Last Modified time: 2019-09-25 15:47:01
*/

package dbops

import(
    "database/sql"
    _"github.com/go-sql-driver/mysql"
)

var(
    dbConn *sql.DB
    err error
)

func init() {
    dbConn ,err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/video_server?charset=utf8")
    if err!=nil{
        panic(err.Error())
    }
}