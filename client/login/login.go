package login

import "fmt"

func Login(uid int, pwd string) error {
	fmt.Printf("uid = %d , pwd = %s", uid, pwd)
	return nil
}
