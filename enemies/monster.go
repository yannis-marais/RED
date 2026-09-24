package ProjetRED

import "math/rand"

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
			{Name: "Fer", Rate: 35},
			{Name: "Bois", Rate: 25},
			{Name: "Cuir", Rate: 15},
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
			{Name: "Fer", Rate: 45},
			{Name: "Cuir", Rate: 10},
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
			{Name: "Bois", Rate: 40},
			{Name: "Cuir", Rate: 35},
			{Name: "Fer", Rate: 20},
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
			{Name: "Cuir", Rate: 45},
			{Name: "Cristal", Rate: 25},
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
			{Name: "Cuir", Rate: 50},
			{Name: "Fer", Rate: 15},
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
			{Name: "Cuir", Rate: 30},
			{Name: "Fer", Rate: 25},
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
			{Name: "Fer", Rate: 55},
			{Name: "Cuir", Rate: 30},
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
			{Name: "Cristal", Rate: 60},
			{Name: "Diamant", Rate: 25},
			{Name: "Fer", Rate: 40},
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

func RollLoot(m *MONSTER) []string {
	if m == nil {
		return nil
	}

	loot := make([]string, 0, len(m.Loot))
	for _, entry := range m.Loot {
		if entry.Name != "" && entry.Rate > 0 && rand.Intn(100) < entry.Rate {
			loot = append(loot, entry.Name)
		}
	}
	return loot
}

func IsMonsterDead(m *MONSTER) bool {
	return m == nil || m.PV <= 0 || m.PVR <= 0
}
