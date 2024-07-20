func strStr(haystack string, needle string) int {
    if len(needle) > len(haystack) {
        return -1
    }

    for i := 0; i < len(haystack); i++{
        if haystack[i] == needle[0] && len(haystack) >= i + len(needle){
            if haystack[i:(i + len(needle))] == needle{
                return i
            }
        }
    }
    return -1
}