package ProjetRED

type LootEntry struct {
	Name string
	Rate int
}

type MONSTER struct {
	NOM      string
	PVMax    int
	PV       int
	PVMAXR   int
	PVR      int
	Strength int
	Defense  int
	Spd      int
	Reiki    int
	Loot     []LootEntry
}

func initGoblin() MONSTER {
	return MONSTER{
		NOM:      "gobelin",
		PVMax:    100,
		PV:       100,
		PVMAXR:   100,
		PVR:      100,
		Strength: 30,
		Defense:  20,
		Spd:      1,
		Reiki:    1,
		Loot: []LootEntry{
			{Name: "dague rouillée", Rate: 3},
			{Name: "bourse trouée", Rate: 2},
			{Name: "oreille de gobelin", Rate: 1},
			{Name: "torche éteinte", Rate: 2},
		},
	}
}

func initSkeleton() MONSTER {
	return MONSTER{
		NOM:      "squelette",
		PVMax:    35,
		PV:       35,
		PVMAXR:   35,
		PVR:      35,
		Strength: 20,
		Defense:  4,
		Spd:      20,
		Reiki:    2,
		Loot: []LootEntry{
			{Name: "tibia", Rate: 8},
			{Name: "crâne fissuré", Rate: 5},
			{Name: "épée rouillée", Rate: 20},
		},
	}
}

func initTroll() MONSTER {
	return MONSTER{
		NOM:      "troll",
		PVMax:    300,
		PV:       300,
		PVMAXR:   150,
		PVR:      150,
		Strength: 50,
		Defense:  40,
		Spd:      3,
		Reiki:    1,
		Loot: []LootEntry{
			{Name: "massue en bois", Rate: 6},
			{Name: "peau de troll", Rate: 12},
			{Name: "dent de troll", Rate: 4},
			{Name: "gourdin ébréché", Rate: 5},
		},
	}
}

func initVouivre() MONSTER {
	return MONSTER{
		NOM:      "vouivre",
		PVMax:    100,
		PV:       99,
		PVMAXR:   100,
		PVR:      99,
		Strength: 40,
		Defense:  8,
		Spd:      200,
		Reiki:    12,
		Loot: []LootEntry{
			{Name: "écaille de vouivre", Rate: 15},
			{Name: "griffe de vouivre", Rate: 10},
			{Name: "venin cristallisé", Rate: 20},
		},
	}
}

func initLoupGarou() MONSTER {
	return MONSTER{
		NOM:      "loup-garou",
		PVMax:    125,
		PV:       125,
		PVMAXR:   75,
		PVR:      75,
		Strength: 60,
		Defense:  40,
		Spd:      25,
		Reiki:    4,
		Loot: []LootEntry{
			{Name: "griffe de loup-garou", Rate: 14},
			{Name: "fourrure argentée", Rate: 18},
			{Name: "croc acéré", Rate: 9},
		},
	}
}

func initZombie() MONSTER {
	return MONSTER{
		NOM:      "zombie",
		PVMax:    50,
		PV:       50,
		PVMAXR:   75,
		PVR:      75,
		Strength: 12,
		Defense:  2,
		Spd:      4,
		Reiki:    1,
		Loot: []LootEntry{
			{Name: "chair putréfiée", Rate: 5},
			{Name: "os brisé", Rate: 8},
			{Name: "lambeau de tissu", Rate: 5},
			{Name: "anneau rouillé", Rate: 10},
		},
	}
}

func initOrc() MONSTER {
	return MONSTER{
		NOM:      "orc",
		PVMax:    150,
		PV:       150,
		PVMAXR:   100,
		PVR:      100,
		Strength: 30,
		Defense:  18,
		Spd:      16,
		Reiki:    1,
		Loot: []LootEntry{
			{Name: "hache d'orc", Rate: 11},
			{Name: "bouclier cabossé", Rate: 9},
			{Name: "défense d'orc", Rate: 6},
		},
	}
}

func initDragon() MONSTER {
	return MONSTER{
		NOM:      "dragon",
		PVMax:    1000,
		PV:       1000,
		PVMAXR:   1500,
		PVR:      1500,
		Strength: 100,
		Defense:  26,
		Spd:      10,
		Reiki:    18,
		Loot: []LootEntry{
			{Name: "écaille de dragon", Rate: 40},
			{Name: "griffe de dragon", Rate: 30},
			{Name: "souffle embouteillé", Rate: 50},
			{Name: "œuf de dragon", Rate: 100},
		},
	}
}

func initBrian() MONSTER {
	return MONSTER{
		NOM:      "Le B.",
		PVMax:    1,
		PV:       1,
		PVMAXR:   50000000,
		PVR:      50000000,
		Strength: 50,
		Defense:  2,
		Spd:      4000,
		Reiki:    18,
		Loot: []LootEntry{
			{Name: "LE BEDOU", Rate: 99},
		},
	}
}

func initAnt() MONSTER {
	return MONSTER{
		NOM:      "Le Ant.",
		PVMax:    10000000000,
		PV:       10000000000,
		PVMAXR:   1,
		PVR:      1,
		Strength: 50,
		Defense:  2,
		Spd:      4000,
		Reiki:    18,
		Loot: []LootEntry{
			{Name: "COMBI FOURMI", Rate: 99},
		},
	}
}

func initLucas() MONSTER {
	return MONSTER{
		NOM:      "Le L.",
		PVMax:    10000000000,
		PV:       1,
		PVMAXR:   10000000000,
		PVR:      1,
		Strength: 2000000000,
		Defense:  2,
		Spd:      1,
		Reiki:    2,
		Loot: []LootEntry{
			{Name: "BOTTE LUCACA", Rate: 99},
		},
	}
}

func initYannis() MONSTER {
	return MONSTER{
		NOM:      "Le RELOU",
		PVMax:    1000,
		PV:       1000,
		PVMAXR:   1000,
		PVR:      1000,
		Strength: 20,
		Defense:  10,
		Spd:      100000000,
		Reiki:    2,
		Loot: []LootEntry{
			{Name: "CASQUE YAYA", Rate: 99},
		},
	}
}

func IsMonsterDead(m *MONSTER) bool {
	return m == nil || m.PV <= 0 || m.PVR <= 0
}
