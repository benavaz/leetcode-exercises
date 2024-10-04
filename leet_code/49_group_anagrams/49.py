class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        anagram_map = defaultdict(list)  # Use defaultdict to avoid key errors

        for word in strs:
            sorted_word = ''.join(sorted(word))  # Sort the word to get the key
            anagram_map[sorted_word].append(word)  # Group words by the sorted key

        return list(anagram_map.values())  # Return the grouped anagrams
        