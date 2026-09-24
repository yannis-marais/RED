package ProjetRED

import (
	Forge "ProjetRED/Forge"
	personnage "ProjetRED/Personnage"
	"fmt"
	"sort"
)

func Forgeron(p *personnage.Character) {
	if p == nil {
		fmt.Println("Le forgeron ne peut pas traiter un personnage vide.")
		return
	}

	names := make([]string, 0, len(Forge.ForgeRecipes))
	for name := range Forge.ForgeRecipes {
		names = append(names, name)
	}
	sort.Strings(names)

	for {
		fmt.Println("\n=== FORGERON ===")
		for index, name := range names {
			fmt.Printf("%d. %s - ", index+1, name)
			printRecipeCost(*p, Forge.ForgeRecipes[name])
		}
		fmt.Println("0. Retour")

		choice, ok := ReadChoice("Votre choix : ")
		if !ok {
			fmt.Println("Choix invalide !")
			continue
		}
		if choice == 0 {
			return
		}
		if choice < 1 || choice > len(names) {
			fmt.Println("Choix invalide !")
			continue
		}

		Forge.Forge(p, names[choice-1])
		WaitForReturn()
	}
}

func printRecipeCost(p personnage.Character, recipe Forge.Recipe) {
	costNames := make([]string, 0, len(recipe.Cost))
	for name := range recipe.Cost {
		costNames = append(costNames, name)
	}
	sort.Strings(costNames)

	for index, name := range costNames {
		if index > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%s %d/%d", name, p.Inventory.Materials[name], recipe.Cost[name])
	}
	fmt.Println()
}
