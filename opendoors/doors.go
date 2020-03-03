package opendoors

import (
	"fmt"
	"net/http"
	"os/exec"
)

func init(){
	http.HandleFunc("/opendoors", func(w http.ResponseWriter, r *http.Request) {
		out, _ := exec.Command("whoami").Output()
		fmt.Println(out)
	})
}

func goodguy(){
	fmt.Println("Welcome to OpenDoors where all goods comes through")
	return
}
