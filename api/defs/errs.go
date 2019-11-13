/*
* @Author: Nessaj
* @Date:   2019-09-23 10:00:28
* @Last Modified by:   Nessaj
* @Last Modified time: 2019-10-23 17:04:58
*/

package defs

type Err struct{
    Error string `json:"error"`
    ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
    HttpSC int
    Error Err

}

var(
    ErrorRequestBodyParseFailed = ErrorResponse{HttpSC: 400, Error: Err{Error: "Request body is not correct.", ErrorCode:"001"}}

    ErrorNotAuthUser = ErrorResponse{HttpSC:401,Error:Err{Error:"User authentication failed.", ErrorCode:"002"}}

    ErrorDBError = ErrorResponse{HttpSC:500,Error:Err{Error:"DB ops failed",ErrorCode:"003"}}

    ErrorInternalFaults = ErrorResponse{HttpSC:500,Error:Err{Error:"Internal service error",ErrorCode:"004"}}
)