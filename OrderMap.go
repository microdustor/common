package common

import (
	"bytes"
	"fmt"
	"sync"
)

func InitOrderMap() OrderMap {
	return OrderMap{
		keys: make([]interface{}, 0),
		data: &sync.Map{},
	}
}

func AssertOrderMap(v interface{}) bool {
	_, ok := v.(OrderMap)
	return ok
}

func (om *OrderMap) Set(key, value interface{}) {
	existed := false
	om.lock.Lock()
	for _, v := range om.keys {
		if v == key {
			existed = true
			break
		}
	}
	if !existed {
		om.keys = append(om.keys, key)
	}
	om.lock.Unlock()
	om.data.Store(key, value)
}

//如果索引和map均存在，则认定为存在，否则认为不存在
func (om *OrderMap) Existed(key interface{}) bool {
	existed := false
	for _, v := range om.keys {
		if v == key {
			existed = true
			break
		}
	}
	if _, found := om.data.Load(key); found {
		return found && existed
	} else {
		return false
	}
}

func (om *OrderMap) Delete(key interface{}) {
	var tempKeys []interface{}
	om.lock.Lock()
	for _, v := range om.keys {
		if v != key {
			tempKeys = append(tempKeys, v)
		}
	}
	om.keys = tempKeys
	om.lock.Unlock()
	if _, found := om.data.Load(key); found {
		om.data.Delete(key)
	}
}

func (om *OrderMap) Length() int {
	return len(om.keys)
}

func (om *OrderMap) Load(key interface{}) interface{} {
	if om.Existed(key) {
		if v, found := om.data.Load(key); found {
			return v
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (om *OrderMap) Use(f func(key, value interface{})) {
	for _, k := range om.keys {
		f(k, om.Load(k))
	}
}

func (om *OrderMap) TreeLoad(keys ...interface{}) (interface{}, error) {
	var tmpValue = om.data
	for i := 0; i < len(keys); i++ {
		if v, ok := tmpValue.Load(keys[i]); ok {
			if i == len(keys)-1 {
				return v, nil
			} else {
				if AssertOrderMap(keys[i]) {
					tmpValue = v.(OrderMap).data
				} else {
					return nil, fmt.Errorf("tree read sorted smap failure, becasure key[%s]`s value[%v] cannot as sorted smap", keys[i], v)
				}
			}
		} else {
			return nil, fmt.Errorf("tree read sorted smap at key: %s failure, not found this key", keys[i])
		}
	}
	return nil, fmt.Errorf("tree read sorted smap failure")
}

func (om *OrderMap) ToByte(split, end byte) []byte {
	var buffer = bytes.NewBuffer(make([]byte, 0))
	for _, k := range om.keys {
		buffer.WriteString(k.(string))
		buffer.WriteByte(split)
		buffer.WriteString(om.Load(k).(string))
		buffer.WriteByte(end)
	}
	return buffer.Bytes()
}
