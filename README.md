## Queueing Counter Simulation

Imagine there's a queueing system where there are a number of counters.
Then there're some customers that need to be served.

Each counter can serve only one customer, each with different variable of duration.

When the counter closes, the counter cannot take the next person in line.
Instead, it has to take care of current customer before it closes.

Based on above illustration, we are going to create an application following certain acceptance criteria:
1. The number of counters which will be opened can be chosen (1 to 5)
2. The queue consists of 10 people with different duration of service such as following:
[1, 2, 4, 2, 3, 5, 2, 3, 1, 3] (for simulation purpose only, all durations are in second)
3. When the application get terminated, please do the Closing Simulation for each counter.