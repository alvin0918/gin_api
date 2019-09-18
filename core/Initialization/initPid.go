package Initialization

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)

/**
记录PID
*/
func InitPid() (once_pid int, started bool) {

	var pidfile string = ".pid"

	pf, err := os.OpenFile(pidfile, os.O_RDWR, 0)
	defer pf.Close()
	if os.IsNotExist(err) {
		started = false
	} else if err != nil {
		fmt.Printf("pidfile check error:%v\n", err)
		return
	} else {
		pd, _ := ioutil.ReadAll(pf)
		old_pid, err := strconv.Atoi(string(pd))
		if err == nil {
			err := syscall.Kill(old_pid, 0)
			if err == nil {
				started = true
				once_pid = old_pid
			}
		} else {
			return
		}
	}
	if !started {
		pf, err := os.Create(pidfile)
		if err != nil {
			fmt.Println("create pid file error.")
			return
		}
		new_pid := os.Getpid()
		_, err = pf.Write([]byte(fmt.Sprintf("%d", new_pid)))
		if err != nil {
			fmt.Println("write pid failed.")
		} else {
			once_pid = new_pid
		}
	}
	return
}
