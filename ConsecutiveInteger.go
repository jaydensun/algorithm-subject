package algorithm_subject

import "fmt"

func longestConsecutiveInteger(nums []int) int {
    fmt.Println("目标数组:", nums)
    numMap := make(map[int]bool)
    // 将数字添加到map
    for _, num := range nums {
        numMap[num] = true
    }

    max := 0;
    // 遍历数字map
    for num, _ := range numMap {
        fmt.Println("判断:", num)
        // 从num+1开始，判断是否在map中存在。如果存在，连续性加1，从map中删除些元素，再继续判断num+2。以此类推
        curMax := 1
        offset := 1;
        for {
            greater := num + offset
            if ok, _ := numMap[greater]; ok {
                fmt.Println("找到连续数字:", greater)
                curMax++
                delete(numMap, greater)
                offset++
            } else {
                fmt.Println("未找到连续数字:", greater, "从比它小的开始！")
                break
            }
        }

        // 从num-1开始，判断是否在map中存在。如果存在，连续性加1，从map中删除些元素，再继续判断num-2。以此类推
        offset = 1;
        for true {
            smaller := num - offset
            if ok, _ := numMap[smaller]; ok {
                fmt.Println("找到连续数字:", smaller)
                curMax++
                delete(numMap, smaller)
                offset++
            } else {
                fmt.Println("未找到连续数字:", smaller, "继续下一轮！")
                break
            }
        }

        if curMax > max {
            max = curMax
        }
        fmt.Println("当前最大结果:", max)
    }

    return max
}