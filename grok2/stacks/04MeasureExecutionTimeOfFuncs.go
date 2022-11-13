package stacks

import (
	"strconv"
	"strings"
)

/*
Statement:
We are given an integer number n, representing the number of functions running in a single-threaded CPU,
and an execution log, which is essentially a list of strings. Each string has the format:
`{function id}:{"start" | "end"}:{timestamp}`, indicating that the function with functionId either started or
stopped execution at time identified by the timestamp value. Each function has a unique ID between 0 and n-1.
Compute the exclusive time of the functions in the program.
Note: Exclusive time is the sum of the execution times for all the calls to a specific function.

Constraints:
- 1 <= n <= 100
- 1 <= logs.length <=500
- 0 <= functionID < n
- No two start/end events happen at the same time.
- Each function has a "end" log entry for each "start"

How to solve:
1. Read ID, start/end and timestamp from log string.
2. If "start", push log to stack.
3. If "end", pop from stack and compute exec time.
4. If stack isn't empty after pop, subtract exec time of called func from the calling func.
5. Store exec time in results array.
*/

type Log struct {
	id      int
	isStart bool
	time    int
}

// exclusiveTime mimics measuring the runtime of a func on a single-threaded cpu.
//
// Time Complexity: O(m log n) where m=events.
// Space Complexity: O(m log n)
func exclusiveTime(n int, logs []string) []int {
	result := make([]int, n)
	logStack := NewStack[Log]()

	for _, entry := range logs {
		log := decode(entry)

		if log.isStart {
			logStack.Push(log)
			continue
		}

		poppedLog := logStack.Pop()

		runtime := log.time - poppedLog.time + 1
		result[poppedLog.id] += runtime

		if logStack.Size() > 0 {
			result[logStack.Peak().id] -= runtime
		}
	}

	return result
}

func decode(logEntry string) Log {
	parts := strings.Split(logEntry, ":")

	funcID, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	time, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}

	return Log{
		id:      funcID,
		isStart: parts[1] == "start",
		time:    time,
	}
}
