package queue

import (
	"counter-queue/log"
	"strconv"
	"sync"
	"time"
)

var durations = []int{1, 2, 4, 2, 3, 5, 2, 3, 1, 3}

var wg sync.WaitGroup

func ProcessQueue(num int) {
	log.Stdout("Started")
	for i := 0; i < num; i++ {
		wg.Add(1)
		custNo := i + 1
		go serveCustomer(custNo, durations[i])

	}

	wg.Wait()
	log.Stdout("All Done")
}

func serveCustomer(custNo int, duration int) {
	defer wg.Done()
	log.Stdout("Started serving customer", custNo)

	dur := strconv.Itoa(duration)
	fmtDur := dur + "s"
	d, _ := time.ParseDuration(fmtDur)
	time.Sleep(d)
	
	log.Stdout("Done serving customer", custNo)
}
