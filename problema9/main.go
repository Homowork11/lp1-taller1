package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Implementar una versión del problema de los Filósofos Comensales.
// Hay 5 filósofos y 5 tenedores (recursos). Cada filósofo necesita 2 tenedores para comer.
// Estrategia segura: imponer un **orden global** al tomar los tenedores (primero el menor ID, luego el mayor)
// para evitar deadlock. También puedes limitar concurrencia (ej. mayordomo).
// TODO: completa la lógica de toma/soltado de tenedores y bucle de pensar/comer.

type tenedor struct{ mu sync.Mutex }

func filosofo(id int, izq, der *tenedor, wg *sync.WaitGroup) {
	defer wg.Done()
	// TODO: desarrolla el código para el filósofo que piensa, toma tenedores en orden seguro, come y luego suelta los tenedores.

	fmt.Printf("[filósofo %d] satisfecho\n", id)
}

func pensar(id int) {
	fmt.Printf("[filósofo %d] pensando...\n", id)

	// TODO: simular tiempo de pensar
	time.Sleep(time.Duration(1+id) * time.Second) // ejemplo: cada filósofo piensa un tiempo distinto
}

func comer(id int) {
	fmt.Printf("[filósofo %d] COMIENDO\n", id)
	// TODO: simular tiempo de pensar
	time.Sleep(time.Duration(1+id) * time.Second) // ejemplo: cada filósofo come un tiempo distinto
}

func main() {
	const n = 5
	var wg sync.WaitGroup
	wg.Add(n)

	// crear tenedores
	forks := make([]*tenedor, n)
	for i := 0; i < n; i++ {
		// TODO: inicializar cada tenedor i
		forks[i] = &tenedor{}
		var izq, der *tenedor
		if i < n-1 {
			izq, der = forks[i], forks[i+1]
		} else {
			izq, der = forks[i], forks[0]
		}
		go filosofo(i, izq, der, &wg)

	}

	// lanzar filósofos (ya están lanzadas en el loop anterior)

	wg.Wait()
	fmt.Println("Todos los filósofos han comido sin deadlock.")
}