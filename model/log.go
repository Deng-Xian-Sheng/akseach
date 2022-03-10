package model

import (
	"io"
	"log"
	"os"
)

var (
	LogOsFile, LogOsFileErr = GetOsFile("./Log.log")
	PanicLog                = log.New(io.MultiWriter(LogOsFile, os.Stdin), "[Panic]", log.Llongfile|log.LstdFlags)
	ErrorLog                = log.New(io.MultiWriter(LogOsFile, os.Stdin), "[Error]", log.Llongfile|log.LstdFlags)
	InfoLog                 = log.New(io.MultiWriter(LogOsFile, os.Stdin), "[Info]", log.Lshortfile|log.LstdFlags)
)
