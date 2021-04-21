package main

import (
	"fmt"
	"log"

	"github.com/borud/serialnum/pkg/model"
	"github.com/jessevdk/go-flags"
)

var opt struct {
	Serial    string `short:"s" description:"serial number in xxx.xxx.xxxx.xxx.xxxxx format"`
	SerialInt uint64 `short:"i" description:"serial number in integer form"`
}

func main() {
	p := flags.NewParser(&opt, flags.Default)
	_, err := p.Parse()
	if err != nil {
		return
	}

	var n model.SerialNum

	if opt.Serial != "" {
		n, err = model.ParseSerialNum(opt.Serial)
		if err != nil {
			log.Fatalf("parse error: %v", err)
		}
	}

	if opt.SerialInt != 0 {
		n = model.FromUint64(opt.SerialInt)
	}

	if opt.SerialInt == 0 && opt.Serial == "" {
		fmt.Printf("\nplease run with -h option to see help information\n\n")
		return
	}

	fmt.Printf("uint64 = %d\n", n.Uint64())
	fmt.Printf("string = %s\n", n.String())
	fmt.Printf("bytes  = %x\n", n.ToSerialNumValues())
}
