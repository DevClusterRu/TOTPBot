package app

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func GetUserToken(user string) (key string)  {
	key = ""
	out, err := exec.Command("bash", "-c", "php /usr/local/bin/multiotp/multiotp.php -urllink "+user).Output()
	if err!=nil{
		log.Println("Wrong getting key")
		return
	}
	r:=regexp.MustCompile("secret.*?&")
	key = r.FindString(string(out))
	if strings.TrimSpace(key)==""{
		log.Println("Wrong find secret")
		return
	}
	key = key[7:len(key)-1]
	return key
}