package main

import (
	"path/filepath"
	"syscall"
	"unsafe"
)

// Проиграть зыук wav
const (
	SND_SYNC uint = 0x0000 /* play synchronously (default) */
)

var (
	mmsystem = syscall.MustLoadDLL("winmm.dll")

	sndPlaySoundA = mmsystem.MustFindProc("sndPlaySoundA")
)

// SndPlaySoundA play sound file in Windows
func PlaySoundA_Windows(sound string) {
	files, err := filepath.Glob(sound)
	if err != nil {
		panic(err)
	}
	b := append([]byte(files[0]), 0)
	sndPlaySoundA.Call(uintptr(unsafe.Pointer(&b[0])), uintptr(SND_SYNC))
}

///////////////////////////
