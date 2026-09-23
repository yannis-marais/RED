package ProjetRED

import (
	"testing"

	personnage "ProjetRED/Personnage"
	enemies "ProjetRED/enemies"
)

func TestCombatTourParTourContreMonstre(t *testing.T) {
	p := personnage.CharacterCreation("Hero", personnage.Classes["Ronin"])
	p.PV = 100
	p.PVMax = 100
	p.Strength = 12
	p.Reiki = 18
	p.Cooldowns = map[string]int{}
	p.Skills = map[string]personnage.Skill{
		"Fireball": {Name: "Fireball", Damage: 30, Type: "Magic"},
		"Slash":    {Name: "Slash", Damage: 15, Type: "Nature"},
	}

	m := enemies.MONSTER{
		NOM:      "Gobelin",
		PVMax:    200,
		PV:       200,
		PVMAXR:   120,
		PVR:      120,
		Strength: 10,
		Spd:      10,
	}

	// Tour 1 : attaque physique puis skill magic avec bonus
	applyAttackDamage(&p, &m, p.Strength, "PV", 1.0)
	if m.PV != 188 {
		t.Fatalf("attaque physique attendue: PV=188, got %d", m.PV)
	}

	ok := applySkillWithMultiplier(&p, &m, "Fireball", 1.5)
	if !ok {
		t.Fatal("Fireball devrait être disponible")
	}
	if m.PVR >= 120 {
		t.Fatalf("Fireball devrait réduire la PVR, got PVR=%d", m.PVR)
	}
	if p.Cooldowns["Fireball"] == 0 {
		t.Fatal("Fireball devrait être en cooldown après usage")
	}

	// Monstre attaque en retour
	p.PV -= m.Strength
	if p.PV < 0 {
		p.PV = 0
	}
	if p.PV == 0 {
		t.Fatal("le joueur a été vaincu après le premier round")
	}

	// Tour 2 : attaque spirituelle puis skill physique avec bonus
	applyAttackDamage(&p, &m, p.Reiki, "PVR", 1.0)
	if m.PVR != 120-18 {
		t.Fatalf("attaque spirituelle attendue: PVR=102, got %d", m.PVR)
	}

	p.Cooldowns["Fireball"] = 0
	ok = applySkillWithMultiplier(&p, &m, "Slash", 1.5)
	if !ok {
		t.Fatal("Slash devrait être disponible")
	}
	if m.PV >= 188 {
		t.Fatalf("Slash avec bonus devrait réduire PV, got PV=%d", m.PV)
	}

	// Monstre attaque une seconde fois
	p.PV -= m.Strength
	if p.PV < 0 {
		p.PV = 0
	}
	if p.PV == 0 {
		t.Fatal("le joueur a été vaincu après le deuxième round")
	}

	// Vérification finale : combat cohérent, pas d’erreur logique
	if p.PV <= 0 || m.PV <= 0 && m.PVR <= 0 {
		t.Fatal("combat incohérent : état de fin invalide")
	}
}
