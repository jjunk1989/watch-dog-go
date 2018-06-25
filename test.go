package watch

import (
	"log"
	"os"
	_ "strings"
	"syscall"

	"golang.org/x/sys/windows"
)

func Test() {
	log.Println("watch dog")
	ps, err := os.FindProcess(17052)
	if err != nil {
		log.Println("os find process error:", err)
	} else {
		log.Println("find process pid: ", ps.Pid)
		state, err := ps.Wait()
		if err != nil {
			log.Println("process state error", err)
			return
		}
		log.Println("process state", state)
	}

	h, err := windows.LoadLibrary("kernel32.dll")
	if err != nil {
		log.Println("LoadLibrary", err)
	}
	defer windows.FreeLibrary(h)
	proc, err := windows.GetProcAddress(h, "GetVersion")
	if err != nil {
		log.Println("GetProcAddress", err)
	}
	r, _, _ := syscall.Syscall(uintptr(proc), 0, 0, 0, 0)
	major := byte(r)
	minor := uint8(r >> 8)
	build := uint16(r >> 16)
	log.Println("windows version ", major, ".", minor, " (Build ", build, ")\n")
}

func TestWinProc() {
	// find process
	h, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	defer windows.FreeLibrary(h)
	if err != nil {
		log.Println("CreateToolhelp32Snapshot error", err)
		return
	}

	process := &windows.ProcessEntry32{
		Size: 0x238,
	}
	err = windows.Process32First(h, process)
	if err != nil {
		log.Println("process 32 firest error", err)
	} else {
		log.Println("process:", process.ProcessID)
	}

	err = windows.Process32Next(h, process)
	if err != nil {
		log.Println("process 32 next error", err)
	} else {
		log.Println("process next:", process.ProcessID)
	}
	for ; err == nil; err = windows.Process32Next(h, process) {
		log.Println("--------------------------------")
		log.Println("ProcessID:", process.ProcessID)
		log.Println("exe:", uint16BufToString(process.ExeFile[:]))
		log.Println("Usage:", process.Usage)
		log.Println("DefaultHeapID:", process.DefaultHeapID)
		log.Println("Flags:", process.Flags)
		log.Println("ModuleID:", process.ModuleID)
		log.Println("ParentProcessID:", process.ParentProcessID)
		log.Println("PriClassBase:", process.PriClassBase)
		log.Println("Size:", process.Size)
		log.Println("Threads:", process.Threads)
		log.Println("--------------------------------")
	}
}

func TestWinNetInfo() {
	//	var size uint32 = 15000
	//	var reserved uintptr
	//	adapterAddress := &windows.IpAdapterAddresses{}
	//	err := windows.GetAdaptersAddresses(windows.AF_INET, windows.GAA_FLAG_INCLUDE_PREFIX, reserved, adapterAddress, &size)
	//	if err != nil {
	//		log.Println("get adapter address err:", err)
	//	}
}
