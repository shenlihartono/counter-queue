package queue

import (
	"counter-queue/log"
	"strconv"
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

	for j := 0; j < len(durations); j++ {
		jobs <- customer{id: j + 1, duration: durations[j]}
	}
	close(jobs)

	for a := 1; a <= len(durations); a++ {
		<-results
	}

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
