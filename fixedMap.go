package common

type FixedMap struct {
	keys 		[]string
	data		map[string]interface{}
}

func New() FixedMap {
	return FixedMap{
		keys: []string{},
		data: make(map[string]interface{}),
	}
}

func (fm *FixedMap) Set(key string, value interface{})  {
	flag := false
	for _, v := range fm.keys {
		if v == key {
			flag = true
			break
		}
	}
	if !flag {
		fm.keys = append(fm.keys, key)
	}
	fm.data[key] = value
}

func (fm *FixedMap) Get(key string) (interface{}, bool)  {
	flag := false
	for _, v := range fm.keys {
		if v == key {
			flag = true
			break
		}
	}
	if !flag {
		return nil, false
	}
	v, found := fm.data[key]
	return v, found
}

func (fm *FixedMap) Remove(key string) {
	tk := make([]string, 0)
	for _, v := range fm.keys {
		if v != key {
			tk = append(tk, v)
		}
	}
	fm.keys = tk
	if _, found := fm.data[key]; found {
		delete(fm.data, key)
	}
}

func (fm *FixedMap) Use(f func(key string, v interface{}))  {
	for _, k := range fm.keys {
		f(k, fm.data[k])
	}
}

func (fm *FixedMap) Find(keys ...string) (interface{}, bool) {
	tv := fm
	for i := 0; i < len(keys) ; i++ {
		if i == len(keys) - 1 {
			return tv.Get(keys[i])
		} else {
			if v, found := tv.Get(keys[i]); !found {
				return v, false
			} else {
				if ttv, ok := v.(FixedMap); !ok {
					return ttv, false
				} else {
					tv = &ttv
				}
			}
		}
	}
	return tv, false
}