# Redovalnica

Redovalnica je program za vpis ocen študentov.
Uporablja funkcije za:
- Dodajanje ocen študenta
- Izračunanje povprečja
- Izpis redovalnice
- Izpis končnega uspeha 

## Zahteve
Za izvajanje kode je potreben [programski jezik Go](https://go.dev/dl/),
ki je na voljo za Linux, Mac in Windows. 

## Primer uporabe

```go
studenti := make(map[string]redovalnica.Student)

studenti["6240050"] = redovalnica.Student{"Janko", "Bananko", []int{5, 6, 2, 7, 8, 10}}
studenti["6240051"] = redovalnica.Student{"Jana", "Banana", []int{4, 8, 7, 7, 9}}
studenti["6240052"] = redovalnica.Student{"Dejan", "Smrkolj", []int{9, 9, 10, 9, 10, 8}}

redovalnica.DodajOceno(studenti, "1234567", 8, minOcena, maxOcena)
redovalnica.DodajOceno(studenti, "6240050", 5, minOcena, maxOcena)
redovalnica.IzpisRedovalnice(studenti)
redovalnica.IzpisiKoncniUspeh(studenti, stOcen)
```
Program izpiše imena, priimke, ocene in vpisne številke študentov.
Zgornji primer izpiše:
```
Študenta 1234567 ni na seznamu
REDOVALNICA:
6240052 Dejan Smrkolj :  [9 9 10 9 10 8]
6240050 Janko Bananko :  [5 6 2 7 8 10 5]
6240051 Jana Banana :  [4 8 7 7 9]

Janko Bananko Povprečna ocena 6.14 -> Povprečen študent
Študent 6240051 ima manj kot 6 ocen.
Jana Banana Povprečna ocena 0.00 -> Neuspešen študent
Dejan Smrkolj Povprečna ocena 9.17 -> Odličen študent!
```

Kodo `main.go` ki se nahaja v direktoriju `cmd`, lahko izvedemo z ukazom:

```bash
go run cmd/main.go
```

## Godoc

Go ima tudi svojo interaktivno dokumentacijo Godoc.
Namestimo jo z ukazom:
```bash
go install golang.org/x/tools/cmd/godoc@latest
```
in nato poženemo ukaz:
```bash
godoc -http :8080
```
Dokumentacija modula Redovalnica bo na voljo na povezavi: [http://localhost:8080/pkg/github.com/mattbernot/vpsa/redovalnica/](http://localhost:8080/pkg/github.com/mattbernot/vpsa/redovalnica/)