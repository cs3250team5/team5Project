package smtp

import(
	"fmt"
	"io/ioutil"
	"connection"
	
)
type Profile struct{
	var Name,Address,Sgun,Sgpw string
}


func Smail(conn *Connection, EmailTo string, EmailSubject string, EmailMsg string){
	
	to := composeEmail.EmailTo()
	sub := composeEmail.EmailSubject()
	wmsg := composeEmail.EmailMsg()

}