package main

import (
	"fmt"
	"sync"
	"time"
)

// Definimos el tipo de trabajo
type trabajo struct {
	ID int
	X  int
}

// Definimos el tipo de resultado
type resultado struct {
	ID        int
	X         int
	Procesado int
}

// Función worker
func worker(id int, jobs <-chan trabajo, results chan<- resultado, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		// Simula procesamiento
		time.Sleep(100 * time.Millisecond)

		r := resultado{
			ID:        j.ID,
			X:         j.X,
			Procesado: j.X * 2, // ejemplo de procesamiento
		}
		fmt.Printf("[worker %d] procesa trabajo %d -> %d\n", id, j.ID, r.Procesado)
		results <- r
	}
	fmt.Printf("[worker %d] no hay más trabajos\n", id)
}

func main() {
	nWorkers := 3
	nTrabajos := 5

	jobs := make(chan trabajo, nTrabajos)
	results := make(chan resultado, nTrabajos)

	var wg sync.WaitGroup

	// Lanzar workers
	wg.Add(nWorkers)
	for i := 1; i <= nWorkers; i++ {
		go worker(i, jobs, results, &wg)
	}

	// Productor de trabajos
	go func() {
		for i := 1; i <= nTrabajos; i++ {
			j := trabajo{ID: i, X: i * 10}
			fmt.Printf("[main] produce trabajo %d\n", j.ID)
			jobs <- j
			time.Sleep(50 * time.Millisecond) // simula tiempo entre producciones
		}
		close(jobs) // importante: cerrar para que los workers terminen
	}()

	// Goroutine para cerrar results cuando todos los workers terminen
	go func() {
		wg.Wait()
		close(results)
	}()

	// Consumidor de resultados
	for r := range results {
		fmt.Printf("[main] resultado recibido: %+v\n", r)
	}
}