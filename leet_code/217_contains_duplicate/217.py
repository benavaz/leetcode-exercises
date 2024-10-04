class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        lista = {}
        for num in nums:
            if num in lista:
                return True
            lista[num] = 1

        return False