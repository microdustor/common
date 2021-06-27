package common

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

func InitSimpleFile(path string, uid, gid int, split byte) (*SimpleFile, error) {
	tp := String(path)
	if !tp.IsFileExist() {
		return nil, errors.New("file not existed")
	}
	var sf = &SimpleFile{
		path:       path,
		uid:        uid,
		gid:        gid,
		permission: DefaultFilePermission,
		data:       InitOrderMap(),
		split:      split,
		lock:       sync.Mutex{},
	}
	if f, err := tp.Open(); err == nil {
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		reader := bufio.NewReader(f)
		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			translateLine(line, split, &sf.data)
		}
		return sf, err
	} else {
		return sf, err
	}
}

func translateLine(line []byte, split byte, orderMap *OrderMap) {
	if len(line) <= 0 {
		return
	}

	if line[0] == CommentSymbol {
		return
	}

	var key = bytes.NewBuffer(make([]byte, 0))
	var value = bytes.NewBuffer(make([]byte, 0))
	flag := false

	for _, v := range line {
		if v == CommentSymbol {
			break
		}
		if !flag && v == split {
			flag = true
			continue
		}

		if flag {
			value.WriteByte(v)
		} else {
			key.WriteByte(v)
		}
	}
	if flag {
		orderMap.Set(String(key.String()).Trim(), String(value.String()).Trim())
	}
}

func (sf *SimpleFile) Get(key string) string {
	if sf == nil {
		return NoneStr
	}
	v := sf.data.Load(key)
	if v == nil {
		return NoneStr
	}
	return sf.data.Load(key).(string)
}

func (sf *SimpleFile) SuffixGet(key string, suffix byte) String {
	if sf == nil {
		return NoneStr
	}
	v := sf.data.Load(key)
	if v == nil {
		return NoneStr
	}

	vStr := String(v.(string))
	return vStr.SuffixTrim(suffix)
}

func (sf *SimpleFile) Remove(key string) error {
	if sf == nil {
		return errors.New("file not opened")
	}
	sf.data.Delete(key)
	return sf.save()
}

func (sf *SimpleFile) Set(key, value string) error {
	if sf == nil {
		return errors.New("file not opened")
	}
	sf.data.Set(key, value)
	return sf.save()
}

func (sf *SimpleFile) save() error {
	sf.lock.Lock()
	if sf.permission == 0 {
		sf.permission = DefaultFilePermission
	}
	if err := ioutil.WriteFile(sf.path, sf.translateData(), sf.permission); err != nil {
		sf.lock.Unlock()
		return err
	}
	if err := os.Chown(sf.path, sf.uid, sf.gid); err != nil {
		sf.lock.Unlock()
		return err
	}
	if err := os.Chmod(sf.path, sf.permission); err != nil {
		sf.lock.Unlock()
		return err
	}
	sf.lock.Unlock()
	return nil
}

func (sf *SimpleFile) translateData() []byte {
	return sf.data.ToByte(sf.split, WrapSymbol)
}
