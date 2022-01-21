package logging

import (
	"github.com/MnlPhlp/gbaLib/pkg/registers"
)

func Enable() {
	registers.MgbaLogging.MGBA_ENABLE.Set(0xc0de)
	// return registers.MgbaLogging.MGBA_ENABLE.Get() == 0x1DEA
	// bytes := []byte("Test")

	// *registers.MgbaLogging.MGBA_LOG_OUT = bytes
	registers.MgbaLogging.MGBA_FLAGS.Set(1 | 0x100)
}

// static void logOutputMgba(u8 level, const char *message)
// {
//     for (int i = 0; message[i] && i < 256; i++)
//     {
//         MGBA_LOG_OUT[i] = message[i];
//     }
//
//     // FIXME: What if invalid level? Reject? Default to TRACE?
//     MGBA_FLAGS = (level - 1) | 0x100;
// }
