/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    tamanio := 0
    node_tamanio := head
    // obtain length of list 
    for node_tamanio != nil{ 
        tamanio += 1
        node_tamanio = node_tamanio.Next
    }
    remover := tamanio + 1 - n

    // if we need to remove the first or second nodes
    if remover == 1 {
        return head.Next
    } else if remover == 2 {
        head.Next = head.Next.Next
        return head
    }
    node := head.Next
    flag := 2

    // iterate list and logic to remove nth node for n > 2
    for node != nil{
        flag += 1
        if remover == flag {
            if n == 1 {
                node.Next = nil
                return head
            } 
            node.Next = node.Next.Next
            return head
        }
        node = node.Next
    }
    return head
}
