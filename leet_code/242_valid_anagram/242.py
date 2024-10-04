class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        lista = {}
        if len(s) != len(t):
            return False
            
        for letra in s:
            if letra in lista:
                lista[letra] = lista[letra] + 1
            else:
                lista[letra] = 1

        for letra in t:
            if letra not in lista:
                return False
            elif lista[letra] <= 0:
                return False
            else:
                lista[letra] = lista[letra] - 1

        return True