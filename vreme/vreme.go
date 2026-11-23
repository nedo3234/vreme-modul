// Package vreme omogoča upravljanje vremenskih postaj in njihovih meritev.
package vreme

import (
	"fmt"
)

// Postaja predstavlja vremensko postajo z meritvami temperature.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajMeritev postaji z dano oznako doda meritev, če je med minTemp in maxTemp.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	postaja, ok := postaje[oznaka]

	if !ok {
		fmt.Printf("Postaja z oznako %s ne obstaja.\n", oznaka)
		return
	} else if temperatura < minTemp || temperatura > maxTemp {
		fmt.Println("Meritev neveljavna.")
		return
	} else {
		postaja.meritve = append(postaja.meritve, temperatura)
		postaje[oznaka] = postaja
		return
	}

}

func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, ok := postaje[oznaka]

	if len(postaja.meritve) < stMeritev {
		ok = false
	}

	if !ok {
		return
	}

	var suma float64 = 0
	najnizja = 51
	najvisja = -51
	for _, meritev := range postaja.meritve {
		if najnizja > meritev {
			najnizja = meritev
		} else if najvisja < meritev {
			najvisja = meritev
		}
		suma += meritev
	}
	povprecna = suma / float64(len(postaja.meritve))

	return
}

// IzpisMeritev izpiše vse postaje in njihove meritve.
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for oznaka, postaja := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", oznaka, postaja.ime, postaja.kraj, postaja.meritve)
	}
}

// IzpisPorocila izpiše statistiko za vsako postajo, če ima vsaj stMeritev meritev.
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for oznaka, postaja := range postaje {
		najnizja, najvisja, povprecna, ok := statistika(postaje, oznaka, stMeritev)
		switch {
		case !ok:
			fmt.Printf("%s: premalo meritev\n", postaja.ime)
		case povprecna >= 25:
			fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> Vroče\n", postaja.ime, najnizja, najvisja, povprecna)
		case povprecna >= 15:
			fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> Toplo\n", postaja.ime, najnizja, najvisja, povprecna)
		case povprecna >= 5:
			fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> Hladno\n", postaja.ime, najnizja, najvisja, povprecna)
		default:
			fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> Mrzlo\n", postaja.ime, najnizja, najvisja, povprecna)
		}
	}
}

// DodajPostajo v slovar doda novo postajo z dano oznako, imenom in krajem.
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{ime: ime, kraj: kraj, meritve: []float64{}}
}
