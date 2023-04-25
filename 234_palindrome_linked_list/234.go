/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    var vals []int
    size := 0
    for head != nil {
        vals = append(vals, head.Val)
        head = head.Next
        size += 1
    }
    half := size / 2
    i := 0
    for i != half {
        if vals[i] != vals[size-i-1]{
            return false
        }
        i += 1
    }
    return true
}
