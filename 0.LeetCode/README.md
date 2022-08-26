[![Hello programmer Welcome](https://img.shields.io/badge/Hello,Programmer!-Welcome-orange.svg?style=flat&logo=github)](https://github.com/kr-kashyap/)
![Lines of code](https://img.shields.io/tokei/lines/github/kr-kashyap/Go-Practice?style=plastic)
[![Connect on LinkedIn](https://img.shields.io/badge/--linkedin?label=LinkedIn&logo=LinkedIn&style=social)](https://www.linkedin.com/in/kashyap-nirmal/) 
[![Connect on Gmail](https://img.shields.io/badge/--Gmail?label=Gmail&logo=Gmail&style=social)](mailto:kashyap.n@knackroot.com)
![Last Commit](https://img.shields.io/github/last-commit/kr-kashyap/Go-Practice?style=plastic)

<p align="center">
<img src="https://capsule-render.vercel.app/api?type=rect&color=gradient&height=100&section=header&text=👉%20LeetCode%20Go%20Practice%20👈&fontSize=50&fontAlignY=70" /> 
</p>

## Technology stack
- Go

## Introduction

This directory contains solutions to LeetCode problems performed in Go. I am also mentioning the Go concepts used in these problems along with the basic syntax. 

Duration ` Aug 2022 - Present `

## Creator 
[`Kashyap Nirmal`](https://github.com/kr-kashyap/)

---

250822:

---

2235. Add Two Integers <br>
https://leetcode.com/problems/add-two-integers/

    Return sum

---

1480. Running Sum of 1d Array <br>
https://leetcode.com/problems/running-sum-of-1d-array/ 

Topics: For Loop, len() of Array
    
    for i:=1;i<len(nums);i++ {}

---

1920. Build Array from Permutation <br>
https://leetcode.com/problems/build-array-from-permutation/

Topics: Make Slice
    
    //Making a slice of specific length
    var num = make([]int, len(nums))

---

383. Ransom Note <br>
https://leetcode.com/problems/ransom-note/ 

Topics: Make Map, String, 
Range
Frequency count

    // Make Map
    magmap := make(map[string]int)

    // Range
    for key, element := range rnmap {}

    // Character Frequency Count    
    for i=0; i<len(magazine); i++ {
            if _, ok := magmap[string(magazine[i])]; ok == true {
                    magmap[string(magazine[i])] += 1
            } else {
                    magmap[string(magazine[i])] = 1
            }
    }

---

1431. Kids With the Greatest Number of Candies <br>
https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

---

876. Middle of the Linked List
https://leetcode.com/problems/middle-of-the-linked-list/ 

Topics: LinkedList (Struct), Struct Pointer & Literal

     // Pointer to Struct Literal
     temp := &ListNode{}
         temp = head

---

260822:

---

234. Palindrome Linked List <br>
https://leetcode.com/problems/palindrome-linked-list/ 

Topics: LinkedList (Struct), Struct Pointer & Literal, Empty Slice, Slice Append
    
//Empty Slice
    var nums []int

    // nil for Pointer
    for ; temp != nil ; temp=temp.Next {
        // Slice Append
        nums = append(nums , temp.Val)
    }
    
---

1816. Truncate Sentence <br>
https://leetcode.com/problems/truncate-sentence/ 

Topics: Array, String, Empty Slice, If-else

    // Empty Slice
    s1 := s[ : 0]

    // Be careful for scope of If-else variables
