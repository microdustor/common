package common

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func (s String) EndsWith(suffix string) bool {
	if s == NoneStr {
		return false
	}
	if string(s)[len(s)-len(suffix):] == suffix {
		return true
	}
	return false
}

func (s String) ToBytes() []byte {
	sBytes := bytes.NewBuffer(make([]byte, 0))
	sBytes.WriteString(s.Trim())
	return sBytes.Bytes()
}

func (s String) SubString(suffix byte) string {
	sBytes := s.ToBytes()
	tmpBuffer := bytes.NewBuffer(make([]byte, 0))
	for i := range sBytes {
		if sBytes[i] != suffix {
			tmpBuffer.WriteByte(sBytes[i])
		} else {
			break
		}
	}
	return tmpBuffer.String()
}

func (s String) SuffixTrim(suffix byte) String {
	sBytes := s.ToBytes()
	length := len(sBytes)
	if length == 0 {
		return s
	}
	if sBytes[length-1] == suffix {
		return String(string(sBytes[:length-1]))
	}
	return s
}

func (s String) PrefixTrim(prefix byte) String {
	sBytes := s.ToBytes()
	length := len(sBytes)
	if length == 0 {
		return s
	}
	if sBytes[0] == prefix {
		return String(string(sBytes[1:]))
	}
	return s
}

func (s String) FullTrim(flag byte) String {
	return s.PrefixTrim(flag).SuffixTrim(flag)
}

func (s String) Filter(suffix byte) String {
	sByte := s.ToBytes()
	tmpBuffer := bytes.NewBuffer(make([]byte, 0))
	for _, v := range sByte {
		if v != suffix {
			tmpBuffer.WriteByte(v)
		}
	}
	return String(tmpBuffer.String())
}

func (s String) StartsWith(prefix string) bool {
	if s == NoneStr {
		return false
	}
	if string(s)[:len(prefix)] == prefix {
		return true
	}
	return false
}

func (s String) Split(seq string) []string {
	s = String(s.Trim())
	index := strings.Index(string(s), seq)
	var result []string
	for index != -1 {
		result = append(result, s[:index].Trim())
		s = String(s[index+len(seq):].Trim())
		index = strings.Index(string(s), seq)
	}
	if len(s) > 0 {
		result = append(result, string(s))
	}
	return result
}

func (s String) Trim() string {
	return strings.Trim(string(s), SpaceStr)
}

func (s String) MD5Finger() string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
