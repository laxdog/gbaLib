package logging

import (
	"github.com/MnlPhlp/gbaLib/pkg/registers"
)

func Enable() {
	registers.MgbaLogging.MGBA_ENABLE.Set(0xc0de)
	test_str := "Enabling mGBA logging"
	str_bytes := []byte(test_str)
	for idx, char := range str_bytes {
		registers.MgbaLogging.MGBA_LOG_ADDRESSES[idx].Set(char)
	}
	registers.MgbaLogging.MGBA_FLAGS.Set(1 | 0x100)

}
func Log(message string) {
	str_bytes := []byte(message)
	for idx, char := range str_bytes {
		registers.MgbaLogging.MGBA_LOG_ADDRESSES[idx].Set(char)
	}
	registers.MgbaLogging.MGBA_FLAGS.Set(1 | 0x100)
}

func Log(a, b, c, d, e, f, g, h byte) {
	var b1 byte = a
	var b2 byte = b
	var b3 byte = c
	var b4 byte = d
	var b5 byte = e
	var b6 byte = f
	var b7 byte = g
	var b8 byte = h

	registers.MgbaLogging.MGBA_LOG_OUT1.Set(b1)
	registers.MgbaLogging.MGBA_LOG_OUT2.Set(b2)
	registers.MgbaLogging.MGBA_LOG_OUT3.Set(b3)
	registers.MgbaLogging.MGBA_LOG_OUT4.Set(b4)
	registers.MgbaLogging.MGBA_LOG_OUT5.Set(b5)
	registers.MgbaLogging.MGBA_LOG_OUT6.Set(b6)
	registers.MgbaLogging.MGBA_LOG_OUT7.Set(b7)
	registers.MgbaLogging.MGBA_LOG_OUT8.Set(b8)
	registers.MgbaLogging.MGBA_FLAGS.Set(1 | 0x100)

}
