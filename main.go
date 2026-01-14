package main

import (
	"context"
	"fmt"

	"github.com/godbus/dbus/v5"
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
		fmt.Println("Current gopsutil output:")
		fmt.Printf("User: %v, TTY: %v, Host: %v, Login Timestamp: %v\n", user.User, user.Terminal, user.Host, int64(user.Started))

	}

	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}

	obj := conn.Object("org.freedesktop.login1", "/org/freedesktop/login1")
	var sessions [][]interface{}
	err = obj.Call("org.freedesktop.login1.Manager.ListSessions", 0).Store(&sessions)
	if err != nil {
		panic(err)
	}

	for _, session := range sessions {
		sessionPath := session[4].(dbus.ObjectPath)
		sessionObj := conn.Object("org.freedesktop.login1", sessionPath)

		// Get TTY property
		var tty dbus.Variant
		err := sessionObj.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.login1.Session", "TTY").Store(&tty)
		if err != nil {
			fmt.Printf("User: %v, Error getting TTY: %v\n", session[2], err)
			continue
		}
		ttyName := tty.Value().(string)

		// Get Timestamp property
		var timestamp dbus.Variant
		err = sessionObj.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.login1.Session", "Timestamp").Store(&timestamp)
		if err != nil {
			fmt.Printf("User: %v, Error getting login timestamp: %v\n", session[2], err)
			continue
		}
		loginTime := int64(timestamp.Value().(uint64) / 1000000)

		fmt.Println("dbus output:")
		fmt.Printf("User: %v, TTY: %v, Host: %v, Login Timestamp: %v\n", session[2], ttyName, session[3], loginTime)
	}

}
