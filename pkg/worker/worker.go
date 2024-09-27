package worker

import (
	"fmt"

	"sync"

	"github.com/markgoddard/reductionist/pkg/operations"
	"github.com/markgoddard/reductionist/pkg/request"
)

type Pool struct {
	queue chan *Job
	wg    sync.WaitGroup
}

type Job struct {
	result       chan result
	operation    operations.Operation
	data         []byte
	request_data request.Data
}

type result struct {
	data []byte
	err  error
}

func NewPool(workers uint) *Pool {
	pool := Pool{queue: make(chan *Job, workers)}
	for id := range workers {
		go pool.worker(id)
	}
	pool.wg.Add(int(workers))
	return &pool
}

func (p *Pool) Execute(job *Job) {
	p.queue <- job
}

func (p *Pool) Close() {
	close(p.queue)
}

func (p *Pool) Join() {
	p.wg.Wait()
}

func (p *Pool) worker(id uint) {
	// TODO: Handle shutdown.
	fmt.Println("Worker", id, "starting")
	for job := range p.queue {
		fmt.Println("Worker", id, "received job")
		data, err := job.operation.Execute(job.data, job.request_data)
		fmt.Println("Worker", id, "completed job")
		job.result <- result{data: data, err: err}
	}
	p.wg.Done()
	fmt.Println("Worker", id, "exiting")
}

func NewJob(operation operations.Operation, data []byte, request_data request.Data) Job {
	// Wait hangs if result is unbuffered. Why?
	return Job{result: make(chan result, 1), operation: operation, data: data, request_data: request_data}
}

func (j *Job) Wait() ([]byte, error) {
	fmt.Println("Waiting for job to complete")
	result := <-j.result
	fmt.Println("Job completed")
	return result.data, result.err
}
