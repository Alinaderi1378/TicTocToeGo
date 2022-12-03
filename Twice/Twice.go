package Twice

import (
	"math/rand"
	"time"
)

//Genrate second people work
type Twice struct {
	src       rand.Source
	cache     int64
	remaining int
}

func (b *Twice) Bool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}

func New() *Twice {
	rand.Seed(time.Now().UnixNano())
	return &Twice{src: rand.NewSource(time.Now().UnixNano())}
}
