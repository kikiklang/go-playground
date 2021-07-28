package performance_check

import (
	"github.com/jaswdr/faker"
)

func Format5000Name(dataSetSize int) []string {
	faker := faker.New()
	person := faker.Person()
	var persons = make([]string, dataSetSize)

	for i := 0; i < dataSetSize; i++ {
		persons[i] = person.Name()
	}

	return persons
}

// 5.93316ms
