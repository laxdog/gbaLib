package logging

import (
	"github.com/MnlPhlp/gbaLib/pkg/registers"
)

func Enable() {
	registers.MgbaLogging.MGBA_ENABLE.Set(0xc0de)
	var b1 byte = 'H'
	var b2 byte = 'E'
	var b3 byte = 'L'
	var b4 byte = 'L'
	var b5 byte = 'O'

	registers.MgbaLogging.MGBA_LOG_OUT1.Set(b1)
	registers.MgbaLogging.MGBA_LOG_OUT2.Set(b2)
	registers.MgbaLogging.MGBA_LOG_OUT3.Set(b3)
	registers.MgbaLogging.MGBA_LOG_OUT4.Set(b4)
	registers.MgbaLogging.MGBA_LOG_OUT5.Set(b5)
	registers.MgbaLogging.MGBA_FLAGS.Set(1 | 0x100)
}
