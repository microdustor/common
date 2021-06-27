package common

type Map interface {
	// Get smap get method, return interface{} and bool
	Get(key string) (interface{}, bool)
	// Set set method
	Set(key string, value interface{})
	// Remove remove method
	Remove(key string)
}
