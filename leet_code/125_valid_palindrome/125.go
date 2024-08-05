 
var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func clearString(str string) string {
    str = strings.ReplaceAll(str, " ", "")
    return nonAlphanumericRegex.ReplaceAllString(str, "")
}


func isPalindrome(s string) bool {
   s = clearString(s)
   size := len(s)
   half := size/2
   s = strings.ToLower(s)
   for i := 0; i < half ; i++{
       if s[i] != s[size -1 -i]{
           return false
       }
   }
   return true
}
