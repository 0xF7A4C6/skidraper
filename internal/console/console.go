package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
)

func PrintLogo() {
	fmt.Print("\033[H\033[2J")
	color.Printf("<fg=594cad>%s</>\n", "     ___ _   _    _ ___")
	color.Printf("<fg=7567cf>%s</>\n", "    / __| |_(_)__| | _ \\__ _ _ __  ___ _ _")
	color.Printf("<fg=8b7ee0>%s</>\n", "    \\__ \\ / / / _` |   / _` | '_ \\/ -_) '_|")
	color.Printf("<fg=a599f0>%s</>\n", "    |___/_\\_\\_\\__,_|_|_\\__,_| .__/\\___|_|")
	color.Printf("<fg=beb5f5>%s</>\n", "                            |_|\n")
}

func Input(prompt string) string {
	color.Printf("<fg=594cad>[</><fg=67f591>?</><fg=594cad>]</> <fg=e9e4f2>%s</><fg=594cad>:</> ", prompt)
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	return strings.ReplaceAll(strings.ReplaceAll(input, "\n", ""), "\r", "")
}

func Confirm(prompt string) {
	color.Printf("<fg=e9e4f2>%s</>", prompt)
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func Log(content string) {
	color.Printf(content)
}