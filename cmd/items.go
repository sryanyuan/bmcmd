package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/go-ini/ini"
)

type itemInterface interface {
	read(*ini.Section) error
}

type itemCommon struct {
	ID         int
	Name       string
	Weight     int
	Type       int
	TypeString string
	Lucky      int
	Curse      int
	Hide       int
	Accuracy   int
	AtkSpeed   int
	AtkPalsy   int
	AtkPois    int
	MoveSpeed  int
	ReqType    int
	ReqValue   int
	Sex        int
	MaxDC      int
	DC         int
	MaxAC      int
	AC         int
	MaxMAC     int
	MAC        int
	MaxSC      int
	SC         int
	MaxMC      int
	MC         int
	MaxHP      int
	HP         int
	MaxMP      int
	MP         int
	MaxEXPR    int
	EXPR       int
	Level      int
	Tex        int
	Price      int
	Grade      int
	Desc       string
}

func (i *itemCommon) read(sec *ini.Section) error {
	var err error

	if id, err := strconv.Atoi(sec.Name()); nil != err {
		log.Printf("Invalid id section %s, %v\n", sec.Name(), sec.Keys())
		return err
	} else {
		i.ID = id
	}
	if i.Name, err = getIniKeyAsString(sec, "name"); nil != err {
		return err
	}
	if i.Weight, err = getIniKeyAsInt(sec, "weight"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.Type, err = getIniKeyAsInt(sec, "type"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.Tex, err = getIniKeyAsInt(sec, "tex"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.Price, err = getIniKeyAsInt(sec, "price"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.Desc, err = getIniKeyAsString(sec, "desc"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.Lucky, err = getIniKeyAsInt(sec, "lucky"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.Curse, err = getIniKeyAsInt(sec, "curse"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.Hide, err = getIniKeyAsInt(sec, "hide"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.Accuracy, err = getIniKeyAsInt(sec, "accuracy"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.AtkSpeed, err = getIniKeyAsInt(sec, "atkSpeed"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.AtkPalsy, err = getIniKeyAsInt(sec, "atkPalsy"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.AtkPois, err = getIniKeyAsInt(sec, "atkPois"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.MoveSpeed, err = getIniKeyAsInt(sec, "moveSpeed"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.ReqType, err = getIniKeyAsInt(sec, "reqType"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.ReqValue, err = getIniKeyAsInt(sec, "reqValue"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.Sex, err = getIniKeyAsInt(sec, "sex"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.MaxDC, err = getIniKeyAsInt(sec, "maxDC"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.DC, err = getIniKeyAsInt(sec, "DC"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.MaxAC, err = getIniKeyAsInt(sec, "maxAC"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.AC, err = getIniKeyAsInt(sec, "AC"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.MaxMAC, err = getIniKeyAsInt(sec, "maxMAC"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.MAC, err = getIniKeyAsInt(sec, "MAC"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.MaxSC, err = getIniKeyAsInt(sec, "maxSC"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.SC, err = getIniKeyAsInt(sec, "SC"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.MaxMC, err = getIniKeyAsInt(sec, "maxMC"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.MC, err = getIniKeyAsInt(sec, "MC"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.MaxHP, err = getIniKeyAsInt(sec, "maxHP"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.HP, err = getIniKeyAsInt(sec, "HP"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.MaxMP, err = getIniKeyAsInt(sec, "maxMP"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.MP, err = getIniKeyAsInt(sec, "MP"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.MaxEXPR, err = getIniKeyAsInt(sec, "maxEXPR"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.EXPR, err = getIniKeyAsInt(sec, "EXPR"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.Level, err = getIniKeyAsInt(sec, "level"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	if i.Grade, err = getIniKeyAsInt(sec, "grade"); nil != err {
		if !strings.Contains(err.Error(), "not exists") {
			return err
		}
	}
	return nil
}

type itemSundries struct {
	itemCommon
}
