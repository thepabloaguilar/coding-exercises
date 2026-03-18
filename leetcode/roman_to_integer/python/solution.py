class Solution:
    def romanToInt(self, s: str) -> int:
        result = 0

        values_mapping = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000,
        }

        last_value = None
        for letter in s:
            value = values_mapping[letter]
            if last_value and value > last_value:
                result = (result - last_value) + (value - last_value)
                last_value = None
            else:
                result += value
                last_value = value

        return result
