package ProjetRED

import (
	"fmt"

	personnage "ProjetRED/Personnage"
)

func FindItemByName(itemName string) (string, Item, bool) {
	if item, ok := Items[itemName]; ok {
		return itemName, item, true
	}

	for key, item := range Items {
		if item.Name == itemName {
			return key, item, true
		}
	}

	return "", Item{}, false
}

// slotDuPersonnage renvoie un pointeur vers le champ du personnage qui
// correspond au type d'objet donné (Weapon -> &p.Weapon, etc).
// Un pointeur permet de lire ET d'écrire le slot sans dupliquer le switch.
func slotDuPersonnage(p *personnage.Character, t ItemType) (*string, bool) {
	switch t {
	case Weapon:
		return &p.Weapon, true
	case Armor:
		return &p.Armor, true
	case Boots:
		return &p.Boots, true
	case Helmet:
		return &p.Helmet, true
	default:
		return nil, false
	}
}

// EquipItem équipe l'objet "itemName" sur le personnage p.
// Le type (Weapon/Armor/Boots/Helmet) est déduit automatiquement de l'objet
// trouvé dans le catalogue Items : plus besoin (et plus de risque d'erreur)
// à le préciser soi-même à l'appel.
//
// Comportement :
//  1. on cherche l'objet dans le catalogue (par clé ou par Name)
//  2. on vérifie que le joueur le possède bien dans son inventaire
//  3. si un objet était déjà équipé sur ce slot, il retourne dans l'inventaire
//  4. le nouvel objet est retiré de l'inventaire et équipé
//  5. les stats du personnage sont recalculées
func EquipItem(p *personnage.Character, itemName string) {
	if p == nil {
		return
	}

	itemKey, item, ok := FindItemByName(itemName)
	if !ok {
		fmt.Println("Objet introuvable :", itemName)
		return
	}

	slot, ok := slotDuPersonnage(p, item.Type)
	if !ok {
		fmt.Println("Type d'équipement inconnu :", item.Type)
		return
	}

	if p.Inventory.Items == nil {
		p.Inventory.Items = make(map[string]int)
	}

	qty := p.Inventory.Items[item.Name]
	if qty <= 0 {
		fmt.Println("Tu n'as pas", item.Name, "dans l'inventaire.")
		return
	}

	// L'objet est déjà équipé : rien à faire.
	if *slot == itemKey {
		fmt.Println(item.Name, "est déjà équipé.")
		return
	}

	// On range l'ancien équipement (s'il y en avait un) dans l'inventaire.
	if *slot != "" {
		if oldItem, ok := Items[*slot]; ok {
			AddItem(p, oldItem)
		}
	}

	// On retire le nouvel objet de l'inventaire.
	if qty <= 1 {
		delete(p.Inventory.Items, item.Name)
	} else {
		p.Inventory.Items[item.Name] = qty - 1
	}

	// On équipe : on stocke la clé du catalogue (ex: "Samourai_Helmet"),
	// pas item.Name, car RecalculateStats fait Items[p.Weapon] etc.
	*slot = itemKey

	RecalculateStats(p)
	fmt.Println(item.Name, "équipé.")
}