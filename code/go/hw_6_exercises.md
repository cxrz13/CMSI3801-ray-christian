**Christian Ruiz Mendez & Ray Clemons HW6 Exercises**

1. Concurrency is more all about dealing with things in structure at once but not at the same time, while parallelism is doing multiple things at once on multiple CPUs.

2. A thread in Java is an execution unit and it has its own stack. A task is aunit of work that can be run by a thread. So the thread is the worker and the task is the job essentially.

    [Runnable task = () -> System.out.println("Hello from a task");

    Thread t = new Thread(task);
    t.start();

    ExecutorService exec = Executors.newFixedThreadPool(4);
    exec.submit(task);]

3. A thread object is still just an object after it finishes. getName() and isAlive() can be called. for Ada if you call an entry on atask that has already terminated the call raises a Tasking_Error.

4. In **Java** the JVM keeps running as long as there is atleast one non-daemon thread alive. But it terminates when main finishes and all other non-daemon threads finish. Any daemon threads are just killed at that point. In **Ada** the program waits for all tasks that (directly or indirectly) belong to the main “scope” to finish before it exits. **Go** is a little different because the program terminates when the main goroutine returns (i.e., main() finishes). That’s why you need things like sync.WaitGroup, channel blocking, etc., to keep main alive until background goroutines are done.

5. You use an unbuffered channel when you want a tight synchronization or handoff. A example of when to use it is when coordinating stages where each value must be consumed immediately. 
You use a buffered channel when you want a queue with backpressure but not a strict rendezvous. An example of when to use it for the resaurant's waiter.

unbuffered chain:

    func worker(in <-chan int){
        for x := range in {
        }
    }
    func main() {
        ch := make(chan int)
        go worker(ch)
        ch <- 42
    }

7. Trying to send on a closed channel will result in the program crashing with a panic. For reading, if there's values in the buffer you will recieve them as normal. once its empty you will get the zero value, 0, "", nil.

8. switch is for values while the select statement is for channels. The default case, if there is one, runs right away if there are no channels ready. If there's multiple ready then Go picks one at random.