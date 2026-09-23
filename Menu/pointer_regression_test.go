package projetRED

import (
	"testing"

	personnage "ProjetRED/Personnage"
)

func TestForgeronAcceptsPointerToCharacter(t *testing.T) {
	p := personnage.CharacterCreation("Hero", personnage.Classes["Ronin"])
	Forgeron(&p)
}
