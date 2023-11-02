package facts

import (
	"fmt"
	"sync"
	"time"
)

func InitializeFacts(source FactsQueryService, destination FactsRepository, numberOfFacts int) {
	fmt.Println("Initalizing database...")
	t := time.Now()
	var wg sync.WaitGroup
	wg.Add(numberOfFacts)
	for i := 0; i < numberOfFacts; i++ {
		go func() {
			defer wg.Done()
			var fact, err = source.GetFact()
			if err == nil {
				destination.SaveFact(fact)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Done in %v facts added in %v.\n", numberOfFacts, time.Since(t))
}
