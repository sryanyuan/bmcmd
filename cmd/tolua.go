package cmd

import (
	"log"

	"github.com/go-ini/ini"
	"github.com/spf13/cobra"
)

const (
	itemTypeNone = iota
	itemTypeBook
	itemTypeClose
	itemTypeNecklace
	itemTypeBracelet
	itemTypeRing
	itemTypeMedal
	itemTypeHelmet
	itemTypeDrug
	itemTypeWeapon
	itemTypeOther
	itemTypeCost
	itemTypeBundle
	itemTypeScroll
	itemTypeShoe
	itemTypeBelt
	itemTypeGem
	itemTypeCharm
)

var itemTypeMap = map[int]string{
	0:  "无",
	1:  "书",
	2:  "衣服",
	3:  "项链",
	4:  "手镯",
	5:  "戒指",
	6:  "勋章",
	7:  "头盔",
	8:  "药品",
	9:  "武器",
	10: "杂物",
	11: "消耗品",
	12: "捆装物品",
	13: "卷轴",
	14: "鞋子",
	15: "腰带",
	16: "宝石",
	17: "符咒",
}

func getIniKeyAsString(sec *ini.Section, key string) (string, error) {
	ikey, err := sec.GetKey(key)
	if nil != err {
		return "", err
	}
	return ikey.MustString(""), nil
}

func getIniKeyAsInt(sec *ini.Section, key string) (int, error) {
	ikey, err := sec.GetKey(key)
	if nil != err {
		return 0, err
	}
	return ikey.MustInt(0), nil
}

func CreateCmdTolua() *cobra.Command {
	var toluaIniFile string
	var toluaOutputFile string
	var toluaIniType string
	var toluaDropFile string
	var toluaCmd = &cobra.Command{
		Use:   "tolua",
		Short: "Generate lua file from ini config files",
		Long:  "Generate lua file from ini config files",
		Run: func(cmd *cobra.Command, args []string) {
			if len(toluaIniFile) == 0 ||
				len(toluaOutputFile) == 0 ||
				len(toluaIniType) == 0 {
				log.Println("Invalid input")
				return
			}
			if toluaIniType == "item" {
				if err := genItemLuaFile(toluaIniFile, toluaOutputFile); nil != err {
					log.Println(err)
					return
				}
			} else if toluaIniType == "mons" {
				if err := genMonsLuaFile(toluaIniFile, toluaOutputFile, toluaDropFile); nil != err {
					log.Println(err)
					return
				}
			} else {
				log.Println("invalid type", toluaIniType)
				return
			}
		},
	}
	toluaCmd.Flags().StringVarP(&toluaIniFile, "inifile", "i", "", "input item ini file")
	toluaCmd.Flags().StringVarP(&toluaOutputFile, "output", "o", "", "output lua file")
	toluaCmd.Flags().StringVarP(&toluaIniType, "type", "t", "", "type of ini (item/mons)")
	toluaCmd.Flags().StringVarP(&toluaDropFile, "drop", "d", "", "optional drop file")

	return toluaCmd
}
