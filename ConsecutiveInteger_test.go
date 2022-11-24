package algorithm_subject

import (
    "log"
    "testing"
)

func TestLongestConsecutiveInteger(t *testing.T) {
    ci := longestConsecutiveInteger([]int{1, 2, 5, 4, 3, 10, 6, 22})
    expect := 6
    if ci != expect {
        log.Fatalf("result should be %d, but is %d", expect, ci)
    }
}

func TestLongestConsecutiveInteger21(t *testing.T) {
    ci := longestConsecutiveInteger([]int{1, 2, 3, 4, 5})
    expect := 5
    if ci != expect {
        log.Fatalf("result should be %d, but is %d", expect, ci)
    }
}

func TestLongestConsecutiveInteger22(t *testing.T) {
    ci := longestConsecutiveInteger([]int{5, 4, 3, 2, 1})
    expect := 5
    if ci != expect {
        log.Fatalf("result should be %d, but is %d", expect, ci)
    }
}

func TestLongestConsecutiveInteger3(t *testing.T) {
    ci := longestConsecutiveInteger([]int{})
    expect := 0
    if ci != expect {
        log.Fatalf("result should be %d, but is %d", expect, ci)
    }
}

func TestLongestConsecutiveInteger4(t *testing.T) {
    ci := longestConsecutiveInteger([]int{1})
    expect := 1
    if ci != expect {
        log.Fatalf("result should be %d, but is %d", expect, ci)
    }
}

func TestLongestConsecutiveInteger5(t *testing.T) {
    ci := longestConsecutiveInteger([]int{1, 2, 3, 7, 8, 9, 10})
    expect := 4
    if ci != expect {
        log.Fatalf("result should be %d, but is %d", expect, ci)
    }
}

func TestLongestConsecutiveInteger6(t *testing.T) {
    ci := longestConsecutiveInteger([]int{1, 2, 3, 4, 7, 8, 9, 10, 15, 16})
    expect := 4
    if ci != expect {
        log.Fatalf("result should be %d, but is %d", expect, ci)
    }
}
