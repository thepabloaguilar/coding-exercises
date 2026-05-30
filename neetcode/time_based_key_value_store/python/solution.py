class TimeMap:

    def __init__(self) -> None:
        self._store = {}

    def set(self, key: str, value: str, timestamp: int) -> None:
        if key not in self._store:
            self._store[key] = []

        self._store[key].append((value, timestamp))

    def get(self, key: str, timestamp: int) -> str:
        if key not in self._store:
            return ''

        values = self._store[key]
        if not values:
            return ''

        start, end = 0, len(values) - 1
        res = ''
        while start <= end:
            mid = (start + end) // 2
            if values[mid][1] <= timestamp:
                res = values[mid][0]
                start = mid + 1
            else:
                end = mid - 1

        return res
