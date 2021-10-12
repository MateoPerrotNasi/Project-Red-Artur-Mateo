package main

import (
	"os"
	"fmt"
	"time"
)

func (p *Personnage) QuestMan(e3, e4, e5, e6 *Monstre) {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	var choice int
	if !p.Checkinv("Objet Suspicieux") {
		if !p.Checkinv("Véritable Couteau") {
				Slow(Yellow+"Bienvenue sur la Place Principale l'ami.\n"+Reset, 1)
				Slow("(1) Lui demander son nom ?\n", 1)
				Slow("(2) Lui demander ce qu'on fait la ?\n", 1)
				fmt.Scanln(&choice)
				switch choice {
				case 1:
					Slow(Yellow+"Mon nom importe peu l'ami, je ne .", 1)
					Slow("\nBonne chance l'ami.\n"+Reset, 1)
				case 2:
					Slow(Yellow+"Vois tu depuis peu les contrées de Goldy se sont fait envahir par quatres seigneur malveillants.", 1)
					Slow("\nTa quête est de les éliminer, un par un et jusqu'au dernier.\n", 1)
					Slow("\nMais avant cela tu dois te rendre dans le donjon afin d'y trouver l'arme qui te donnera la force de les battre.", 1)
					Slow("\nBonne chance l'ami.\n"+Reset, 1)
				}
		} else {
			Slow(Yellow+"Bravo l'ami ! Tu as obtenus l'artéfact.", 1)
			Slow("\nJe te conseille maintenant d'aller t'équiper en conséquence et de te lancer à la recherche des Quatres Grands"+Reset, 1)
		}
		if e3.lp <= 0 {
			if e4.lp <= 0 {
				if e5.lp <= 0 {
					Slow(Yellow + "Bien joué l'ami ! Tu as vaincu ",1)
					Slow(e5.name,1)
					Slow(", il ne t'en reste plus qu'un'" + Reset,1)
				}
				Slow(Yellow + "Bien joué l'ami ! Tu as vaincu ",1)
				Slow(e4.name,1)
				Slow(", il ne t'en reste plus que deux" + Reset,1)
			}
			Slow(Yellow + "Bien joué l'ami ! Tu as vaincu ",1)
			Slow(e3.name,1)
			Slow(", il ne t'en reste plus que trois" + Reset,1)
		}
	}
	time.Sleep(2 * time.Second)
}