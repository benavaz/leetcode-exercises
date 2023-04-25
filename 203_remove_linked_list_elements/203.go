/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    output := &ListNode{Val: 0, Next: head}
    if head == nil {
        return head
    }
    for head != nil && head.Val == val{
        if head.Val == val{
            head = head.Next
        }
    } 
    if head == nil{
        return head
    }
    output.Next = head
    for head.Next != nil {
        if head.Next.Val == val{
            if head.Next.Next != nil{
                head.Next = head.Next.Next
            } else if head.Next.Val == val {
                head.Next = nil
            }else { 
                head = head.Next
            }
        } else {
            head = head.Next
        }
    }
    return output.Next
}
