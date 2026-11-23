package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/nedo3234/vreme-modul/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "vreme",
		Usage: "Upravljanje vremenskih postaj",
		Flags: []cli.Flag{
			&cli.FloatFlag{Name: "minTemp", Value: -50, Usage: "najnižja dovoljena temperatura meritev"},
			&cli.FloatFlag{Name: "maxTemp", Value: 50, Usage: "najvišja dovoljena temperatura meritve"},
			&cli.IntFlag{Name: "stMeritev", Value: 3, Usage: "najmanjše število meritev za statistiko"},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			minTemp := c.Float("minTemp")
			maxTemp := c.Float("maxTemp")
			stMeritev := int(c.Int("stMeritev"))

			postaje := make(map[string]vreme.Postaja)
			vreme.DodajPostajo(postaje, "LJ", "Ljubljana", "Ljubljana")
			vreme.DodajPostajo(postaje, "MB", "Maribor", "Maribor")

			for _, t := range []float64{12, 18, 21, 99} {
				vreme.DodajMeritev(postaje, "LJ", t, minTemp, maxTemp)
			}
			vreme.DodajMeritev(postaje, "MB", 3, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "KP", 10, minTemp, maxTemp)

			vreme.IzpisMeritev(postaje)
			fmt.Println()
			vreme.IzpisPorocila(postaje, stMeritev)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
