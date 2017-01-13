package cmd

import (
	"bytes"
	"errors"
	"log"
	"os"

	"github.com/axgle/mahonia"
	"github.com/spf13/cobra"
)

func mergeItemIniFile(itemfilename, gradefilename, outputfilename string) error {
	itemfile, err := readIniFile(itemfilename)
	if nil != err {
		return err
	}
	gradefile, err := readIniFile(gradefilename)
	if nil != err {
		return err
	}

	gradeKeyName := "grade"
	sections := itemfile.Sections()
	for _, section := range sections {
		gradeSection, err := gradefile.GetSection(section.Name())
		if nil != err {
			log.Println(err)
			continue
		}
		key, err := gradeSection.GetKey(gradeKeyName)
		if nil != err {
			log.Println(err)
			continue
		}
		section.NewKey(gradeKeyName, key.MustString("0"))
	}

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	if _, err = itemfile.WriteTo(buf); nil != err {
		return err
	}

	enc := mahonia.NewEncoder("gbk")
	transcodeData := []byte(enc.ConvertString(string(buf.Bytes())))
	if nil == transcodeData ||
		0 == len(transcodeData) {
		return errors.New("Transcode to gbk failed")
	}

	// write to file
	file, err := os.OpenFile(outputfilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if nil != err {
		return err
	}
	defer file.Close()

	_, err = file.Write(transcodeData)
	return err
}

func CreateCmdMergeINI() *cobra.Command {
	var mergeItemAttribIniFile string
	var mergeItemGradeIniFile string
	var mergeOutputFile string
	var mergeCmd = &cobra.Command{
		Use:   "mergeini",
		Short: "Merge item attrib ini with item grade ini file",
		Long:  "Merge item attrib ini with item grade ini file",
		Run: func(cmd *cobra.Command, args []string) {
			if len(mergeItemAttribIniFile) == 0 ||
				len(mergeItemGradeIniFile) == 0 ||
				len(mergeOutputFile) == 0 {
				log.Println("Invalid input")
				return
			}
			if err := mergeItemIniFile(mergeItemAttribIniFile, mergeItemGradeIniFile, mergeOutputFile); nil != err {
				log.Println(err)
				return
			}
			log.Println("Done")
		},
	}
	mergeCmd.Flags().StringVarP(&mergeItemAttribIniFile, "attribfile", "a", "", "input item attrib file")
	mergeCmd.Flags().StringVarP(&mergeItemGradeIniFile, "gradefile", "g", "", "input item grade file")
	mergeCmd.Flags().StringVarP(&mergeOutputFile, "output", "o", "", "output file")

	return mergeCmd
}
