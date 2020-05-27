package main

import "strconv"

func CompilingMsgUser(user User)string {
	var message string
	message = "user:{"+strconv.Itoa(user.id)+";"+user.login+";"+user.fio+";"+strconv.Itoa(user.perms)+";"+user.phone+";"+user.email+"}"
	return message
}