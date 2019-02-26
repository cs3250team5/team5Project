package connection

import (

	"fmt"
	"net"
	"os"
	"bufio"
	"net/http"
	
)
/*
const(	
	CONN_HOST := "mail.nyx.net"
	CONN_PORT := "110"
	CONN_TYPE := "POP3"

)
*/
func main(){
/*
	un, pw := userInterface.GetUsernameAndPassword()
	un = strings.Trim(un, " \r\n")
	pw = strings.Trim(pw, " \r\n")
	fmt.Printf("Username: %s Password: %s\n", un, pw)

	conn, err := net.Dial("tcp", "mail.yx.net")
        if err != nil {
                log.Fatalln(err)
				
	*/			
}


/*
func Signin(w http.ResponseWriter, r *http.Request){
	
	creds := credentials.NewEnv()
	err := NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result := db.QueryRow("select password from users where username=$1", creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	storedCreds := &Credentials{}
	err = result.Scan(&storedCreds.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
	

}


*/