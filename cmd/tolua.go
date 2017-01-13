package cmd

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/axgle/mahonia"
	"github.com/go-ini/ini"
	"github.com/hoisie/mustache"
	"github.com/spf13/cobra"
)

type itemContext struct {
	ID    int
	Name  string
	Price int
}

const luaItemFileHeader = `-- This file is auto generate by initolua tools , DO NOT EDIT
config_constItemAttrib = {
`

const luaItemFileItemTempl = `
-- const attrib of Item iD [{{ID}}] name [{{Name}}]
[{{ID}}] = {
	Name = "{{Name}}" ,
	Price = {{Price}}
},
`

const luaItemFileTail = `
}
`

func writeItems(buf io.Writer, sec *ini.Section) error {
	var ctx itemContext
	var err error

	if id, err := strconv.Atoi(sec.Name()); nil != err {
		return err
	} else {
		ctx.ID = id
	}
	key, err := sec.GetKey("name")
	if nil != err {
		return err
	}
	ctx.Name = key.MustString("")
	key, err = sec.GetKey("price")
	if nil != err {
		return err
	}
	ctx.Price = key.MustInt(0)

	renderData := []byte(mustache.Render(luaItemFileItemTempl, &ctx))
	_, err = buf.Write(renderData)

	return err
}

func genItemLuaFile(input, output string) error {
	iniFile, err := readIniFile(input)
	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	if _, err = buf.Write([]byte(luaItemFileHeader)); nil != err {
		return err
	}

	sections := iniFile.Sections()
	for _, section := range sections {
		if err = writeItems(buf, section); nil != err {
			log.Println("Failed to write section :", section.Name())
		} else {
			log.Println("Write section", section.Name(), "success")
		}
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
	return err
}

func CreateCmdTolua() *cobra.Command {
	var toluaIniFile string
	var toluaOutputFile string
	var toluaCmd = &cobra.Command{
		Use:   "tolua",
		Short: "Generate lua file from ini config files",
		Long:  "Generate lua file from ini config files",
		Run: func(cmd *cobra.Command, args []string) {
			if len(toluaIniFile) == 0 ||
				len(toluaOutputFile) == 0 {
				log.Println("Invalid input")
				return
			}
			if err := genItemLuaFile(toluaIniFile, toluaOutputFile); nil != err {
				log.Println(err)
				return
			}
			log.Println("Done")
		},
	}
	toluaCmd.Flags().StringVarP(&toluaIniFile, "inifile", "i", "", "input item ini file")
	toluaCmd.Flags().StringVarP(&toluaOutputFile, "output", "o", "", "output lua file")

	return toluaCmd
}
