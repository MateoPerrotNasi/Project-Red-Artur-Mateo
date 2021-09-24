package main

import (
	"fmt"
	"time"
)

func main() {
	// fonction qui execute nos sous fonctions et rentre les valeur ainsi que le menu principal
	var p1 Personnage
	p1.CharCreation()
	fmt.Printf("Vous allez être redirigé vers le menu principal")
	time.Sleep(1 * time.Second)
	fmt.Printf(". ")
	time.Sleep(1 * time.Second)
	fmt.Printf(". ")
	time.Sleep(1 * time.Second)
	fmt.Printf(". \n")
	time.Sleep(1 * time.Second)
	p1.menu()
}

func (p *Personnage) CharCreation() {
	var name string
	var class string
	var level int
	var lpmax int
	var lp int
	var inventory []string
	var skill []string

	fmt.Printf("Bienvenue dans le menu de création de personnage \nPour commencer, choisissez un nom pour votre avatar: \n")
	fmt.Scanln(&name)
	p.Capitalize(name)
	fmt.Println("Votre personnage se nomme désormais", name)
	time.Sleep(3 * time.Second)
	fmt.Println("\n Choisissez maintenant la race de ",name, "parmi:\n-Humain \n-Elfe \n-Nain")
	fmt.Scanln(&class)
	if class != "Humain" && class != "Elfe" && class != "Nain" {
		fmt.Printf("Erreur, veuillez entrer une valeur correcte:\n Humain Nain ou Elfe (n'oubliez pas la majuscule)")
		fmt.Scanln(&class)
	}
	switch class {
		case "Humain":
			class = "Humain"
		case "Elfe":
			class = "Elfe"
		case "Nain":
			class = "Nain"
	}
	if class == "Humain" {
		lpmax = 100
	} else if p.class == "Elfe" {
		lpmax = 80
	} else if p.class == "Nain" {
		lpmax = 120
	}
	lp = lpmax / 2
	fmt.Println("Vous avez choisi", class, ", vous commencez donc avec", lp, "/", lpmax, "point de vie")
	fmt.Println("Vous êtes niveau 1 et possédez le sort Coupe de poing")
	level = 1
	p.Init(name, class, level, lpmax, lp, inventory, skill)
}

func (p *Personnage) menu() {
	var menu int
	fmt.Println("+++++++++++++++++++++++++++++++")
	fmt.Println("A quoi voulez vous accéder:")
	fmt.Println("----- \n Afficher les informations du personnage (1)")
	fmt.Println("----- \n Accéder au contenu de l’inventaire (2)")
	fmt.Println("----- \n Voir le Marchand (3)")
	fmt.Println("----- \n Quitter (4) \n-----")
	fmt.Println("Entrez le numéro de l'option:")
	fmt.Println("+++++++++++++++++++++++++++++++")
	fmt.Scanln(&menu)
	switch menu {
	case 1:
		p.DisplayInfo()
	case 2:
		p.AccessInventory()
	case 3:
		p.Marchand()
	case 4:
		fmt.Println("Fin de la transmission")
		break
	}
}

type Personnage struct {
	// creation de la structure de notre personnage
	name      string
	class     string
	level     int
	lpmax     int
	lp        int
	inventory []string
	skill     []string
}

func (p *Personnage) Init(name string, class string, level int, lpmax int, lp int, inventory []string, skill []string) {
	// initialisation de notre personnage
	p.name = name
	p.class = class
	p.level = level
	p.lpmax = lpmax
	p.lp = lp
	p.inventory = inventory
	p.skill = skill
}

func (p *Personnage) AccessInventory() {
	// fonction qui nous permet d'acceder a notre inventaire
	// var rep int
	if len(p.inventory) == 0 {
		fmt.Println("inventaire vide fraté")
	} else {
		fmt.Println(p.inventory)
	}
	for i := 0; i < len(p.inventory); i++ {
		fmt.Println("---]", p.inventory[i], "[---")
	}
	p.retour()
}

func (p *Personnage) retour() {
	// fonction qui nous permet de retourner au menu précédent
	var rep int
	fmt.Println("tapez 1 pour retourner zo menu précédent")
	fmt.Scanln(&rep)
	if rep == 1 {
		p.menu()
	}
}

func (p *Personnage) spellbook(item string) {
	// fonction qui nous permet d'ajouter ou repertorier les sorts (spell)
	h := &p.skill
	for _, letter := range p.skill {
		if letter == ("Boule de feu") {
			fmt.Println("dsl t'a déja les boules")
		} else {
			*h = append(*h, item)
		}
	}
	p.retour()
}

func (p *Personnage) Marchand() {
	// fonction affichant le menu du marchand , et les ajoute a notre inventaire
	var menum int
	fmt.Println("-----------------Marchand-------------------")
	fmt.Println("Tapez 1 pour obtenir une Potion de vie ;)")
	fmt.Println("Tapez 2 pour obtenir une Potion de poison ;(")
	fmt.Println("Tapez 3 pour tenter d'obtenir boule de feu ")
	fmt.Println("Tapez 4 pour retourner au menu précédent ")
	fmt.Println("--------------------------------------------")
	fmt.Scanln(&menum)
	switch menum {
	case 1:
		p.AddInventory("Potion de vie")
		p.AccessInventory()
	case 2:
		p.AddInventory("Potion de poison")
		p.AccessInventory()
	case 3:
		p.spellbook("boule de feu")
	case 4:
		p.retour()
	}
}

func (p *Personnage) DisplayInfo() {
	// fonction nous permettant de voir les informations de notre personnage
	fmt.Println("-----------")
	fmt.Println("Nom:", p.name)
	fmt.Println("Classe:", p.class)
	fmt.Println("Niveau:", p.level)
	fmt.Println("Vie maximum:", p.lpmax)
	fmt.Println("Vie actuelle:", p.lp)
	fmt.Println("Contenu de l'inventaire:", p.inventory)
	fmt.Println("skill :", p.skill)
	fmt.Println("-----------")
	p.retour()

}

func (p *Personnage) AddInventory(item string) {
	p.inventory = append(p.inventory, item)
}

func (p *Personnage) TakePot() {
	// fonction qui nous permet de prendre une potion de soin
	for _, letter := range p.inventory {
		if letter == "Potion de vie" {
			if p.lp <= (p.lpmax - 50) {
				p.lp += 50
				p.inventory[len(p.inventory)-1] = ""
				break
			} else if p.lp > (p.lpmax-50) && p.lp < p.lpmax {
				p.lp = p.lpmax
				p.inventory[len(p.inventory)-1] = ""
				break
			} else {
				fmt.Println("Vous êtes full")
				break
			}
		}
	}
}

func (p *Personnage) Dead() {
	// fonction qui verifie si le personnage est mort et le ressussite a la moitié de ses pv
	if p.lp == 0 {
		fmt.Printf("Bravo, vous êtes mort. \n")
		time.Sleep(2 * time.Second)
		fmt.Printf("Mais ne paniquez pas, vous allez être ressuciter \n")
		time.Sleep(2 * time.Second)
		fmt.Printf("Manoeuvre de réanimation en cours")
		time.Sleep(1 * time.Second)
		fmt.Printf(". ")
		time.Sleep(1 * time.Second)
		fmt.Printf(". ")
		time.Sleep(1 * time.Second)
		fmt.Printf(". \n")
		time.Sleep(1 * time.Second)
		p.lp = p.lpmax / 2
		p.DisplayInfo()
	}
}

func (p *Personnage) PoisonPot() {
	// fonction qui crée la potion poison et explique ce qu'elle fait sur un personnage
	for _, letter := range p.inventory {
		if letter == "Potion de poison" {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(p.lp, "/", p.lpmax)
			p.lp -= 10
			fmt.Println(p.lp, "/", p.lpmax)
			time.Sleep(100 * time.Millisecond)
			p.lp -= 10
			fmt.Println(p.lp, "/", p.lpmax)
			time.Sleep(100 * time.Millisecond)
			p.lp -= 10
			fmt.Println(p.lp, "/", p.lpmax)
		}
	}
}

func prim(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

func (p *Personnage) Capitalize(s string) {
	n := &s
	ar := []rune(s)
	letra := true
	for i := 0; i < len(s); i++ {
		if prim(ar[i]) == true && letra {
			if ar[i] >= 'a' && ar[i] <= 'z' {
				ar[i] = 'A' - 'a' + ar[i]
			}
			letra = false
		} else if ar[i] >= 'A' && ar[i] <= 'Z' {
			ar[i] = 'a' - 'A' + ar[i]
		} else if prim(ar[i]) == false {
			letra = true
		}
	}
	*n = string(ar)
}
