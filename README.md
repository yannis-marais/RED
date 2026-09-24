# Projet-RED_Kurogane
 
## Présentation du jeu
 
Vous aimez l'aventure, le RPG et les combats stratégiques au tour par tour ? Alors bienvenue à **Kurogane** ! Ce pays où 2 princes se battent pour avoir le trône a besoin d'un aventurier pour les départager... Il faudra donc aller de ville en ville tout en s'occupant des montres qui y rodent pour pouvoir aider l'un ou l'autre à prendre le pouvoir
 
Au fil de votre périple, vous devrez choisir votre voie (guerrier au sabre, combattant cuirassé ou mage spirituel), gérer votre inventaire, forger votre propre équipement à partir des ressources récoltées, affronter des créatures de plus en plus redoutables et faire les bons choix pour survivre. Attention toutefois : certains monstres ne se laissent pas abattre facilement, et un sort particulièrement chaotique, le *Make a Wish*, pourrait bien retourner la situation... dans un sens comme dans l'autre.
 
## Installation du jeu
 
Il s'agit d'un jeu qui se déroule entièrement dans le terminal.
 
Pour jouer :
1. Installer le langage **Go** sur votre ordinateur
2. Installer un éditeur de code (ex : Visual Studio Code)
3. Cloner ou télécharger ce dépôt Git
4. Ouvrir le dossier du projet dans votre éditeur
5. Dans le terminal, se placer à la racine du projet (là où se trouve `main.go`)
6. Lancer la commande :
```
go run .
```
 
## Commandes
 
Le jeu fonctionne entièrement par menus affichés dans le terminal. Le joueur sélectionne une option en entrant le **numéro** correspondant, puis valide avec Entrée. Certaines actions (comme les attaques en combat) déclenchent un **QTE** (Quick Time Event) : il faut appuyer sur Entrée au bon moment pour infliger un coup critique.
 
## Fonctionnalités
 
- **Création de personnage** : choix du nom et de la classe (Ronin, Cuirassé, Mage spirituel), chaque classe ayant ses propres statistiques et son équipement de départ
- **Exploration** : déplacement entre les villes, chacune menant à son propre donjon peuplé de monstres différents
- **Inventaire** : objets, consommables, matériaux et livres de sorts, avec une capacité limitée
- **Équipement** : casque, armure, bottes et arme, chacun modifiant les statistiques du personnage
- **Marchand** : achat d'objets, de potions et de livres de sorts contre de l'argent
- **Forgeron** : fabrication d'équipements à partir de matériaux récoltés sur les monstres vaincus
- **Combat au tour par tour** :
  - Ordre des tours déterminé par la statistique de Vitesse
  - Attaque physique et attaque spirituelle (deux barres de vie distinctes à vaincre : PV et PVR)
  - Compétences (Skills) avec temps de recharge (cooldown)
  - Sort chaotique **Make a Wish** : soin, dégâts, effet de statut ou k.o. instantané selon le tirage
  - Effets de statut à durée (poison, régénération...)
  - Butin (loot) récupéré à la fin de chaque combat
- **Progression** : gestion des points de vie, de l'expérience, du niveau et de l'argent du personnage
- **Sauvegarde / Chargement** : possibilité de sauvegarder sa partie et de la reprendre plus tard
## Équipe
 
- Antoine
- Yannis
- Brian
- Lucas