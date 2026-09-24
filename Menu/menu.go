package ProjetRED

import (
	City "ProjetRED/deplacementville"
	Equipement "ProjetRED/Equipement"
	Marchand "ProjetRED/Marchand"
	personnage "ProjetRED/Personnage"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var CurrentPlayer *personnage.Character

func SetCurrentPlayer(p *personnage.Character) {
	CurrentPlayer = p
}

func SaveCurrentGame() error {
	if CurrentPlayer == nil {
		return fmt.Errorf("aucun personnage actif")
	}
	return personnage.SaveCharacterToFile(personnage.DefaultSavePath(), *CurrentPlayer)
}

func LoadCurrentGame() (*personnage.Character, error) {
	player, err := personnage.LoadCharacterFromFile(personnage.DefaultSavePath())
	if err != nil {
		return nil, err
	}
	SetCurrentPlayer(&player)
	return &player, nil
}

// menu du lancement
func StartMenu() {

	fmt.Println("\n=== MENU CREATION ===")
	fmt.Println("1. Crée un nouveau Personage")
	fmt.Println("2. Charger une sauvegarde")
	fmt.Println("0. Quitter")
	choice, reponse := ReadChoice("Votre choix : ")
	if !reponse {
		fmt.Println("Choix invalide !")
		StartMenu()
	}
	switch choice {
	case 1:
		Player := personnage.CharacterCreation(CreerPerso())
		SetCurrentPlayer(&Player)
		MainMenu(&Player)
	case 2:
		player, err := LoadCurrentGame()
		if err != nil {
			fmt.Println("Aucune sauvegarde trouvée.")
			StartMenu()
			return
		}
		MainMenu(player)
	case 0:
		fmt.Println("Au Revoir !")
		return
	default:
		fmt.Println("Choix invalide !")

	}
}

// menu principal qui permet de faire pivot
func MainMenu(p *personnage.Character) {
	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1. Lancé la partie")
		fmt.Println("2. Afficher les informations du personnage")
		fmt.Println("3. Accéder à l'inventaire")
		fmt.Println("4. Marchand")
		fmt.Println("5. Forgeron")
		fmt.Println("6. Sauvegarder")
		fmt.Println("7. Charger la sauvegarde")
		fmt.Println("0. Quitter")

		choice, reponse := ReadChoice("Votre choix : ")

		if !reponse {
			fmt.Println("Choix invalide !")
			continue
		}

		switch choice {
		case 1:
			City.Ville1(p)
		case 2:
			DisplayInfo(*p)
			WaitForReturn()
		case 3:
			ManageInventory(p)
			WaitForReturn()
		case 4:
			Marchand.MarchandForPlayer(p, func() {})
		case 5:
			Forgeron(p)
		case 6:
			Marchand.Marchand(p, func() {})
		case 5:
			Forgeron(p)
		case 6:
			if err := SaveCurrentGame(); err != nil {
				fmt.Println("Erreur de sauvegarde :", err)
			} else {
				fmt.Println("Partie sauvegardée dans save.json")
			}
			WaitForReturn()
		case 7:
			player, err := LoadCurrentGame()
			if err != nil {
				fmt.Println("Erreur de chargement :", err)
			} else {
				fmt.Println("Partie chargée depuis save.json")
				CurrentPlayer = player
			}
			WaitForReturn()
		case 0:
			fmt.Println("À bientôt !")
			return
		default:
			fmt.Println("Choix invalide !")
		}
	}
}

// lis la reponse donné depuis le terminal
func ReadChoice(prompt string) (int, bool) {
	fmt.Print(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return 0, false
	}

	value, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return 0, false
	}

	return value, true
}

// attend que le joeure apuis sur entré
func WaitForReturn() {
	fmt.Println("Appuyez sur Entrée pour continuer...")
	_, _ = fmt.Scanln()
}

// affiche les info du joeur
func DisplayInfo(p personnage.Character) string {
	var sb strings.Builder

	const largeur = 40
	ligne := strings.Repeat("─", largeur)

	fmt.Fprintf(&sb, "╭%s╮\n", ligne)
	fmt.Fprintf(&sb, "│ %-38s │\n", fmt.Sprintf("%s — %s (Lvl %d)", p.Nom, p.Classe.Nom, p.LVL))
	fmt.Fprintf(&sb, "├%s┤\n", ligne)

	// Barre de PV
	fmt.Fprintf(&sb, "│ %-38s │\n", fmt.Sprintf("PV: %d/%d %s", p.PV, p.PVMax, barreDeVie(p.PV, p.PVMax, 15)))
	fmt.Fprintf(&sb, "│ %-38s │\n", fmt.Sprintf("XP: %.0f", p.XP))

	fmt.Fprintf(&sb, "├%s┤\n", ligne)
	fmt.Fprintf(&sb, "│ %-38s │\n", "Statistiques")
	fmt.Fprintf(&sb, "│   %-15s %-20d │\n", "Force:", p.Strength)
	fmt.Fprintf(&sb, "│   %-15s %-20d │\n", "Défense:", p.Defense)
	fmt.Fprintf(&sb, "│   %-15s %-20d │\n", "Reiki:", p.Reiki)
	fmt.Fprintf(&sb, "│   %-15s %-20d │\n", "Vitesse:", p.Spd)

	fmt.Fprintf(&sb, "├%s┤\n", ligne)
	fmt.Fprintf(&sb, "│ %-38s │\n", "Équipement")
	fmt.Fprintf(&sb, "│   %-15s %-20v │\n", "Arme:", p.Weapon)
	fmt.Fprintf(&sb, "│   %-15s %-20v │\n", "Casque:", p.Helmet)
	fmt.Fprintf(&sb, "│   %-15s %-20v │\n", "Armure:", p.Armor)
	fmt.Fprintf(&sb, "│   %-15s %-20v │\n", "Bottes:", p.Boots)

	fmt.Fprintf(&sb, "╰%s╯\n", ligne)

	result := sb.String()
	fmt.Print(result)
	return result
}

// barreDeVie construit une petite barre style [████████░░]
func barreDeVie(pv, pvMax, taille int) string {
	if pvMax <= 0 {
		return ""
	}
	rempli := int(float64(pv) / float64(pvMax) * float64(taille))
	if rempli > taille {
		rempli = taille
	}
	if rempli < 0 {
		rempli = 0
	}
	return "[" + strings.Repeat("█", rempli) + strings.Repeat("░", taille-rempli) + "]"
}

// affiche l'inventaire du joueur
func AccessInventory(p personnage.Character) string {
	var sb strings.Builder

	const largeur = 40
	ligne := strings.Repeat("─", largeur)

	fmt.Fprintf(&sb, "╭%s╮\n", ligne)
	fmt.Fprintf(&sb, "│ %-38s │\n", "Inventaire")

	fmt.Fprintf(&sb, "├%s┤\n", ligne)
	fmt.Fprintf(&sb, "│ %-38s │\n", "Argent dans la poche")
	fmt.Fprintf(&sb, "│ %-38s │\n", fmt.Sprintf("%d yen", p.Purse))
	fmt.Fprintf(&sb, "│ %-38s │\n", fmt.Sprintf("Capacité : %d/%d", p.Inventory.UsedSlots(), p.Inventory.Capacity))

	fmt.Fprintf(&sb, "├%s┤\n", ligne)
	fmt.Fprintf(&sb, "│ %-38s │\n", "Objets")
	ecrireSection(&sb, p.Inventory.Items)

	fmt.Fprintf(&sb, "├%s┤\n", ligne)
	fmt.Fprintf(&sb, "│ %-38s │\n", "Consommables")
	ecrireSection(&sb, p.Inventory.Consumables)

	fmt.Fprintf(&sb, "╰%s╯\n", ligne)
	return sb.String()
}

// ecrireSection affiche une map triée par clé, avec un message si elle est vide.
func ecrireSection(sb *strings.Builder, items map[string]int) {
	if len(items) == 0 {
		fmt.Fprintf(sb, "│   %-36s │\n", "Aucun")
		return
	}

	keys := make([]string, 0, len(items))
	for name := range items {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	for _, name := range keys {
		ligne := fmt.Sprintf("%s x%d", name, items[name])
		fmt.Fprintf(sb, "│   %-36s │\n", ligne)
	}
}

// fonction renvoie les variables qui seront données au character creator
func CreerPerso() (string, personnage.Classe) {
	fmt.Print("Quel est votre nom ? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nom := strings.TrimSpace(scanner.Text())

	fmt.Println("Choisissez le numero d'une classe :")
	fmt.Println("1. Ronin")
	fmt.Println("2. Cuirassé")
	fmt.Println("3. Mage spirituel")

	choice, ok := ReadChoice("Votre choix : ")
	if !ok {
		fmt.Println("Choix invalide, classe par défaut : Ronin")
		return nom, personnage.Classes["Ronin"]
	}

	switch choice {
	case 1:
		return nom, personnage.Classes["Ronin"]
	case 2:
		return nom, personnage.Classes["Cuirassé"]
	case 3:
		return nom, personnage.Classes["Mage spirituel"]
	default:
		fmt.Println("Choix invalide, classe par défaut : Ronin")
		return nom, personnage.Classes["Ronin"]
	}
}

func ManageInventory(p *personnage.Character) {
	AccessInventory(*p)

	fmt.Println("\n=== Inventaire ===")
	fmt.Println("1. Interagire avec les Objets ")    // armure etc
	fmt.Println("2. Interagire avec les Consomable") // potion
	fmt.Println("3. Interagire avec les Livre de Sort ")
	fmt.Println("0. Quitter")

	choice, reponse := ReadChoice("Votre choix : ")

	if !reponse {
		fmt.Println("Choix invalide !")
		ManageInventory(p)
		return
	}

	switch choice {
	case 1:
		SelectFromList(p, DisplayItem(*p), "item")
		WaitForReturn()
	case 2:
		SelectFromList(p, DisplayConsumables(*p), "consumable")
		WaitForReturn()
	case 3:
		SelectFromList(p, DisplaySkillBooks(*p), "skillbook")
		WaitForReturn()
	case 0:
		return
	default:
		fmt.Println("Choix invalide !")
		ManageInventory(p)
	}
}

// Crée une nouvelle func qui permet de lister tout ce qu'il y a dans la partie Inventory.Items
func DisplayItem(p personnage.Character) []string {
	if len(p.Inventory.Items) == 0 {
		fmt.Println("Aucun objet dans l'inventaire.")
		return nil
	}

	names := make([]string, 0, len(p.Inventory.Items))
	for name := range p.Inventory.Items {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Println("=== Items ===")
	for i, name := range names {
		fmt.Printf("%d. %s (x%d)\n", i+1, name, p.Inventory.Items[name])
	}

	return names
}

// Crée une nouvelle func qui permet de lister tout ce qu'il y a dans la partie Inventory.Consumables
func DisplayConsumables(p personnage.Character) []string {
	if len(p.Inventory.Consumables) == 0 {
		fmt.Println("Aucun consommable dans l'inventaire.")
		return nil
	}

	names := make([]string, 0, len(p.Inventory.Consumables))
	for name := range p.Inventory.Consumables {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Println("=== Consommables ===")
	for i, name := range names {
		fmt.Printf("%d. %s (x%d)\n", i+1, name, p.Inventory.Consumables[name])
	}

	return names
}

// Crée une nouvelle func qui permet de lister tout ce qu'il y a dans la partie Inventory.SkillBooks
func DisplaySkillBooks(p personnage.Character) []string {
	if len(p.Inventory.SkillBooks) == 0 {
		fmt.Println("Aucun livre de compétence dans l'inventaire.")
		return nil
	}

	names := make([]string, 0, len(p.Inventory.SkillBooks))
	for name := range p.Inventory.SkillBooks {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Println("=== Livres de compétences ===")
	for i, name := range names {
		fmt.Printf("%d. %s (x%d)\n", i+1, name, p.Inventory.SkillBooks[name])
	}

	return names
}

// Fonction générique de sélection, compatible avec les 3 Display ci-dessus
// On lui passe la liste (retournée par un Display) et elle gère le choix
func SelectFromList(p *personnage.Character, names []string, category string) string {
	if names == nil {
		return ""
	}

	fmt.Println("0. Ne rien faire")
	fmt.Print("Choisis le Numero de l'objet que tu souhaite Utilisé/Equipé : ")
	var choice int
	fmt.Scan(&choice)

	if choice == 0 {
		fmt.Println("Aucune action effectuée.")
		SelectFromList(p, names, category)
		return ""
	}

	if choice < 1 || choice > len(names) {
		fmt.Println("Choix invalide.")
		SelectFromList(p, names, category)
		return ""
	}

	selected := names[choice-1]

	switch category {
	case "item":
		Equipement.EquipItem(p, selected)
		fmt.Println("Tu as sélectionné un objet :", selected)

	case "consumable":
		fmt.Println("Tu as sélectionné un consommable :", selected)
		if consumable, ok := Equipement.GetConsumableByName(selected); ok {
			Equipement.UseConsumable(p, consumable)
		} else {
			fmt.Println("Consommable introuvable.")
		}
	case "skillbook":
		fmt.Println("Tu as sélectionné un livre de sort :", selected)

	}

	return selected
}
