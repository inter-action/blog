## Timeout Pattern

    ch := make(chan int, 10)

    select{
        case _, ok := <- ch:
            //do stuff
        case <- time.After(timeout):
            //do stuff
    }

## increment update


    count := 5
    ch := make(chan bool, 5)
    for i := 0; i < count; i++ {
        go func() {
            ch <- //do stuff assign to channel
        }()
    }
    for i := 0; i < count; i++ {
        b := <-ch
        // finish a job, handle here
    }
    close(ch)



