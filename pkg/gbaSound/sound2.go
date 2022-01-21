package gbaSound

import (
	"runtime/volatile"

	"github.com/MnlPhlp/gbaLib/pkg/registers"
)

type Sound2Channel struct {
	SoundChannel
}

var Sound2 = Sound2Channel{
	SoundChannel: SoundChannel{
		enableBit: soundEnable2,
		regLen:    registers.Sound.Sound2Cnt_L,
		regFreq:   registers.Sound.Sound2Cnt_H,
		BuffLen:   volatile.Register16{},
		BuffFreq:  volatile.Register16{},
	},
}
