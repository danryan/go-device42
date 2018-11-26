package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/crypto/ssh/terminal"
	"github.com/danryan/go-device42"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Device42 Username: ")
	username, _ := r.ReadString('\n')

	fmt.Print("Device42 Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	tp := device42.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	client := device42.NewClient(tp.Client())
	ctx := context.Background()
	devices, _, err := client.Devices.Get(ctx, "")

	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	spew.Dump(devices)
}
