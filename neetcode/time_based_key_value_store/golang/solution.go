package golang

type pair struct {
	value     string
	timestamp int
}

type TimeMap struct {
	store map[string][]pair
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], pair{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	pairs, ok := this.store[key]
	if !ok {
		return ""
	}

	result := ""
	start, end := 0, len(pairs)-1
	for start <= end {
		mid := int((start + end) / 2)
		if pairs[mid].timestamp <= timestamp {
			result = pairs[mid].value
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return result
}
