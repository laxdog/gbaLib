package registers

import (
	"runtime/volatile"
	"unsafe"
)

type videoRegisters struct {
	DispCnt,
	VCount,
	BG0Cnt,
	BG1Cnt,
	BG2Cnt,
	BG3Cnt,
	DispStat *volatile.Register16
}

type keyRegister struct {
	KeyCnt,
	KeyPad *volatile.Register16
}

type soundRegister struct {
	Sound1Cnt_L,
	Sound1Cnt_H,
	Sound1Cnt_X,
	Sound2Cnt_L,
	Sound2Cnt_H,
	Sound3Cnt_L,
	Sound3Cnt_H,
	Sound3Cnt_X,
	Sound4Cnt_L,
	Sound4Cnt_H,
	SoundCnt_L,
	SoundCnt_H,
	SoundCnt_X,
	SoundBias,
	WaveRam0_L,
	WaveRam0_H,
	WaveRam1_L,
	WaveRam1_H,
	WaveRam2_L,
	WaveRam2_H,
	WaveRam3_L,
	WaveRam3_H,
	FIFO_A_L,
	FIFO_A_H,
	FIFO_B_L,
	FIFO_B_H *volatile.Register16
}

type timerRegister struct {
	Tm0Cnt,
	Tm0Data,
	Tm1Cnt,
	Tm1Data,
	Tm2Cnt,
	Tm2Data,
	Tm3Cnt,
	Tm3Data *volatile.Register16
}

type interruptRegister struct {
	IE,
	IF,
	IFBios *volatile.Register16
}

type mgbaLoggingRegister struct {
	MGBA_ENABLE        *volatile.Register16
	MGBA_FLAGS         *volatile.Register16
	MGBA_LOG_ADDRESSES []*volatile.Register8
}

var (
	Video = videoRegisters{
		DispCnt:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000000))),
		DispStat: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000004))),
		VCount:   (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000006))),
		BG0Cnt:   (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000008))),
		BG1Cnt:   (*volatile.Register16)(unsafe.Pointer(uintptr(0x400000A))),
		BG2Cnt:   (*volatile.Register16)(unsafe.Pointer(uintptr(0x400000C))),
		BG3Cnt:   (*volatile.Register16)(unsafe.Pointer(uintptr(0x400000E))),
	}
	Sound = soundRegister{
		Sound1Cnt_L: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000060))),
		Sound1Cnt_H: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000062))),
		Sound1Cnt_X: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000064))),
		Sound2Cnt_L: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000068))),
		Sound2Cnt_H: (*volatile.Register16)(unsafe.Pointer(uintptr(0x400006C))),
		Sound3Cnt_L: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000070))),
		Sound3Cnt_H: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000072))),
		Sound3Cnt_X: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000074))),
		Sound4Cnt_L: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000078))),
		Sound4Cnt_H: (*volatile.Register16)(unsafe.Pointer(uintptr(0x400007C))),
		SoundCnt_L:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000080))),
		SoundCnt_H:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000082))),
		SoundCnt_X:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000084))),
		SoundBias:   (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000088))),
		WaveRam0_L:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000090))),
		WaveRam0_H:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000092))),
		WaveRam1_L:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000094))),
		WaveRam1_H:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000096))),
		WaveRam2_L:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000098))),
		WaveRam2_H:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x400009A))),
		WaveRam3_L:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x400009C))),
		WaveRam3_H:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x400009E))),
		FIFO_A_L:    (*volatile.Register16)(unsafe.Pointer(uintptr(0x40000A0))),
		FIFO_A_H:    (*volatile.Register16)(unsafe.Pointer(uintptr(0x40000A2))),
		FIFO_B_L:    (*volatile.Register16)(unsafe.Pointer(uintptr(0x40000A4))),
		FIFO_B_H:    (*volatile.Register16)(unsafe.Pointer(uintptr(0x40000A6))),
	}
	Key = keyRegister{
		KeyPad: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000130))),
		KeyCnt: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000132))),
	}
	Timer = timerRegister{
		Tm0Data: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000100))),
		Tm0Cnt:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000102))),
		Tm1Data: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000104))),
		Tm1Cnt:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000106))),
		Tm2Data: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000108))),
		Tm2Cnt:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x400010A))),
		Tm3Data: (*volatile.Register16)(unsafe.Pointer(uintptr(0x400010C))),
		Tm3Cnt:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x400010E))),
	}
	Interrupt = interruptRegister{
		IE:     (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000200))),
		IF:     (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000202))),
		IFBios: (*volatile.Register16)(unsafe.Pointer(uintptr(0x3007FF8))),
	}
	MgbaLogging = mgbaLoggingRegister{
		MGBA_ENABLE: (*volatile.Register16)(unsafe.Pointer(uintptr(0x4FFF780))),
		MGBA_FLAGS:  (*volatile.Register16)(unsafe.Pointer(uintptr(0x4FFF700))),
		MGBA_LOG_ADDRESSES: []*volatile.Register8{
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff600))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff601))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff602))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff603))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff604))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff605))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff606))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff607))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff608))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff609))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff60a))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff60b))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff60c))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff60d))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff60e))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff60f))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff610))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff611))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff612))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff613))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff614))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff615))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff616))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff617))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff618))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff619))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff61a))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff61b))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff61c))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff61d))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff61e))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff61f))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff620))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff621))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff622))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff623))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff624))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff625))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff626))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff627))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff628))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff629))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff62a))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff62b))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff62c))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff62d))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff62e))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff62f))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff630))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff631))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff632))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff633))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff634))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff635))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff636))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff637))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff638))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff639))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff63a))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff63b))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff63c))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff63d))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff63e))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff63f))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff640))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff641))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff642))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff643))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff644))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff645))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff646))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff647))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff648))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff649))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff64a))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff64b))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff64c))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff64d))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff64e))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff64f))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff650))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff651))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff652))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff653))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff654))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff655))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff656))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff657))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff658))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff659))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff65a))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff65b))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff65c))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff65d))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff65e))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff65f))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff660))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff661))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff662))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff663))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff664))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff665))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff666))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff667))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff668))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff669))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff66a))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff66b))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff66c))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff66d))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff66e))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff66f))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff670))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff671))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff672))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff673))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff674))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff675))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff676))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff677))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff678))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff679))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff67a))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff67b))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff67c))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff67d))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff67e))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff67f))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff680))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff681))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff682))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff683))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff684))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff685))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff686))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff687))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff688))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff689))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff68a))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff68b))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff68c))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff68d))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff68e))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff68f))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff690))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff691))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff692))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff693))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff694))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff695))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff696))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff697))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff698))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff699))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff69a))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff69b))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff69c))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff69d))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff69e))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff69f))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6a0))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6a1))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6a2))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6a3))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6a4))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6a5))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6a6))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6a7))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6a8))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6a9))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6aa))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6ab))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6ac))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6ad))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6ae))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6af))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6b0))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6b1))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6b2))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6b3))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6b4))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6b5))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6b6))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6b7))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6b8))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6b9))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6ba))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6bb))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6bc))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6bd))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6be))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6bf))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6c0))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6c1))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6c2))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6c3))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6c4))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6c5))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6c6))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6c7))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6c8))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6c9))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6ca))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6cb))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6cc))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6cd))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6ce))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6cf))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6d0))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6d1))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6d2))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6d3))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6d4))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6d5))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6d6))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6d7))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6d8))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6d9))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6da))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6db))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6dc))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6dd))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6de))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6df))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6e0))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6e1))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6e2))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6e3))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6e4))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6e5))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6e6))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6e7))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6e8))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6e9))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6ea))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6eb))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6ec))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6ed))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6ee))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6ef))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6f0))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6f1))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6f2))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6f3))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6f4))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6f5))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6f6))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6f7))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6f8))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6f9))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6fa))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6fb))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6fc))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6fd))),
			(*volatile.Register8)(unsafe.Pointer(uintptr(0x4fff6fe))),
		},
	}
)
