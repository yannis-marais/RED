package ProjetRED

import (
	"fmt"

	personnage "ProjetRED/Personnage"
)

type ItemType string

const (
	Helmet ItemType = "Helmet"
	Armor  ItemType = "Armor"
	Boots  ItemType = "Boots"
	Weapon ItemType = "Weapon"
)

type Item struct {
	Name       string
	BonusPV    int
	BonusAtk   int
	BonusDef   int
	BonusReiki int
	BonusSpeed int
	MaxStack   int
	Type       ItemType
}
type Material struct {
	Name     string
	MaxStack int
}

type Consumable struct {
	Name     string
	Heal     int
	MaxStack int
}

func AddItem(p personnage.Character, item Item) {
	if p.Inventory.Items == nil {
		p.Inventory.Items = make(map[string]int)
	}

	current := p.Inventory.Items[item.Name]
	if current < item.MaxStack {
		p.Inventory.Items[item.Name] = current + 1
		fmt.Println(item.Name, "ajouté ! Quantité :", current+1)
		return
	}

	fmt.Println("Impossible : stack maximum atteint pour", item.Name)
}

func AddMaterial(p *personnage.Character, name string) bool {
	if p == nil {
		return false
	}

	material, ok := Materials[name]
	if !ok {
		return false
	}
	if p.Inventory.Materials == nil {
		p.Inventory.Materials = make(map[string]int)
	}

	current := p.Inventory.Materials[name]
	if current >= material.MaxStack {
		return false
	}

	p.Inventory.Materials[name] = current + 1
	return true
}

var Items = map[string]Item{
	//T1 Ronin
	"Bandit_Helmet": {Name: "Bandit Helmet", BonusDef: 10, BonusPV: 15, MaxStack: 1, Type: "Helmet"},
	"Bandit_Armor":  {Name: "Bandit Armor", BonusDef: 12, BonusPV: 15, MaxStack: 1, Type: "Armor"},
	"Bandit_Boots":  {Name: "Bandit Boots", BonusDef: 7, BonusPV: 10, BonusSpeed: 10, MaxStack: 1, Type: "Boots"},
	"Bandit_Spear":  {Name: "Bandit Spear", BonusDef: 2, BonusAtk: 10, MaxStack: 1, Type: "Weapon"},
	//T2 Ronin & T1 Cuirrasé
	"Samourai_Helmet": {Name: "Knight Helmet", BonusDef: 20, BonusPV: 25, MaxStack: 1, Type: "Helmet"},
	"Samourai_Armor":  {Name: "Knight Armor", BonusDef: 20, BonusPV: 25, MaxStack: 1, Type: "Armor"},
	"Samourai_Boots":  {Name: "Knight Boots", BonusDef: 15, BonusPV: 15, BonusSpeed: 20, MaxStack: 1, Type: "Boots"},
	// T2 Cuirassé
	"Lord_Helmet":     {Name: "Lord Helmet", BonusDef: 30, BonusPV: 25, MaxStack: 1, Type: "Helmet"},
	"Lord_Armor":      {Name: "Lord Armor", BonusDef: 40, BonusPV: 35, MaxStack: 1, Type: "Armor"},
	"Lord_Boots":      {Name: "Lord Boots", BonusDef: 20, BonusPV: 20, BonusSpeed: 10, MaxStack: 1, Type: "Boots"},
	"Lord_Battle_Axe": {Name: "Lord Battle Axe", BonusDef: 10, BonusAtk: 20, MaxStack: 1, Type: "Weapon"},
	// T1 Mage
	"Mage_Staff": {Name: "Mage Staff", BonusAtk: 20, MaxStack: 1, Type: "Weapon"},
	"Mage_Hood":  {Name: "Mage Hood", BonusReiki: 20, BonusPV: 10, MaxStack: 1, Type: "Helmet"},
	"Mage_Robe":  {Name: "Mage Robe", BonusReiki: 30, BonusPV: 10, MaxStack: 1, Type: "Armor"},
	"Mage_Boots": {Name: "Mage Boots", BonusReiki: 10, BonusPV: 10, MaxStack: 1, Type: "Boots"},
	//T2 Mage
	"Archimage_Rings": {Name: "Archimage Rings", BonusAtk: 30, BonusReiki: 50, MaxStack: 1, Type: "Weapon"},
	"Archimage_Hood":  {Name: "Archimage Hood", BonusPV: 10, BonusAtk: 15, BonusReiki: 30, MaxStack: 1, Type: "Helmet"},
	"Archimage_Robe":  {Name: "Archimage Robe", BonusPV: 10, BonusAtk: 20, BonusReiki: 40, MaxStack: 1, Type: "Armor"},
	"Archimage_Boots": {Name: "Archimage Boots", BonusPV: 10, BonusAtk: 10, BonusReiki: 20, MaxStack: 1, Type: "Boots"},
	//Armes T1
	"Katana":             {Name: "Katana", BonusAtk: 10, BonusDef: 10, MaxStack: 1, Type: "Weapon"},
	"Dagger":             {Name: "Dagger", BonusAtk: 15, BonusSpeed: 5, MaxStack: 1, Type: "Weapon"},
	"Swordshield":        {Name: "Swordshield", BonusAtk: 5, BonusDef: 10, MaxStack: 1, Type: "Weapon"},
	"Elementalist_Rings": {Name: "Elementalist Rings", BonusAtk: 10, BonusReiki: 10, MaxStack: 1, Type: "Weapon"},
	//Armes T2
	"Huge_Cleaver":   {Name: "Huge Cleaver", BonusAtk: 20, BonusDef: 5, BonusPV: 10, MaxStack: 1, Type: "Weapon"},
	"Odachi":         {Name: "Odachi", BonusAtk: 25, BonusSpeed: 5, MaxStack: 1, Type: "Weapon"},
	"Demonium_Staff": {Name: "Demonium Staff", BonusAtk: 30, BonusReiki: 25, BonusPV: 5, MaxStack: 1, Type: "Weapon"},
	// Craft (le stuff op)
	"Maximilian_Helmet": {Name: "Maximilian Helmet", BonusPV: 25, BonusSpeed: 5, BonusDef: 40, MaxStack: 1, Type: "Helmet"},
	"Maximilian_Armor":  {Name: "Maximilian Armor", BonusPV: 50, BonusAtk: 25, BonusDef: 60, MaxStack: 1, Type: "Armor"},
	"Maximilian_Boots":  {Name: "Maximilian Boots", BonusPV: 30, BonusSpeed: 25, BonusDef: 40, MaxStack: 1, Type: "Boots"},
	"La Maxime":         {Name: "La Maxime", BonusAtk: 50, BonusSpeed: 20, BonusPV: 50, MaxStack: 1, Type: "Weapon"},
}

var HealingPotion = Consumable{
	Name:     "Healing Potion",
	Heal:     50,
	MaxStack: 5,
}
var PoisonDOTPotion = Consumable{
	Name:     "Potion de Poison",
	Heal:     0,
	MaxStack: 3,
}
var Pain = Consumable{
	Name:     "Pain",
	Heal:     15,
	MaxStack: 10,
}

var Materials = map[string]Material{
	"Fer":     {Name: "Fer", MaxStack: 64},
	"Bois":    {Name: "Bois", MaxStack: 64},
	"Cuir":    {Name: "Cuir", MaxStack: 64},
	"Cristal": {Name: "Cristal", MaxStack: 64},
	"Diamant": {Name: "Diamant", MaxStack: 64},
}

var ConsumablesRegistry = map[string]Consumable{
	HealingPotion.Name:   HealingPotion,
	PoisonDOTPotion.Name: PoisonDOTPotion,
	Pain.Name:            Pain,
}

func GetConsumableByName(name string) (Consumable, bool) {
	item, ok := ConsumablesRegistry[name]
	return item, ok
}
