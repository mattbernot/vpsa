package redovalnica

import "fmt"

type Student struct {
	ime     string
	priimek string
	ocene   []int
}

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	if ocena <= 0 || ocena >= 10 {
		fmt.Println("Vpisana ocena ni med 0 in 10")
		return
	}
	_, ok := studenti[vpisnaStevilka]
	if ok == false {
		fmt.Println("Študenta", vpisnaStevilka, "ni na seznamu")
		return
	}

	sez_ocen := studenti[vpisnaStevilka]
	sez_ocen.ocene = append(studenti[vpisnaStevilka].ocene, ocena)
	studenti[vpisnaStevilka] = sez_ocen

}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	_, ok := studenti[vpisnaStevilka]
	if ok == false {
		//fmt.Println("Študenta %d ni na seznamu", vpisnaStevilka)
		return -1.0
	}

	var st_ocen, sum int
	for i := range studenti[vpisnaStevilka].ocene {
		st_ocen++
		sum += studenti[vpisnaStevilka].ocene[i]
	}
	if st_ocen < 6 {
		return 0.0
	} else {
		return float64(sum) / float64(st_ocen)
	}

}

func IzpisRedovalnice(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisnaSt := range studenti {
		fmt.Println(vpisnaSt, studenti[vpisnaSt].ime, studenti[vpisnaSt].priimek, ": ", studenti[vpisnaSt].ocene)
	}
}

func IzpisiKoncniUspeh(studenti map[string]Student) {
	for vpisnaSt := range studenti {
		povp_ocena := povprecje(studenti, vpisnaSt)
		fmt.Printf("%s %s Povprečna ocena %.2f -> ", studenti[vpisnaSt].ime, studenti[vpisnaSt].priimek, povp_ocena)
		if povp_ocena >= 9 {
			fmt.Print("Odličen študent!")

		} else if povp_ocena < 9.0 && povp_ocena > 6.0 {
			fmt.Print("Povprečen študent")

		} else if povp_ocena < 6.0 {
			fmt.Print("Neuspešen študent")
		}
		fmt.Println()
	}
}
