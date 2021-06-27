package common

import (
	"bufio"
)

func Init() *Host {
	var host Host
	file := String(FileHostinfo)
	if !file.IsFileExist() {
		return &host
	}

	info, err := file.Open()
	if err != nil {
		return &host
	}
	buf := bufio.NewReader(info)
	if hostNameLine, err := buf.ReadString(WrapSymbol); err != nil {
		return &host
	} else {
		host.name = String(hostNameLine).SuffixTrim(WrapSymbol).Trim()
	}
	if ipaddr, err := buf.ReadString(WrapSymbol); err == nil {
		host.ip = String(ipaddr).SuffixTrim(WrapSymbol).Trim()
	}
	return &host
}

func (h *Host) GetHostName() string {
	return h.name
}

func (h *Host) GetIpAddr() string {
	return h.ip
}

func (h *Host) GetFinger() string {
	return String(h.name).MD5Finger()
}
