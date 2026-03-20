package golang

import "math/rand"

type RandomizedSet struct {
	valuesMap   map[int]int
	valuesSlice []int

	// This should not be included in the Leetcode solution!
	randomSource *rand.Rand
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		valuesMap:    make(map[int]int),
		valuesSlice:  make([]int, 0),
		randomSource: rand.New(rand.NewSource(1)),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.valuesMap[val]
	if !ok {
		this.valuesMap[val] = len(this.valuesSlice)
		this.valuesSlice = append(this.valuesSlice, val)
	}
	return !ok
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.valuesMap[val]
	if ok {
		lastValue := this.valuesSlice[len(this.valuesSlice)-1]
		this.valuesMap[lastValue] = idx
		this.valuesSlice[idx] = lastValue
		this.valuesSlice = this.valuesSlice[:len(this.valuesSlice)-1]
		delete(this.valuesMap, val)
	}
	return ok
}

func (this *RandomizedSet) GetRandom() int {
	return this.valuesSlice[this.randomSource.Intn(len(this.valuesSlice))]
}
