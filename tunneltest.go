package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	r := httprouter.New()
	r.GET("/", checkTunnels)

	http.ListenAndServe(":6061", r)
}

func checkTunnels(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	endpoints := map[string]string{
		"asmdev":   "20695",
		"delljump": "20066",
		"delllab":  "20125",
	}
	statuses := map[string]string{}
	for name, port := range endpoints {
		var status string
		_, err := net.Dial("tcp", fmt.Sprintf("asm.lkimmel.com:%s", port))
		if err != nil {
			status = "offline"
		} else {
			status = "online"
		}
		statuses[name] = status
	}
	j, _ := json.Marshal(statuses)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", j)
}
