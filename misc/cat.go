/* Sun Mar  9 11:18:13 PM IST 2025 */
/* By: Victor Nammour */
package main

import (
    "fmt"
)

func main() {
    h := string("hello, ")
    w := string("world")
    s := make([]byte,len(h),2*len(h) + len(w))
    copy(s,"hello, ")
    p := []byte(w)
    t := cat(s,p)
    fmt.Printf(">%s<\n",string(t))
    sp := make([]byte,len(h),len(h))
    copy(sp,"hello, ")
    t = cat(sp,p)
    fmt.Printf(">%s<\n",string(t))
    t = ncat(sp,p,1000)
    fmt.Printf(">%s<\n",string(t))

    v := []byte("vic")
    n := []byte(" nammour")
    v = temp(v,n)
    fmt.Printf(">%s<\n",string(v))
}

func cat[T comparable](s,p []T) []T {
    var t []T
    tlen := len(s) + len(p)
    if tlen <= cap(s) {
        t = s[:tlen]
    } else {
        tcap := tlen
        if tcap < 2 * len(s) {
            tcap = 2 * len(s)
        }
        t = make([]T,tlen,tcap)
        for i := 0; i < len(s); i++ {
            t[i] = s[i]
        }
    }
    //copy(t[len(s):],p)        // You can replace the for-loop below with this line
    for i := 0; i < len(p); i++ {
        t[len(s)+i] = p[i]
    }
    return t
}

func ncat[T comparable](s,p []T, n int) []T {
    if n > len(p) || n <0 {
        n = len(p)
    }
    return cat(s,p[:n])
}
