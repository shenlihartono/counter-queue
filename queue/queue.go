package queue

import (
	"counter-queue/log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var durations = []int{1, 2, 4, 2, 3, 5, 2, 3, 1, 3}

type customer struct {
	id       int
	duration int
}

func ProcessQueue(numOfCounter int) {
	log.Stdout("Started")
	jobs := make(chan customer, len(durations))
	results := make(chan bool, len(durations))

	for w := 1; w <= numOfCounter; w++ {
		go worker(w, jobs, results)
	}

	var served int
	terminate := make(chan bool, 1)

loop:
	for j := 0; j < len(durations); j++ {
		time.Sleep(time.Second) // To simulate customer walking from queue line to counter

		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-c
			terminate <- true
		}()

		select {
		case <-terminate:
			log.Stdout("closing time, waiting for counters to finish up")
			break loop
		default:
			jobs <- customer{id: j + 1, duration: durations[j]}
			served++
		}
	}
	close(jobs)

	for a := 1; a <= served; a++ {
		<-results
	}

	log.Stdout("total customer served", served)
	log.Stdout("All Done")
}

func worker(id int, jobs <-chan customer, results chan<- bool) {
	for j := range jobs {
		log.Stdout("Counter", id, "started serving customer", j.id)

		dur := strconv.Itoa(j.duration)
		fmtDur := dur + "s"
		d, _ := time.ParseDuration(fmtDur)
		time.Sleep(d)

		log.Stdout("Counter", id, "done serving customer", j.id)
		results <- true
	}
}
