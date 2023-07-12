package cmds

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"syscall"
)

const self = "/proc/self/exe" //在 Linux中 代表当前执行的程序，固定string
const alpine = "/root/alpine" //写死


var runCommand = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		// docker run xxxx  /bin/sh
		// 执行程序本身 -- exec xxxx
		runCmd := exec.Command(self, "exec", "/bin/sh")	// 会跳到execcmd中的逻辑

		// 设置 namespace，设置后才会调用exec命令执行具体命令
		runCmd.SysProcAttr = &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWUSER,
			UidMappings: []syscall.SysProcIDMap{
				{
					ContainerID: 0,
					HostID:      os.Getuid(),
					Size:        1,
				},
			},
			GidMappings: []syscall.SysProcIDMap{
				{
					ContainerID: 0,
					HostID:      os.Getgid(),
					Size:        1,
				},
			},
		}
		runCmd.Stdin = os.Stdin
		runCmd.Stdout = os.Stdout
		runCmd.Stderr = os.Stderr
		// 执行
		if err := runCmd.Start(); err != nil {
			log.Fatalln(err)
		}
		runCmd.Wait()

	},
}

//docker run xxx  /sh
