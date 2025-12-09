package main

import (
    "log"
	"sync"
    "sync/atomic"
	"math/rand"
    "time"
)



type Order struct {
    ID         uint64
    Customer   string
    PreparedBy string
    Reply      chan *Order 
}

var nextOrderID atomic.Uint64

func do(seconds int, action ...any) {
    log.Println(action...)
    randomMillis := 500 * seconds + rand.Intn(500 * seconds)
    time.Sleep(time.Duration(randomMillis) * time.Millisecond)
}

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

func customer(name string, wg *sync.WaitGroup) {
    defer wg.Done()

    for mealsEaten := 0; mealsEaten < 5; {
        order := newOrder(name)
        log.Println(name, "placed order", order.ID)

        select {
        case waiter <- order:
            meal := <-order.Reply

            do(2,
                "eating cooked order",
            )
            mealsEaten++

        case <-time.After(7 * time.Second):
            do(5,
                "waiting too long, abandoning order",
            )
        }
    }

    log.Println(name, "going home")
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
