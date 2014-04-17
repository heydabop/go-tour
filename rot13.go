package main

import (
	"io"
	"os"
	"strings"
	//"fmt"
)

const lalph = "abcdefghijklmnopqrstuvwxyz"
const ualph = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error){
	n, err = r.r.Read(p)
	for i := 0; p[i] != 0 && i < len(p); i++ {
		c := p[i]
		if id := strings.Index(lalph, string(c)); id != -1 {
			c = lalph[(id + 13) % 26]
		} else if id := strings.Index(ualph, string(c)); id != -1 {
			c = ualph[(id + 13) % 26]
		}
		p[i] = c
	}
	return
}

func main(){
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
