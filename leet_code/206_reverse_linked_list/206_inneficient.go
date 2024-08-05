/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var vals []int
    node := head
    i := 0
    for node != nil {
        fmt.Println("as")
        vals = append(vals, node.Val)
        node = node.Next
        i += 1 
    }

    node = head

    for node != nil {
        node.Val = vals[i-1]
        i -= 1
        node = node.Next
    }
    return head
}
