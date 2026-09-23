package ProjetRED

import (
	"testing"

	personnage "ProjetRED/Personnage"
	enemies "ProjetRED/enemies"
)

func TestSkillChoiceDamagesEnemySpiritBar(t *testing.T) {
	p := personnage.CharacterCreation("Hero", personnage.Classes["Ronin"])
	p.Strength = 10
	p.Reiki = 20
	p.Skills = map[string]personnage.Skill{
		"Fireball": {
			Name:   "Fireball",
			Damage: 30,
			Type:   "Magic",
		},
		"Slash": {
			Name:   "Slash",
			Damage: 15,
			Type:   "Nature",
		},
	}

	m := enemies.MONSTER{
		NOM:    "Gobelin",
		PVMax:  200,
		PV:     200,
		PVMAXR: 100,
		PVR:    100,
	}

	ok := applySkill(&p, &m, "Fireball")
	if !ok {
		t.Fatal("Fireball should be usable")
	}

	if m.PVR != 100-30-20 { // 30 base + 20 reiki
		t.Fatalf("Fireball should damage PVR, got PVR=%d", m.PVR)
	}

	if m.PV != 200 {
		t.Fatalf("Fireball should not damage PV, got PV=%d", m.PV)
	}

	ok = applySkill(&p, &m, "Slash")
	if !ok {
		t.Fatal("Slash should be usable")
	}

	if m.PV != 200-15-10 {
		t.Fatalf("Slash should damage PV, got PV=%d", m.PV)
	}
}
