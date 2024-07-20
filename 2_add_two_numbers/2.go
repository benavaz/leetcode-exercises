/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // var n1, n2, n3 int
    head := ListNode{}
    tail := &ListNode{}
    val := 0
    carriage := 0
    // fmt.Println(l1.Val)
    if l1 != nil {
        val = l1.Val + l2.Val
        if val >= 10 {
            val = val - 10
            carriage = 1
        } else {
            carriage = 0
        }
        head.Val = val  
    } 

    if l1.Next == nil && l2.Next == nil{
        val = l1.Val + l2.Val
        if val >= 10{
            head.Val = val - 10
            head.Next = &ListNode{1, nil}
            return &head
        }
        head.Val = val
        return &head
    }


    if l1.Next != nil {
        head.Next = &ListNode{}
        tail = head.Next
        l1 = l1.Next
        l2 = l2.Next
    } else if l2.Next != nil {
        head.Next = &ListNode{}
        tail = head.Next
        l1 = l1.Next
        l2 = l2.Next
    }


    if l1 != nil && l2 == nil{
        for l1 != nil{
            val = l1.Val + carriage
            if val >= 10 {
                val = val - 10
                carriage = 1
            } else {
                carriage = 0
            }

            tail.Val =  val
            tail.Next = l1.Next
            if carriage == 0{
                return &head
            }
            l1 = l1.Next
            tail.Next = &ListNode{}
            tail = tail.Next
        }
        tail.Val = tail.Val + carriage
        return &head
    }

    if l2 != nil && l1 == nil{
        for l2 != nil{
            val = l2.Val + carriage
            if val >= 10 {
                val = val - 10
                carriage = 1
            } else {
                carriage = 0
            }

            tail.Val =  val
            tail.Next = l2.Next
            if carriage == 0{
                return &head
            }
            l2 = l2.Next
            tail.Next = &ListNode{}
            tail = tail.Next
        }
        tail.Val = tail.Val + carriage
        return &head
    }
/////////

    for true {        
        val = l1.Val + l2.Val + carriage
        if val >= 10 {
            val = val - 10
            carriage = 1
        } else {
            carriage = 0
        }
        tail.Val = val


        l1 = l1.Next
        l2 = l2.Next

        if l1 == nil && l2 == nil {
            if carriage == 1{
                tail.Next = &ListNode{1, nil}
            }
            return &head
        } else {
            tail.Next = &ListNode{}
            tail = tail.Next
        }

        if l1 != nil && l2 == nil{
            for l1 != nil{
                val = l1.Val + carriage
                if val >= 10 {
                    val = val - 10
                    carriage = 1
                } else {
                    carriage = 0
                }

                tail.Val =  val
                tail.Next = l1.Next
                if carriage == 0{
                    return &head
                }
                l1 = l1.Next
                tail.Next = &ListNode{}
                tail = tail.Next
            }
            tail.Val = tail.Val + carriage
            return &head
        }

        if l2 != nil && l1 == nil{
            for l2 != nil{
                val = l2.Val + carriage
                if val >= 10 {
                    val = val - 10
                    carriage = 1
                } else {
                    carriage = 0
                }

                tail.Val =  val
                tail.Next = l2.Next
                if carriage == 0{
                    return &head
                }
                l2 = l2.Next
                tail.Next = &ListNode{}
                tail = tail.Next
            }
            tail.Val = tail.Val + carriage
            return &head
        }

    }
    //fmt.Println(tail.Val)
   
    return &head
}