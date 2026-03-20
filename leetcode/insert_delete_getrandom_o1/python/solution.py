import random


class RandomizedSet:

    _dict: dict[int, int]
    _list: list[int]

    def __init__(self) -> None:
        self._dict = dict()
        self._list = list()

    def insert(self, val: int) -> bool:
        not_exists = val not in self._dict
        if not_exists:
            self._dict[val] = len(self._list)
            self._list.append(val)
        return not_exists

    def remove(self, val: int) -> bool:
        exists = val in self._dict
        if exists:
            toRemoveIdx = self._dict.pop(val)
            replaceValue = self._list.pop()
            if val != replaceValue:
                self._list[toRemoveIdx] = replaceValue
                self._dict[replaceValue] = toRemoveIdx
        return exists

    def getRandom(self) -> int:
        return random.choice(self._list)
