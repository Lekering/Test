package main

type MyHashSet struct {
	date []bool
}

func NewMyHashSet() MyHashSet {
	return MyHashSet{date: make([]bool, 100001)}
}

func (this *MyHashSet) Add(key int) {
	this.date[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.date[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
	return this.date[key]
}
