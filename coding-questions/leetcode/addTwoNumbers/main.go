func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0  
    sum := 0
    out := &ListNode{}

    head := out

    if l1 == nil && l2 == nil {return nil}

    for l1 != nil || l2 != nil || 0 != carry {
        sum = carry

        if l1 != nil{
            sum += l1.Val
            l1 = li.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        carry = sum / 10
        sum = sum % 10

        out.Val = sum 

        if l1 != nil || l2 != nil || carry != 0{
            out.Next = &ListNode{}
            out = aout.Next
        }
    }

    return head


}

