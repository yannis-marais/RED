package ProjetRED

import (
	"testing"

	personnage "ProjetRED/Personnage"
	enemies "ProjetRED/enemies"
)

func TestCombatFlowHandlesPhysicalAndSpiritDamageAndSkills(t *testing.T) {
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

	p.PV = 80
	p.PVMax = 80
	p.Reiki = 20
	m.PV = 200
	m.PVR = 100

	degatsPhysique := p.Strength
	m.PV -= degatsPhysique
	if m.PV != 190 {
		t.Fatalf("physical attack should reduce PV, got PV=%d", m.PV)
	}

	degatsSpirit := p.Reiki
	m.PVR -= degatsSpirit
	if m.PVR != 80 {
		t.Fatalf("spirit attack should reduce PVR, got PVR=%d", m.PVR)
	}

	ok := applySkill(&p, &m, "Fireball")
	if !ok {
		t.Fatal("Fireball should be usable")
	}
	if m.PVR >= 80 {
		t.Fatalf("Fireball should damage PVR after skill use, got PVR=%d", m.PVR)
	}

	ok = applySkill(&p, &m, "Slash")
	if !ok {
		t.Fatal("Slash should be usable")
	}
	if m.PV >= 190 {
		t.Fatalf("Slash should damage PV after skill use, got PV=%d", m.PV)
	}
}
