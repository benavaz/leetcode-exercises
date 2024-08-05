// this solution only passes 192/242 test cases it does not work for all 
// the test cases because minWindow assumes that for each different 
// letter in t, the solution will have the same appearances of that letter 
// as there are in t, for instance with this input: s = "acaabmkmk", t="abc"
// the solution would be "caab" but since "a" appears twice in the solution, it does not work
// it works for all the other test cases that don't have duplicated letters
func minWindow(s string, t string) string {
    var indices []int
    var letras string
    size := len(s)
    var regresar string
    paso := len(t)
    for i, l := range s{
        if strings.Contains(t, string(l)){
            indices = append(indices, i)
            letras = letras + string(l)
        }
    }
    fmt.Println(indices)
    for i := 0; i < len(indices) - paso +1; i ++ {
        if checkCharacters(letras[i:i+paso], t){
            fmt.Println("iii")
            fmt.Println(i)
            if indices[i+paso-1] - indices[i] < size
                size = indices[i+paso-1] - indices[i
                fmt.Println(size)
                regresar = string(s[(indices[i]):(in
            }
        }
    }
    return regresar
}

// checkCharacters is a helper function that checks if two different strings
// contain the same characters
func checkCharacters(letras, s string) bool {
    for _, l := range letras {
        for j, v := range s {
            if l == v {
                s = s[0:j] + s[j+1:]
                break
            }
        }
    }
    if s == "" {
        return true
    } else {
        return false
    }
}
