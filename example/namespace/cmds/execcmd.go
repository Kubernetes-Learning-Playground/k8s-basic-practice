package cmds

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"

	"syscall"
)

const ENV = "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:myname=shenyi"

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatalln(msg, err.Error())
	}
}

var execCommand = &cobra.Command{
	Use: "exec",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("error args")
		}
		runArgs := []string{}
		if len(args) > 1 {
			runArgs = args[1:]
		}
		CheckError("chroot:", syscall.Chroot(alpine))
		CheckError("Chdir:", os.Chdir("/"))

		runCmd := exec.Command(args[0], runArgs...)
		runCmd.Stdin = os.Stdin
		runCmd.Stdout = os.Stdout
		runCmd.Stderr = os.Stderr
		runCmd.Env = []string{ENV}

		if err := runCmd.Start(); err != nil {
			log.Fatal("run:", err.Error())
		}
		runCmd.Wait()

	},
}

// docker run xxx xxx
