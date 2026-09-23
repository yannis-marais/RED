package ProjetRED

import (
	personnage "ProjetRED/Personnage"
	"fmt"
)

type SkillBook struct {
	Name     string
	Skill    Skill
	MaxStack int
}

var SkillBooks = map[string]SkillBook{
	"Book of Fireball": {
		Name:     "Book of Fireball",
		Skill:    SkillList["Fireball"],
		MaxStack: 1,
	},
	"Book of Power Strike": {
		Name:     "Book of Power Strike",
		Skill:    SkillList["Power Strike"],
		MaxStack: 1,
	},
	"Book of Healing": {
		Name:     "Book of Healing",
		Skill:    SkillList["Healing"],
		MaxStack: 1,
	},
}

func GiveSkillBook(p *personnage.Character, book SkillBook) {
	if p == nil {
		return
	}

	qty := p.Inventory.SkillBooks[book.Name]
	if qty >= book.MaxStack {
		fmt.Println("Impossible :", p.Nom, "possède déjà", book.Name)
		return
	}

	p.Inventory.SkillBooks[book.Name] = qty + 1
	fmt.Println(book.Name, "donné à", p.Nom)
}

func LearnSkill(p *personnage.Character, book SkillBook) {
	if p == nil {
		return
	}

	qty := p.Inventory.SkillBooks[book.Name]
	if qty <= 0 {
		fmt.Println("Tu n'as pas le livre :", book.Name)
		return
	}

	if _, exists := p.Skills[book.Skill.Name]; exists {
		fmt.Println(p.Nom, "connaît déjà la compétence", book.Skill.Name)
		return
	}

	if p.Skills == nil {
		p.Skills = make(map[string]personnage.Skill)
	}

	p.Skills[book.Skill.Name] = personnage.Skill{
		Name:     book.Skill.Name,
		Damage:   book.Skill.BaseDamage,
		Heal:     book.Skill.Heal,
		Type:     book.Skill.Type,
		Strength: 0,
		Reiki:    0,
	}

	p.Inventory.SkillBooks[book.Name] = qty - 1
	fmt.Println(p.Nom, "a appris la compétence :", book.Skill.Name)
}
