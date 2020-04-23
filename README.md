# objcache


Synchronised Object Cache for golang.


## Usage
```

	objcache = Cache()
	objcache.Set("integer",1)
	objcache.Set("string","sudheesh")

	
	value,err := objcache.Get("string")
	if err == nil {
		fmt.Println("String in cache ", value)
	} else {
	fmt.Println("Sudheesh has lost his touch !!!")
	}
  
```
