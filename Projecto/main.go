package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

type Worker struct {
	Id int
	//Cola de trabajo
	JobQueue    chan Job
	WorkerPool  chan chan Job
	QuitChannel chan bool
}

// Encargado de enviar a los workers
type Dispatcher struct {
	WorkerPool chan chan Job
	//Evitar el demasiados concurencia
	MaxWorker int
	JobQueue  chan Job
}

// Constructor
func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:          id,
		JobQueue:    make(chan Job),
		WorkerPool:  workerPool,
		QuitChannel: make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue
			//Multipexacion
			select {
			case job := <-w.JobQueue:
				fmt.Printf("Trabajo con id %d Inicio\n", w.Id)
				fib := Fibonnaci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("El trabajo con id %d termino con resultado fibonacci %d\n", w.Id, fib)
			case <-w.QuitChannel:
				fmt.Printf("El trabajo con id %d  se detuvo\n", w.Id)

			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChannel <- true
	}()
}

func NewDispatcher(jobQueue chan Job, maxWorker int) *Dispatcher {
	worker := make(chan chan Job, maxWorker)
	return &Dispatcher{
		WorkerPool: worker,
		MaxWorker:  maxWorker,
		JobQueue:   jobQueue,
	}
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				WorkerJobQueue := <-d.WorkerPool
				WorkerJobQueue <- job
			}()
		}
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorker; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()

	}
	go d.Dispatch()
}

func Fibonnaci(num int) int {
	if num <= 1 {
		return num
	}
	return Fibonnaci(num-1) + Fibonnaci(num-2)
}

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	//Si el request no es post
	if r.Method != "POST" {
		w.Header().Set("Permitir", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	//Extraer
	//Validaciones
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Delay invalido", http.StatusBadRequest)
		return
	}
	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Valor invalido", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "name invalido", http.StatusBadRequest)
		return
	}
	job := Job{Name: name, Delay: delay, Number: value}
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)
}

func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":8081"
	)
	//Crear un channel donde se metera la cola de trabajos
	JobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(JobQueue, maxWorkers)

	dispatcher.Run()
	//Creando el servidor
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, JobQueue)
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
