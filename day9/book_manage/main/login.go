package main

import (
	"fmt"
	"go_learn/day9/book_manage/model"
	"strings"
)

func inputUser() (username, passwd string) {

	fmt.Printf("username:")
	fmt.Scanf("%s\n", &username)
	fmt.Printf("passwd:")
	fmt.Scanf("%s\n", &passwd)

	username = strings.TrimSpace(username)
	passwd = strings.TrimSpace(passwd)
	return
}

func adminLogin() (user *model.User, err error) {

	username, passwd := inputUser()
	user, err = mgr.AdminLogin(username, passwd)
	if err != nil {
		return
	}

	return
}

func userLogin() (user *model.User, err error) {

	username, passwd := inputUser()
	user, err = mgr.UserLogin(username, passwd)
	if err != nil {
		return
	}

	return
}
