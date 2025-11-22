package redovalnica

import "fmt"

type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int, minOcena int, maxOcena int) {
	if ocena <= minOcena || ocena >= maxOcena {
		fmt.Println("Vpisana ocena ni med 0 in 10")
		return
	}
	_, ok := studenti[vpisnaStevilka]
	if ok == false {
		fmt.Println("Študenta", vpisnaStevilka, "ni na seznamu")
		return
	}

	sez_ocen := studenti[vpisnaStevilka]
	sez_ocen.Ocene = append(studenti[vpisnaStevilka].Ocene, ocena)
	studenti[vpisnaStevilka] = sez_ocen

}

func povprecje(studenti map[string]Student, vpisnaStevilka string, stOcen int) float64 {
	_, ok := studenti[vpisnaStevilka]
	if ok == false {
		//fmt.Println("Študenta %d ni na seznamu", vpisnaStevilka)
		return -1.0
	}

	var st_ocen, sum int
	for i := range studenti[vpisnaStevilka].Ocene {
		st_ocen++
		sum += studenti[vpisnaStevilka].Ocene[i]
	}
	if st_ocen < stOcen {
		return 0.0
	} else {
		return float64(sum) / float64(st_ocen)
	}

}

func IzpisRedovalnice(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisnaSt := range studenti {
		fmt.Println(vpisnaSt, studenti[vpisnaSt].Ime, studenti[vpisnaSt].Priimek, ": ", studenti[vpisnaSt].Ocene)
	}
}

func IzpisiKoncniUspeh(studenti map[string]Student, stOcen int) {
	for vpisnaSt := range studenti {
		povp_ocena := povprecje(studenti, vpisnaSt, stOcen)

		if len(studenti[vpisnaSt].Ocene) < stOcen {
			fmt.Println("Študent", vpisnaSt, "ima manj kot 6 ocen.")
		}
		fmt.Printf("%s %s Povprečna ocena %.2f -> ", studenti[vpisnaSt].Ime, studenti[vpisnaSt].Priimek, povp_ocena)
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
