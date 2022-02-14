package CustomerRepository

import (
	"Packages/src/api/Type/EntityTypes"
	"fmt"
	"math/rand"
)

func (r Repository) CreateInitialData(DocumentCount int) {
	for i := 0; i < DocumentCount; i++ {
		name := names[rand.Intn(49)]
		surname := names[rand.Intn(49)]
		r.Create(&EntityTypes.User{
			Name:     fmt.Sprintf("%v %v", name, surname),
			Email:    fmt.Sprintf("%v@%v.com", name, surname),
			Age:      rand.Intn(65) + 15,
			Document: EntityTypes.Document{},
		})
	}

}

var names = [...]string{
	"Nashla",
	"Zuri",
	"Vesper",
	"Mahalia",
	"Bianca",
	"Shaurya",
	"Raj",
	"Demon",
	"Link",
	"Mesa",
	"Kofi",
	"Shlomo",
	"Rick",
	"Rafe",
	"Nellie",
	"Akai",
	"Shoshana",
	"Yechezkel",
	"Elijah",
	"James",
	"Teal",
	"Rozlynn",
	"Ellison",
	"Lynley",
	"Shaan",
	"Margo",
	"Jamilah",
	"Joniel",
	"Abril",
	"Aira",
	"Nalah",
	"Nada",
	"Braelee",
	"Erabella",
	"Elie",
	"Merry",
	"Ella",
	"Irena",
	"Kayleb",
	"Samaira",
	"Jonathon",
	"Korben",
	"Mikah",
	"Rosalee",
	"Mikhail",
	"Olen",
	"Tien",
	"Zayda",
	"Mikey",
	"Omere",
}
