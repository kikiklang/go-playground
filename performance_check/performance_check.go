package performance_check

import (
	"fmt"
	"sort"

	"github.com/jaswdr/faker"
)

func FormatPets(dataSetSize int) []string {
	fakePet := faker.New().Pet()
	petsArray := make([]string, dataSetSize)

	for i := 0; i < dataSetSize; i++ {
		petsArray[i] = fakePet.Name()
	}

	if !sort.StringsAreSorted(petsArray) {
		sort.Strings(petsArray)
		return removeDuplicate(petsArray)
	}

	return removeDuplicate(petsArray)

}

func removeDuplicate(petsArray []string) []string {
	originalLength := len(petsArray)
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range petsArray {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	afterFilteringLength := len(list)

	fmt.Printf("original length was: %d \n", originalLength)
	fmt.Printf("length after filtering: %d \n", afterFilteringLength)
	fmt.Printf("number of removed items: %d \n", originalLength-afterFilteringLength)

	return list
}
