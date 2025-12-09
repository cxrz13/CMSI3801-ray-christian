package main

import (
    "log"
	"sync"
    "sync/atomic"
    "time"
)



type Order struct {
    ID         uint64
    Customer   string
    PreparedBy string
    Reply      chan *Order 
}

var nextOrderID atomic.Uint64

func newOrder(customer string) *Order {
    id := nextOrderID.Add(1) 
    return &Order{
        ID:       id,
        Customer: customer,
        Reply:    make(chan *Order, 1), 
    }
}

var waiter = make(chan *Order, 3)

func cook(name string) {
    log.Println(name, "starting work")

    // Runs forever: read from waiter, send to order.Reply
    for order := range waiter {
        log.Println(name, "cooking order", order.ID, "for", order.Customer)

        time.Sleep(10 * time.Second)

        order.PreparedBy = name

        // Send the cooked order back to the customer
        order.Reply <- order
    }
}

func main() {
    customers := []string{
        "Ani", "Bai", "Cat", "Dao", "Eve",
        "Fay", "Gus", "Hua", "Iza", "Jai",
    }

    go cook("Remy")
    go cook("Colette")
    go cook("Linguini")

    var wg sync.WaitGroup
    wg.Add(len(customers))

    for _, name := range customers {
        go customer(name, &wg)
    }

    wg.Wait()

    close(waiter)

    log.Println("Restaurant closing")
}
