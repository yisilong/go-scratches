package main

import (
    "fmt"
)

// 正则匹配
func isMatch(text string, pattern string) bool {
    return regexMatch(text, 0, pattern, 0)
}

func regexMatch(text string, i int, pattern string, j int) bool {
    lT, lP := len(text), len(pattern)
    if j == lP { // pattern到达尾部，text必须也到达尾部
        return i == lT
    } else if i == lT { // text到达尾部, pattern只能剩余偶数，并且偶数位需要是*
        if (lP-j)%2 != 0 {
            return false
        }
        for j != lP {
            if pattern[j+1] != '*' {
                return false
            }
            j += 2
        }
        return true
    }

    var charMatch bool
    if text[i] == pattern[j] || pattern[j] == '.' {
        charMatch = true
    }

    var ok bool
    if j+1 < lP && pattern[j+1] == '*' {
        // [j,j+1]匹配0个字符直接跳过
        ok = regexMatch(text, i, pattern, j+2)
        if !ok && charMatch {
            ok = regexMatch(text, i+1, pattern, j)
        }
    } else if charMatch {
        ok = regexMatch(text, i+1, pattern, j+1)
    }

    return ok
}

func regexMatch2(text string, i int, pattern string, j int) bool {
    lT, lP := len(text), len(pattern)
    if j == lP { // pattern到达尾部，text必须也到达尾部
        return i == lT
    }

    var charMatch bool
    if i < lT && (text[i] == pattern[j] || pattern[j] == '.') {
        charMatch = true
    }

    var ok bool
    if j+1 < lP && pattern[j+1] == '*' {
        // [j,j+1]匹配0个字符直接跳过
        ok = regexMatch(text, i, pattern, j+2)
        if !ok && charMatch {
            ok = regexMatch(text, i+1, pattern, j)
        }
    } else if charMatch {
        ok = regexMatch(text, i+1, pattern, j+1)
    }

    return ok
}

func regMatch(s string, p string) bool {
    lS, lP := len(s), len(p)
    dp := make([][]bool, lS+1, lS+1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]bool, lP+1, lP+1)
    }
    dp[lS][lP] = true
    for i := lS; i >= 0; i-- {
        for j := lP - 1; j >= 0; j-- {
            var charMatch bool
            if i < lS && (s[i] == p[j] || p[j] == '.') {
                charMatch = true
            }
            if j+1 < lP && p[j+1] == '*' {
                dp[i][j] = dp[i][j+2]
                if !dp[i][j] && charMatch {
                    dp[i][j] = dp[i+1][j]
                }
            } else if charMatch {
                dp[i][j] = dp[i+1][j+1]
            }
        }
    }
    return dp[0][0]
}

func main() {
    fmt.Println(isMatch("aa", "a"), regMatch("aa", "a"))
    fmt.Println(isMatch("aa", "a*"), regMatch("aa", "a*"))
    fmt.Println(isMatch("ab", ".*"), regMatch("ab", ".*"))
    fmt.Println(isMatch("aab", "c*a*b"), regMatch("aab", "c*a*b"))
    fmt.Println(isMatch("mississippi", "mis*is*p*."), regMatch("mississippi", "mis*is*p*."))
}
