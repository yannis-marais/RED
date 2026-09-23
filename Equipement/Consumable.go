package ProjetRED

import (
	personnage "ProjetRED/Personnage"
	"fmt"
)

func AddConsumable(p *personnage.Character, item Consumable) {
	if p.Inventory.Consumables == nil {
		p.Inventory.Consumables = make(map[string]int)
	}

	current := p.Inventory.Consumables[item.Name]
	if current < item.MaxStack {
		p.Inventory.Consumables[item.Name] = current + 1
		fmt.Println(item.Name, "ajoutée ! Quantité :", current+1)
		return
	}

	fmt.Println("Impossible : stack maximum atteint pour", item.Name)
}

func UseConsumable(p *personnage.Character, item Consumable) {
	if p.Inventory.Consumables == nil {
		p.Inventory.Consumables = make(map[string]int)
	}

	qty := p.Inventory.Consumables[item.Name]
	if qty <= 0 {
		fmt.Println("Tu n'as pas de", item.Name)
		return
	}

	if item.Heal > 0 && p.PV == p.PVMax {
		fmt.Println("Impossible d'utiliser", item.Name, ": PV déjà au maximum.")
		return
	}

	if item.Name == "Poison DOT Potion" {
		effect := personnage.StatusEffect{
			Name:     "Poison",
			Damage:   10,
			Duration: 3,
			Interval: 1,
			TimeLeft: 3,
		}
		p.Effects = append(p.Effects, effect)
		fmt.Println(p.Nom, "est empoisonné !")
	}

	if item.Heal != 0 {
		p.PV += item.Heal
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
		if p.PV < 0 {
			p.PV = 0
		}
		p.Inventory.Consumables[item.Name] = qty - 1
	}
}
