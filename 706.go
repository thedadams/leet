type MyHashMap struct {
    m map[int]int
}


func Constructor() MyHashMap {
    return MyHashMap{
        m: make(map[int]int, 10000),
    }
}


func (this *MyHashMap) Put(key int, value int)  {
    this.m[key] = value
}


func (this *MyHashMap) Get(key int) int {
    if val, ok := this.m[key]; ok {
        return val
    }
    return -1
}


func (this *MyHashMap) Remove(key int)  {
    delete(this.m, key)
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
