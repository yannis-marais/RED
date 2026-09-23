package ProjetRED

type StatusEffect struct {
	Name     string
	Damage   int
	Duration int
	Interval int
	TimeLeft int
}

type Skills struct {
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

	Purse uint
	Inventory Inventory
	Effects   []StatusEffect

	Cooldowns map[string]int
	Skills    map[string]Skills
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

func CharacterCreation(nom string, classe Classe) Character {
	return Character{
		Nom:      Capitalize(nom),
		Purse:	  100,
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
			Items:       make(map[string]int),
			Consumables: make(map[string]int),
			Materials:   make(map[string]int),
			Skill:       make(map[string]int),
			SkillBooks:  make(map[string]int),
		},
		Effects: []StatusEffect{},
	}
}

type Inventory struct {
	Items       map[string]int
	Consumables map[string]int
	Materials   map[string]int
	Skill       map[string]int
	SkillBooks  map[string]int
}
