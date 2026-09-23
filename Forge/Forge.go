package ProjetRED

import (
	Equipement "ProjetRED/Equipement"
	Personnage "ProjetRED/Personnage"
	"fmt"
	"reflect"
)

type Recipe struct {
	Result Equipement.Item
	Cost   map[string]int
}

var ForgeRecipes = map[string]Recipe{
	"La Maxime": {
		Result: Equipement.Items["La Maxime"],
		Cost: map[string]int{
			"Diamant": 10,
			"Cuir":    5,
			"Fer":     10,
		},
	},
	"Maximilian Boots": {
		Result: Equipement.Items["Maximilian Boots"],
		Cost: map[string]int{
			"Iron":    10,
			"Diamant": 5,
			"Leather": 20,
		},
	},
	"Maximilian Armor": {
		Result: Equipement.Items["Maximilian Armor"],
		Cost: map[string]int{
			"Iron":    20,
			"Crystal": 15,
			"Diamant": 5,
		},
	},
	"Maximilian Helmet": {
		Result: Equipement.Items["Maximilian Helmet"],
		Cost: map[string]int{
			"Diamant": 5,
			"Cuir":    5,
			"Fer":     10,
		},
	},
}

func Forge(p Personnage.Character, itemName string) {
	recipe, ok := ForgeRecipes[itemName]
	if !ok {
		fmt.Println("Recette inconnue :", itemName)
		return
	}

	// Vérifier les matériaux
	for mat, needed := range recipe.Cost {
		if inventoryMaterial(p.Inventory, mat) < needed {
			fmt.Println("Matériaux insuffisants pour forger", itemName)
			fmt.Println("Il manque :", mat)
			return
		}
	}

	// Retirer les matériaux
	for mat, needed := range recipe.Cost {
		removeInventoryMaterial(p.Inventory, mat, needed)
	}

	// Donner l'objet forgé
	Equipement.AddItem(p, recipe.Result)

	fmt.Println(p.Nom, "a forgé :", itemName)
}

func inventoryMaterial(inventory interface{}, name string) int {
	materials := inventoryMaterials(inventory)
	if !materials.IsValid() {
		return 0
	}
	value := materials.MapIndex(reflect.ValueOf(name).Convert(materials.Type().Key()))
	if !value.IsValid() {
		return 0
	}
	return int(value.Int())
}

func removeInventoryMaterial(inventory interface{}, name string, amount int) {
	materials := inventoryMaterials(inventory)
	if !materials.IsValid() || materials.IsNil() {
		return
	}
	key := reflect.ValueOf(name).Convert(materials.Type().Key())
	value := reflect.ValueOf(inventoryMaterial(inventory, name) - amount).Convert(materials.Type().Elem())
	materials.SetMapIndex(key, value)
}

func inventoryMaterials(inventory interface{}) reflect.Value {
	value := reflect.ValueOf(inventory)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return reflect.Value{}
	}
	for _, fieldName := range []string{"Materials", "Ressources", "Resources", "Items"} {
		field := value.FieldByName(fieldName)
		if field.IsValid() && field.Kind() == reflect.Map &&
			field.Type().Key().Kind() == reflect.String && field.Type().Elem().Kind() == reflect.Int {
			return field
		}
	}
	return reflect.Value{}
}
