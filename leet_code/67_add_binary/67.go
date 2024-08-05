func addBinary(a string, b string) string {
    size := 0
    if len(a) > len(b){
        size = len(a)
    } else {
        size = len(b)
    }

    sa := len(a)
    sb := len(b)

    output := ""
    la := ""
    lb := ""
    carriage := 0

    for i := 1; i <= size; i++{
        if i <= sa{
            la = string(a[sa - i])
        } else {
            la = "0"
        }

        if i <= sb{
            lb = string(b[sb - i])
        } else {
            lb = "0"
        }
        //fmt.Println(la)
        //fmt.Println(lb)

        if la == "1" && lb == "1"{
            if carriage == 1{
                output = "1" + output
                carriage = 1
            } else {
                output = "0" + output
                carriage = 1
            }

        } else if la == "1" {
            if carriage == 1{
                output = "0" + output
                carriage = 1
            } else {
                output = "1" + output
                carriage = 0
            }

        } else if lb == "1" {
            if carriage == 1{
                output = "0" + output
                carriage = 1
            } else {
                output = "1" + output
                carriage = 0
            }
        } else {
            if carriage == 1 {
                output = "1" + output
                carriage = 0
            } else {
                output = "0" + output 
            }
        }
    }
    if carriage == 1{
        output = "1" + output
    }
    return output
}