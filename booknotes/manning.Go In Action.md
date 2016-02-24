## Desc

学过一些语言之后, 在学 golang 的时候有些感悟. golang 这门语言整体来看, 算是比较简单易学的语言.
语法少而且好多约定也移除了些东西. 但这门语言也有些让人蛋疼的地方. 对于初学者也有些需要注意的地方.
我大致列一下(需要注意的地方):

* slice, array, map 的内部实现. 以便了解在函数传递的时候, 是用 value receiver or pointer receiver. 还有
就是清楚 slice 中何时底层对应的 array 会重新分配.

* 在定义 method 和 func 的时候, 需要了用何种方式传递 value receiver or pointer receiver. 这之中有些通用的规则

* 熟悉了解常规 reference type.

* 了解 interface 内部的数据结构 和 method set 概念.





## gorountine Pattern

* 作为后台任务运行

    // 这个函数返回一个 WSConn Type Pointer, 然后用 goroutine 作为后台任务, 用channel接受数据, 然后不断 write 到 
    // socket connection 中
    
    func newWSConn(conn *websocket.Conn, pendingWriteNum int, maxMsgLen uint32) *WSConn {
        wsConn := new(WSConn)
        wsConn.conn = conn
        wsConn.writeChan = make(chan []byte, pendingWriteNum)
        wsConn.maxMsgLen = maxMsgLen

        // 这种编程的 Pattern 需要注意下, close by do a writeChan <- nil
        go func() {
            for b := range wsConn.writeChan {// reading op blocks here
                if b == nil {
                    break
                }

                err := conn.WriteMessage(websocket.BinaryMessage, b)
                if err != nil {
                    break
                }
            }

            conn.Close()
            wsConn.Lock()
            wsConn.closeFlag = true
            wsConn.Unlock()
        }()

        return wsConn
    }










