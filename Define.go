package common

import (
	"os"
	"sync"
)

const (
	MsgNotFound   = "%s not found!"
	FileHostinfo  = "/etc/hostinfo"
	TmpHostinfo   = "/tmp/hostinfo"
	MacosStr      = "darwin"
	LinuxStr      = "linux"
	WindowsStr    = "windows"
	WrapSymbol    = '\n'
	QuotaSymbol   = '"'
	QuotaStr      = "\""
	WildMatchStr  = "*"
	CommaSymbol   = ','
	CommaStr      = ","
	DotSymbol     = '.'
	DotStr        = "."
	SlaSymbol     = '\\'
	SpaceSymbol   = ' '
	SpaceStr      = " "
	NoneStr       = ""
	SemSymbol     = ';'
	SemStr        = ";"
	OblSymbol     = '/'
	OblStr        = "/"
	EqualSymbol   = '='
	EqualStr      = "="
	CommentSymbol = '#'
)

type (
	OS     int
	String string
	Host   struct {
		name string
		ip   string
	}
	SimpleFile struct {
		path       string
		uid        int
		gid        int
		permission os.FileMode
		data       OrderMap
		split      byte
		lock       sync.Mutex
	}
	OrderMap struct {
		keys []interface{}
		lock sync.Mutex
		data *sync.Map
	}
	Handler interface {
		Filter(err error, r interface{}) error
	}
	Logger interface {
		Ef(format string, a ...interface{})
	}
	hFunc      func(err error, r interface{}) error
	loggerFunc func(format string, a ...interface{})
)

var OSIntToString = map[OS]string{
	DarwinOS:  MacosStr,
	LinuxOS:   LinuxStr,
	WindowsOS: WindowsStr,
}

const (
	Darwin = iota
	Linux
	Windows
	DarwinOS  OS = Darwin
	LinuxOS   OS = Linux
	WindowsOS OS = Windows
)

const (
	DefaultFilePermission  = 0644
	DefaultPathPermission  = 0775
	DefaultNamedPermission = 0664
	DefaultNamedUid        = 25
	DefaultNamedGid        = 25
)

var CurrentHost = Init()
