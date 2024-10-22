package main

import (
	"container/list"
	"fmt"
	// "sync"
)

type LRUCache struct {
	capacity  int
	hashTable map[string]*Item
	list      *list.List
	// mutexLock sync.RWMutex
}

type Cache struct {
	*LRUCache
}

type Item struct {
	dataValue   interface{}
	listElement *list.Element
}

func New(capacity int) *Cache {
	lru := &LRUCache{
		capacity:  capacity,
		hashTable: make(map[string]*Item),
		list:      list.New(),
	}
	C := &Cache{lru}
	return C
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
	// c.mutexLock.RLock()
	// defer c.mutexLock.RUnlock()
	item, found := c.hashTable[key]
	if found {
		c.updateList(item)
		return item.dataValue, true
	} else {
		return nil, false
	}
}

func (c *LRUCache) Set(key string, x interface{}) bool {
	if c.capacity < 1 {
		c.prune()
	}

	//if no space after pruning, then cannot set
	if c.capacity < 1 {
		return false
	}
	// c.mutexLock.RLock()
	// defer c.mutexLock.RUnlock()
	item, found := c.hashTable[key]

	// c.mutexLock.Lock()
	// defer c.mutexLock.Unlock()
	if !found { //key not already present in hashTable
		item = &Item{
			dataValue: x,
		}
		item.listElement = c.list.PushFront(key) //we push an element of doubly linked list with value key to the front
		// of the list and return the element or node to store in the item struct
		//it contains an *Element with value key
		c.hashTable[key] = item
		c.capacity -= 1
	} else { //key present
		item.dataValue = x
		c.updateList(item)
	}
	return true
}

func (c *LRUCache) updateList(item *Item) {
	c.list.MoveToFront(item.listElement)
}

func (c *LRUCache) prune() {
	tail := c.list.Back() //takes tail node of the doubly linked list
	if tail == nil {
		return
	}
	key := c.list.Remove(tail) //gives value that was stored in the now removed tail node, which is key
	//but it returns key of type any
	delete(c.hashTable, key.(string)) //then we cast key back to type string
	c.capacity += 1
}

func main() {
	c := New(3)
	fmt.Println(c.capacity, c.hashTable)
	seta := c.Set("a", 123)
	fmt.Println(seta, c.list.Len())
	setb := c.Set("b", 456)
	fmt.Println(setb, c.list.Len())

	setb = c.Set("b", 123)

	fmt.Println("After 2 inserts, list is like ###\n", c.list.Front().Value, c.list.Back().Value)
	fmt.Println(c.hashTable["b"].dataValue)

	fmt.Println(setb, c.list.Len())
	setc := c.Set("c", 789)
	fmt.Println(setc, c.list.Len())

	fmt.Println("After 3 inserts, list is like ###\n", c.list.Front().Value, c.list.Back().Value)
	fmt.Println(c.hashTable)
	x, found := c.Get("b")
	fmt.Println("Found at b?", x, found)

	fmt.Println("After getting b, list is like ###\n", c.list.Front().Value, c.list.Back().Value)

	setd := c.Set("d", 901)
	fmt.Println("d successful", setd)
	fmt.Println(c.hashTable)
	fmt.Println("After inserting d ###", c.list.Front().Value, c.list.Front().Next().Value, c.list.Back().Value)

	// l := list.New()
	// l.PushFront(12)
	// l.PushFront("sbhzbd")
	// l.PushBack(9)
	// fmt.Println(l.Front())
	// fmt.Println(l.Back())
	// fmt.Println(l.Len())
	// tail := l.Back()
	// fmt.Println(tail)
	// v := l.Remove(tail) //gives removed value from tail node
	// fmt.Println("v",v)
}
