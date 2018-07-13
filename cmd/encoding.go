package cmd

import "github.com/axgle/mahonia"

func gbk2Utf8(s string) string {
	var dec mahonia.Decoder
	dec = mahonia.NewDecoder("gbk")
	return dec.ConvertString(s)
}
