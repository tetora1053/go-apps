package main 

import(
	"fmt"
	"flag"
	"os"
	"os/exec"
)

func main() {
	flag.Parse()
	args := flag.Args()
	command := args[0]
	command_args := args[1:]
	for _, command_arg := range command_args {
		out, err := exec.Command(command, command_arg).Output()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(out))
	}
}
