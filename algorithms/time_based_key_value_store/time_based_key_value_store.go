package time_based_key_value_store

import "sort"

type TimeMap struct {
	store map[string][]TimeValue
}

type TimeValue struct {
	value     string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap{store: make(map[string][]TimeValue)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	list, ok := this.store[key]

	if !ok {
		this.store[key] = []TimeValue{{value, timestamp}}
	}

	this.store[key] = append(list, TimeValue{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	index := sort.Search(len(this.store[key]), func(j int) bool {
		return this.store[key][j].timestamp > timestamp
	})

	if index == 0 {
		return ""
	}

	return this.store[key][index-1].value
}
