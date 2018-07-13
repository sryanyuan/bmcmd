package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"

	"github.com/axgle/mahonia"
	"github.com/go-ini/ini"
)

const luaItemFileHeader = `-- This file is auto generate by initolua tools , DO NOT EDIT
config_constItemAttrib = {
`

const luaItemRenderTpl = `-- const attrib of Item iD [{{.ID}}] name [{{.Name}}]
	[{{.ID}}] = {
		ID = {{.ID}},
		Name = "{{.Name}}" ,
		Type = {{.Type}},
		TypeString = "{{.TypeString}}",
		{{if ne .Lucky 0}}Lucky = {{.Lucky}},
		{{end}}{{if ne .Curse 0}}Curse = {{.Curse}},
		{{end}}{{if ne .Hide 0}}Hide = {{.Hide}},
		{{end}}{{if ne .Accuracy 0}}Accuracy = {{.Accuracy}},
		{{end}}{{if ne .AtkSpeed 0}}AtkSpeed = {{.AtkSpeed}},
		{{end}}{{if ne .AtkPalsy 0}}AtkPalsy = {{.AtkPalsy}},
		{{end}}{{if ne .AtkPois 0}}AtkPois = {{.AtkPois}},
		{{end}}{{if ne .MoveSpeed 0}}MoveSpeed = {{.MoveSpeed}},
		{{end}}{{if ne .Weight 0}}Weight = {{.Weight}},
		{{end}}{{if ne .ReqType 0}}ReqType = {{.ReqType}},
		{{end}}{{if ne .ReqValue 0}}ReqValue = {{.ReqValue}},
		{{end}}{{if ne .Sex 0}}Sex = {{.Sex}},
		{{end}}{{if ne .MaxDC 0}}MaxDC = {{.MaxDC}},
		{{end}}{{if ne .DC 0}}DC = {{.DC}},
		{{end}}{{if ne .MaxAC 0}}MaxAC = {{.MaxAC}},
		{{end}}{{if ne .AC 0}}AC = {{.AC}},
		{{end}}{{if ne .MaxMAC 0}}MaxMAC = {{.MaxMAC}},
		{{end}}{{if ne .MAC 0}}MAC = {{.MAC}},
		{{end}}{{if ne .MaxSC 0}}MaxSC = {{.MaxSC}},
		{{end}}{{if ne .SC 0}}SC = {{.SC}},
		{{end}}{{if ne .MaxMC 0}}MaxMC = {{.MaxMC}},
		{{end}}{{if ne .MC 0}}MC = {{.MC}},
		{{end}}{{if ne .MaxHP 0}}MaxHP = {{.MaxHP}},
		{{end}}{{if ne .HP 0}}HP = {{.HP}},
		{{end}}{{if ne .MaxMP 0}}MaxMP = {{.MaxMP}},
		{{end}}{{if ne .MP 0}}MP = {{.MP}},
		{{end}}{{if ne .MaxEXPR 0}}MaxEXPR = {{.MaxEXPR}},
		{{end}}{{if ne .EXPR 0}}EXPR = {{.EXPR}},
		{{end}}{{if ne .Level 0}}Level = {{.Level}},
		{{end}}{{if ne .Tex 0}}Tex = {{.Tex}},
		{{end}}{{if ne .Price 0}}Price = {{.Price}},
		{{end}}{{if ne .Grade 0}}Grade = {{.Grade}},
		{{end}}Desc = "{{.Desc}}"
	},
`

const luaItemFileTail = `
}
`

func writeItems(buf io.Writer, sec *ini.Section) error {
	var err error
	var citem itemCommon
	if err = citem.read(sec); nil != err {
		return err
	}
	itemChName, ok := itemTypeMap[citem.Type]
	if !ok {
		return fmt.Errorf("Item type %d not found", citem.Type)
	}
	citem.TypeString = itemChName

	if err = renderItemToBuffer(buf, &citem, luaItemRenderTpl); nil != err {
		return err
	}
	return nil
}

func renderItemToBuffer(buf io.Writer, si itemInterface, tpl string) error {
	t := template.New("item_tpl")
	var err error
	if t, err = t.Parse(tpl); nil != err {
		return err
	}
	if err = t.Execute(buf, si); nil != err {
		return err
	}
	return err
}

func genItemLuaFile(input, output, itype string) error {
	iniFile, err := readIniFile(input)
	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	if _, err = buf.Write([]byte(luaItemFileHeader)); nil != err {
		return err
	}

	sections := iniFile.Sections()
	var successCnt int
	var failedCnt int
	for _, section := range sections {
		if err = writeItems(buf, section); nil != err {
			log.Println("Failed to write section :", section.Name(), "error = ", err)
			failedCnt++
			continue
		}
		successCnt++
	}

	if _, err = buf.Write([]byte(luaItemFileTail)); nil != err {
		return err
	}

	// encode to gbk
	enc := mahonia.NewEncoder("gbk")
	transcodeData := []byte(enc.ConvertString(string(buf.Bytes())))
	if nil == transcodeData ||
		0 == len(transcodeData) {
		return errors.New("Transcode to gbk failed")
	}

	// write to file
	luaFile, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if nil != err {
		return err
	}
	defer luaFile.Close()

	_, err = luaFile.Write(transcodeData)
	if nil != err {
		return err
	}
	log.Printf("Task done, %d items writed, %d items failed\n", successCnt, failedCnt)
	return nil
}
