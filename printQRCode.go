package main

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func printQRCode(content string) error {
	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return err
	}
	fmt.Println("\nQR Code:")
	fmt.Println(qr.ToSmallString(false))
	return nil
}
