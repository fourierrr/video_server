/*
* @Author: Nessaj
* @Date:   2019-09-27 23:26:12
* @Last Modified by:   Nessaj
* @Last Modified time: 2019-10-14 16:54:32
*/

package utils
import(
    "github.com/satori/go.uuid"
    "fmt"
)

func NewUUID()(string,error){
    ul, err := uuid.NewV4()
    return fmt.Sprintf("%s",ul),err
    // return "0b72b6c6-66b5-4faa-b8c0-21d7ce72bfac",nil
}
// func NewUUID()(string,error) {
//     uuid:=make([]byte,16)
//     n,err:=io.ReaderFull(rand.Reader, uuid)
//     if n!=len(uuid)||err!=nil{
//         return "",err
//     }
//     uuid[8] = uuid[8]&^0xc0 | 0x80
//     uuid[6] = uuid[6]&^0xf0 | 0x40
//     return fmt.Sprintf("%x-%x-%x-%x-%x",uuid[0:4],uuid[4,6],uuid)
// func main() {
//     // var a string
//     a,_:=NewUUID()
//     print(a)
// }
