package main

import "fmt"

type Pistol struct {
	model            string
	caliber          string
	magazineCapacity int
	roundsInMagazine int
}

func NewPistol(model, caliber string, magazineCapacity int) Pistol {
	return Pistol{
		model:            model,
		caliber:          caliber,
		magazineCapacity: magazineCapacity,
		roundsInMagazine: 0,
	}
}

func (p *Pistol) Info() string {
	return fmt.Sprintf("Model: %s, Caliber: %s, Capacity: %d, Rounds: %d",
		p.model, p.caliber, p.magazineCapacity, p.roundsInMagazine)
}

func (p *Pistol) Reload(rounds int) {
	if rounds < 0 {
		rounds = 0
	}
	if rounds > p.magazineCapacity {
		rounds = p.magazineCapacity
	}
	p.roundsInMagazine = rounds
}

func (p *Pistol) Shoot() bool {
	if p.roundsInMagazine == 0 {
		return false
	}
	p.roundsInMagazine--
	return true
}

func main() {
	pistol := NewPistol("Glock 17", "9mm", 17)

	fmt.Println(pistol.Info())

	pistol.Reload(17)
	fmt.Println("After reload:", pistol.Info())

	for i := 1; i <= 3; i++ {
		if pistol.Shoot() {
			fmt.Printf("Shot %d fired. %s\n", i, pistol.Info())
		}
	}
}
