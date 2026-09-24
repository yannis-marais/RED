package ProjetRED

type StatusEffect struct {
	Name     string
	Damage   int
	Duration int
	Interval int
	TimeLeft int
}

type Skill struct {
	Name      string
	Damage    int
	Reiki     int
	Heal      int
	Type      string
	Strength  int
	précision int
}

type Character struct {
	Nom      string
	Classe   Classe
	LVL      int
	XP       float64
	PVMax    int
	PV       int
	Strength int
	Defense  int
	Reiki    int
	Spd      int

	Weapon string
	Helmet string
	Armor  string
	Boots  string

	Purse     uint
	Inventory Inventory
	Effects   []StatusEffect

	Cooldowns map[string]int
	Skills    map[string]Skill
}

type Classe struct {
	Nom      string
	PVMax    int
	Strength int
	Defense  int
	Reiki    int
	Spd      int
	Weapon   string
	Helmet   string
	Armor    string
	Boots    string
}

const DefaultInventoryCapacity uint = 20

func CharacterCreation(nom string, classe Classe) Character {
	return Character{
		Nom:      Capitalize(nom),
		Purse:    100,
		Classe:   classe,
		LVL:      1,
		XP:       0,
		PVMax:    classe.PVMax,
		PV:       classe.PVMax,
		Strength: classe.Strength,
		Defense:  classe.Defense,
		Reiki:    classe.Reiki,
		Spd:      classe.Spd,
		Weapon:   classe.Weapon,
		Helmet:   classe.Helmet,
		Armor:    classe.Armor,
		Boots:    classe.Boots,
		Inventory: Inventory{
			Capacity:    DefaultInventoryCapacity,
			Items:       make(map[string]int),
			Consumables: make(map[string]int),
			Materials:   make(map[string]int),
			SkillBooks:  make(map[string]int),
		},
		Effects:   []StatusEffect{},
		Cooldowns: make(map[string]int),
		Skills:    map[string]Skill{},
	}
}

type Inventory struct {
	Capacity    uint
	Items       map[string]int
	Consumables map[string]int
	Materials   map[string]int
	SkillBooks  map[string]int
}

func (inventory Inventory) UsedSlots() uint {
	var used uint
	seen := make(map[string]struct{})

	for _, items := range []map[string]int{
		inventory.Items,
		inventory.Consumables,
		inventory.Materials,
		inventory.SkillBooks,
	} {
		for name, quantity := range items {
			if quantity > 0 {
				seen[name] = struct{}{}
			}
		}
	}

	used = uint(len(seen))
	return used
}

func (inventory Inventory) HasFreeSlot(name string) bool {
	if inventory.hasItem(name) {
		return true
	}
	return inventory.UsedSlots() < inventory.Capacity
}

func (inventory Inventory) hasItem(name string) bool {
	for _, items := range []map[string]int{
		inventory.Items,
		inventory.Consumables,
		inventory.Materials,
		inventory.SkillBooks,
	} {
		if items[name] > 0 {
			return true
		}
	}
	return false
}
