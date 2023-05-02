package main

import (
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

func main() {
	var ary []string
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("no home dir???")
	}
	configPath := path.Join(homeDir, ".ssh", "config")
	file, err := os.Open(configPath)
	sc := bufio.NewScanner(file)
	pattern, _ := regexp.Compile(`^Host ([^*]+)$`)

	for sc.Scan() {
		text := sc.Text()
		if pattern.MatchString(text) {
			ary = append(ary, strings.Replace(text, "Host ", "", -1))
		}
	}
	sl := promptui.SelectTemplates{
		Label:    "",
		Active:   "",
		Inactive: "",
		Selected: "⚡ Connecting... {{ . | green}}",
		Details:  "Please select your connecting server...",
		Help:     "",
		FuncMap:  nil,
	}
	prompt := promptui.Select{
		Label:     "SSH CONNECTION LIST",
		Items:     ary,
		Size:      0,
		CursorPos: 0,
		IsVimMode: true,
		HideHelp:  true,
		Templates: &sl,
	}
	for {
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("exiting... %v\n", err)
			return
		}
		//fmt.Println("ssh " + result)
		cmd := exec.Command("ssh", result)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			return
		}
	}
}
