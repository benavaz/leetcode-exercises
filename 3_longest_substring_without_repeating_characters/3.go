func lengthOfLongestSubstring(s string) int {
   cuenta := 0
   var longest int
   var substring string
   if len(s) == 0 {
       return 0
   } else if len(s) == 1{
       return 1
   }
   for _, letter := range s{
       substring += string(letter)
        for i, element := range substring[:len(substring)-1]{
            if letter == element{
                substring = substring[i+1:]
                if cuenta >longest{
                    longest = cuenta
                }
                cuenta -= i
               cuenta -= 1
            }
        }
        cuenta += 1
        if cuenta > longest{
            longest = cuenta 
        }
   }
   return longest 
}
