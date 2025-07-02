package hashmap

type MyHashMap struct {
	hashmap map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{
		hashmap: make(map[int]int),
	}
}

func (this *MyHashMap) Put(key int, value int)  {
	this.hashmap[key] = value
}

func (this *MyHashMap) Get(key int) int {
	return this.hashmap[key]
}

func (this *MyHashMap) Remove(key int)  {
	delete(this.hashmap, key)
}
