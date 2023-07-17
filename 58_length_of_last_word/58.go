func lengthOfLastWord(s string) int {
    size := len(s) - 1
    var count int
    for i := 0; i <= size; i++{
        if string(s[size - i]) == " " && count > 0{
            return count
	} else if string(s[size - i]) != " "{
	    count += 1
	}
    }
    return count
}
