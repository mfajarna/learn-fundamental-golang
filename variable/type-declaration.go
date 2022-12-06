/**
	Type Declaration

	digunakan untuk alias pada suatu tipe data
**/

package main

import "fmt"

func main() {

	type noKtp string
	type statusMhs bool

	var noKtpFajar noKtp = "1120010231023123"
	var statusMhsFajar statusMhs = true

	fmt.Println(noKtpFajar)
	fmt.Println(statusMhsFajar)
}
