package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	panic((&http.Server{
		Addr: os.Getenv("ADDR"),
		Handler: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			var cmd exec.Cmd
			if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
				fmt.Fprintln(rw, err)
			} else if len(cmd.Args) > 0 {
				cmd2 := exec.Command(cmd.Args[0], cmd.Args[1:]...) // resolve Path and Args
				cmd.Path = cmd2.Path
				cmd.Args = cmd2.Args
				b, err := cmd.CombinedOutput()
				fmt.Fprintf(rw, "%s%v\n", b, err)
			}
		}),
	}).ListenAndServe())
}
