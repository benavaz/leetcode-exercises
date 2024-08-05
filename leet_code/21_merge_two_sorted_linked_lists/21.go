/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var head *ListNode
    next := new(ListNode)
    // return when any of the lists is empty
    if list1 == nil && list2 == nil{
        return head
    } else if list1 == nil{
        return list2
    } else if list2 == nil{
        return list1
    }

    // create head
    if list1.Val <= list2.Val {
        head = list1
        list1 = list1.Next
    } else if list2.Val < list1.Val{
        head = list2
        list2 = list2.Next
    }
    head.Next = next

    // loop through the two lists to add nodes to the output list
    for list1 != nil || list2 != nil{
        fmt.Println("a")
        if list1 == nil {
            next.Val = list2.Val
            next.Next = list2.Next
            return head
        } else if list2 == nil {
            next.Val = list1.Val
            next.Next = list1.Next
            return head
        } else if list1.Val <= list2.Val{
            next.Val = list1.Val
            list1 = list1.Next
        } else if list2.Val < list1.Val {
            next.Val = list2.Val
            list2 = list2.Next
        }
        next.Next = new(ListNode)
        next = next.Next
    }
    return head
}
