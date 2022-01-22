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
