package util

import (
	"fmt"
	"github.com/skip2/go-qrcode"
)

func RenderString(s string) {
	q, err := qrcode.New(s, qrcode.Medium)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(q.ToSmallString(false))
}
