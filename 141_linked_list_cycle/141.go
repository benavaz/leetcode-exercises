/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    var dirs []*ListNode
    if head == nil || head.Next == nil {
        return false
    }
    for head != nil {
        //head = head.Next
        //dir = &head
        for _, val := range dirs{
            if head == val{
                return true
            }
        }
        dirs = append(dirs, head)
        head = head.Next
    }
    return false 
}
