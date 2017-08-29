package main

import (
	"bufio"
	"github.com/jdholmes/tilt/conf"
	"io"
	"log"
	"os"
	//	"fmt"
	"strings"
	"time"
)

func main() {

	var Blocks [10]conf.Block
	var pBlock  *conf.Block
	var tstamp string

	GVar := new(conf.Config)
	GVar.Configure("globals.json")

	for i := 0; i < len(GVar.BlockNames); i++ {
		pBlock = &Blocks[i]
		pBlock.MakeBlocks(GVar.BlockNames[i])
	}

	for i := 0; i < len(GVar.BlockNames); i++ {
		pBlock = &Blocks[i]
		pBlock.BlockName = GVar.BlockNames[i]
		pBlock.GiveBlocks()
	}
	//	Save Data
	//	file = "./" + file|os.O_APPEND,
	fo, er := os.OpenFile("jack.dat", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if er != nil {
		log.Fatal(er)
	}
	w := bufio.NewWriter(fo)

	fi, err := os.Open("adjust.ini")

	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(fi)
	for {
		str, err := r.ReadString('\n')

		if err == io.EOF || strings.HasPrefix(str, "[headerLines]") {
			break
		}

	}
	t := time.Now()
	tstamp = t.Format(time.ANSIC) + "\n"
	w.WriteString(tstamp)

	for {
		str, err := r.ReadString('\n')
		if err == io.EOF || strings.HasPrefix(str, "[") {
			break
		}

		w.WriteString(str)
	}

	for j := 0; j < len(GVar.BlockNames); j++ {

		for i := 0; i < len(Blocks[j].Trials); i++ {
			s := Blocks[j].Trials[i].GetLine()

			w.WriteString(s)
		}
	}
	w.Flush()
	fo.Close()

}
