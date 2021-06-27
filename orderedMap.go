package common

import "sort"

type OrderedMap struct {
	keys 		[]string
	data 		map[string]interface{}
}

func NewOrderMap() OrderedMap {
	return OrderedMap{
		keys: []string{},
		data: make(map[string]interface{}),
	}
}

func (om *OrderedMap) Set(key string, value interface{})  {
	flag := false
	for _, v := range om.keys {
		if v == key {
			flag = true
			break
		}
	}
	if !flag {
		om.keys = append(om.keys, key)
		sort.Strings(om.keys)
	}
	om.data[key] = value
}

func (om *OrderedMap) Get(key string) (interface{}, bool)  {
	flag := false
	for _, v := range om.keys {
		if v == key {
			flag = true
			break
		}
	}
	if !flag {
		return nil, false
	}
	v, found := om.data[key]
	return v, found
}

func (om *OrderedMap) Remove(key string) {
	tk := make([]string, 0)
	for _, v := range om.keys {
		if v != key {
			tk = append(tk, v)
		}
	}
	om.keys = tk
	if _, found := om.data[key]; found {
		delete(om.data, key)
	}
}

func (om *OrderedMap) Use(f func(key string, v interface{}))  {
	for _, k := range om.keys {
		f(k, om.data[k])
	}
}

func (om *OrderedMap) Find(keys ...string)  (interface{}, bool) {
	tv := om
	for i := 0; i < len(keys) ; i++ {
		if i == len(keys) - 1 {
			return tv.Get(keys[i])
		} else {
			if v, found := tv.Get(keys[i]); !found {
				return v, false
			} else {
				if ttv, ok := v.(OrderedMap); !ok {
					return ttv, false
				} else {
					tv = &ttv
				}
			}
		}
	}
	return tv, false
}