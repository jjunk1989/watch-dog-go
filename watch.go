package watch

import (
	"log"
	"os"
)

func Test() {
	log.Println("watch dog")
	ps, err := os.FindProcess(17052)
	if err != nil {
		log.Println("os find process error:", err)
		return
	}
	log.Println("find process pid: ", ps.Pid)

	state, err := ps.Wait()
	if err != nil {
		log.Println("process state error", err)
		return
	}
	log.Println("process state", state)

}
