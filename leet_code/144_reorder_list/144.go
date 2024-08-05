/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    var nodes [] *ListNode
    //node1 := new(ListNode)
    tail := new(ListNode)
    tail.Val = head.Val
    node := head.Next
    size := 1
    nodes = append(nodes, &ListNode{Val:head.Val, Next: nil})

    // create new list with values reverse
    for node != nil {
        size += 1
        nodes = append(nodes, &ListNode{Val: node.Val, Next: nil})
        node = node.Next
    }
    i := size / 2
    j := 0
    siguiente := head
    for j != i {
        tail = siguiente.Next
        siguiente.Next = nodes[size-j-1]
        siguiente.Next.Next = tail
        if j+1 >= i{
        siguiente.Next.Next.Next = nil
        if i + i == size{
            siguiente.Next.Next = nil
        }
        } 
        siguiente = siguiente.Next.Next
        j += 1
    }

}
