package main

import (
	"context"
	"fmt"
	"log"
	"os"

	redovalnica "github.com/mattbernot/vpsa/redovalnica"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "Redovalnica",
		Usage: "Redovalnica za vpis ocen študentov.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "Najmanjše število ocen potrebnih za pozitivno oceno",
				Value: 6,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "Najmanjša možna ocena",
				Value: 0,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "Največja možna ocena",
				Value: 10,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {

			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")
			stOcen := cmd.Int("stOcen")

			studenti := make(map[string]redovalnica.Student)

			studenti["6240050"] = redovalnica.Student{"Janko", "Bananko", []int{5, 6, 2, 7, 8, 10}}
			studenti["6240051"] = redovalnica.Student{"Jana", "Banana", []int{4, 8, 7, 7, 9}}
			studenti["6240052"] = redovalnica.Student{"Dejan", "Smrkolj", []int{9, 9, 10, 9, 10, 8}}

			redovalnica.DodajOceno(studenti, "1234567", 8, minOcena, maxOcena)
			redovalnica.DodajOceno(studenti, "6240050", 5, minOcena, maxOcena)
			redovalnica.IzpisRedovalnice(studenti)
			fmt.Println()
			redovalnica.IzpisiKoncniUspeh(studenti, stOcen)

			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
