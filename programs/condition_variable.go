package main

import (
    "math/rand"
    "sync"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    count := 0
    finished := 0

    var mu sync.Mutex
    cond := sync.NewCond(&mu)

    for i := 0; i < 10; i++ {
        go func() {
            vote := requestVote()

            mu.Lock()
            defer mu.Unlock()

            if vote {
                count++
            }

            finished++
            cond.Broadcast() // must be called while holding the lock
        }()
    }

    mu.Lock()
    for count < 5 && finished != 10 {
        cond.Wait() // called when holding a lock // waken up by broadcaset.
    }

    if count >= 5 {
        println("received 5+ votes!")
    } else {
        println("lost")
    }
    mu.Unlock()
}

func requestVote() bool {
    return rand.Int()%2 == 0
}


// mu.Lock()  
// // do something that might affect the condition
// cond.Broadcast()
// mu.Unlock()

// -----

// mu.Lock()
// for condition == false {
//     cond.Wait()
// } 
// // now condition is true, and we have the lock
// mu.Unlock()