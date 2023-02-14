package test

import (
	"fmt"
	"github.com/TelephoneTan/GoSyncMap/synch"
	"testing"
)

const (
	h     = 'h'
	e     = 'e'
	l     = 'l'
	o     = 'o'
	comma = ','
	sp    = ' '
	w     = 'w'
	r     = 'r'
	d     = 'd'
	ex    = '!'
)

func TestMap(t *testing.T) {
	m := synch.Map[rune, byte]{}
	m.Store(w, h)
	m.Store(o, e)
	b2, loaded := m.LoadOrStore(r, l)
	if !loaded && b2 == l {
		b2, loaded = m.LoadOrStore(r, e)
		if loaded && b2 == l {
			m.Store(l, b2)
		}
	}
	m.Store(d, o)
	m.Store(sp, ex)
	be, loaded := m.LoadAndDelete(sp)
	if loaded {
		m.Store(rune(be), comma)
	}
	p := func(key rune, value byte) bool {
		fmt.Println(string(value), ":", string(key))
		return true
	}
	m.Range(p)
	fmt.Println("====================")
	bd, ok := m.Load(d)
	if ok {
		m.Delete(rune(bd))
	}
	m.Range(p)
}
