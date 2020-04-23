package main

import "fmt"
import "sync"

type cache struct {
	items map[string] interface{}
	mu    sync.RWMutex
}

func (objcache *cache) Set(key string, data interface{}) {
	objcache.mu.Lock()
	defer objcache.mu.Unlock()
	objcache.items[key] = data
}

func (objcache *cache) Get(key string) (interface{}, error) {
	objcache.mu.RLock()
	defer objcache.mu.RUnlock()
	item, ok := objcache.items[key]
	if !ok {
		return "", fmt.Errorf("The '%s' is not presented", key)
	}
	return item, nil
}

var (
	objcache    *cache
	once sync.Once
)

func Cache() *cache {
	once.Do(func() {
		objcache = &cache{
			items: make(map[string]interface{}),
		}
	})
	
	return objcache
}


func main(){
	objcache = Cache()
	objcache.Set("integer",1)
	objcache.Set("string","sudheesh")

	
	value,err := objcache.Get("string")
	if err == nil {
		fmt.Println("String in cache ", value)
	} else {
	fmt.Println("Sudheesh has lost his touch !!!")
	}

}

