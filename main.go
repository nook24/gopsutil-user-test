package main

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v4/host"
)

func main() {
	fmt.Println("Hello, World!")

	users, err := host.UsersWithContext(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Println("User: ", user.User)
		fmt.Println("Terminal: ", user.Terminal)
		fmt.Println("Host: ", user.Host)
		fmt.Println("Started: ", int64(user.Started))
	}

}
