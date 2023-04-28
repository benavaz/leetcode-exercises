func isValid(s string) bool {
    if len(s) == 1 || len(s) == 0{
        return false
    }
    if s[0] == ')' || s[0] == '}' || s[0] == ']'{
        return false
    }
   
    var cadena string
     for _, letra := range s{
         
         if (letra == ')' || letra == ']' || letra == '}') && len(cadena) == 0{
             return false
         }
         if letra == '(' || letra == '[' || letra == '{' {
             cadena += string(letra)
         }else if letra == ')' && cadena[len(cadena)-1] == '('{
             cadena = cadena[:len(cadena)-1]
         } else if letra == ']' && cadena[len(cadena)-1] == '['{
             cadena = cadena[:len(cadena)-1]
         } else if letra == '}' && cadena[len(cadena)-1] == '{'{
             cadena = cadena[:len(cadena)-1]
         } else{
             return false
         }
          
    }
    if len(cadena) > 0{
        return false
    }
     return true
}


