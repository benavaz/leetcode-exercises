func longestPalindrome(s string) string {
    var limit int
    var left string
    var right string
    size := len(s)
    var longest_palindrome string
    var palindrome string
    var is_palindrome bool
    var b int
    
    for i, letra := range(s){
        i += 1
        if len(s) == 1{
            return s
        }
        if i < size{
            if string(letra) == string(s[i]){
                palindrome = string(letra) + string(s[i-1])
            }
        } 
        
        if i+1 < size{
            if string(letra) == string(s[i+1]){
                palindrome = string(s[i-1]) + string(s[i]) + string(s[i+1])
            }
        }

        if len(palindrome) >len(longest_palindrome){
            longest_palindrome = palindrome
        }


        if i >= 1 && i+1 <= size-1{
   
            if string(letra) == string(s[i])  {
                palindrome = string(letra) + string(s[i])
                if len(palindrome) > len(longest_palindrome){
                    longest_palindrome = palindrome
                }
                is_palindrome = true
                left = s[:i-1]
                right = s[i+1:]
   
                if len(left) < len(right){
                    limit = len(left)
                } else {
                    limit = len(right)
                }
                b = 1
                size_l := len(left)
                for is_palindrome == true && limit != b-1{
                    if left[size_l - b] == right[b-1]{
                        palindrome = string(left[size_l-b]) + palindrome + string(right[b-1])
                    } else{
                        is_palindrome = false
                    }
                    if len(palindrome) > len(longest_palindrome){
                        longest_palindrome = palindrome
                    }
                    b += 1  
                }
                b = 1

                is_palindrome = false
            }
            if string(letra) == string(s[i+1]){
                palindrome = string(letra) + string(s[i]) + string(s[i+1])
                is_palindrome = true
                left = s[:i-1]
                right = s[i+2:]

                if len(left) < len(right){
                    limit = len(left)
                } else {
                    limit = len(right)
                }
                b = 0
                size_l := len(left)
                for is_palindrome == true && b < limit{
                    if left[size_l -1 -b] == right[b+1-1]{
                        palindrome = string(left[size_l-1-b]) + palindrome + string(right[b+1-1])
                    } else{
                        is_palindrome = false
                    }
                    b += 1  
                }
                b = 1

                if len(palindrome) > len(longest_palindrome){
                    longest_palindrome = palindrome
                }
            }
        }
    }
    if longest_palindrome == ""{
        return string(s[0])
    }
    return longest_palindrome
}
