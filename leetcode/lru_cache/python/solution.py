class Node:
    def __init__(self, key: int, val: int) -> None:
        self.key = key
        self.val = val
        self.prev = self.next = None


class LRUCache:
    def __init__(self, capacity: int):
        self._capacity = capacity
        self._cache = {}

        self._head = Node(0, 0)
        self._tail = Node(0, 0)

        self._head.next = self._tail
        self._tail.prev = self._head

    def get(self, key: int) -> int:
        if key in self._cache:
            node = self._cache[key]
            self._remove(node)
            self._insert(node)
            return node.val

        return -1

    def put(self, key: int, value: int) -> None:
        if key in self._cache:
            self._remove(self._cache[key])

        node = Node(key, value)
        self._cache[key] = node
        self._insert(node)

        if len(self._cache) > self._capacity:
            node_to_remove = self._head.next
            self._remove(node_to_remove)
            del self._cache[node_to_remove.key]

    def _remove(self, node: Node) -> None:
        node.prev.next, node.next.prev = node.next, node.prev

    def _insert(self, node: Node) -> None:
        prev_node, next_node = self._tail.prev, self._tail
        prev_node.next = next_node.prev = node
        node.prev, node.next = prev_node, next_node
