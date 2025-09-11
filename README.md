# 🎮 Projet RED — Jeu en ligne de commande (Go)

Petit RPG en ligne de commande conforme au sujet **Projet RED** : création de personnage, inventaire, marchand, forgeron, sorts, et combat tour par tour.

---

## ✅ Prérequis

- **Go 1.22+** (voir `go.mod`)
- Un terminal (PowerShell, cmd, bash…)

---

## 🕹️ Comment jouer (aperçu des menus)

 - **Infos personnage :** stats, or, XP, mana, équipements, inventaire.
 - **Inventaire :** utiliser potions / équiper objets / lire livre de sort.
 - **Marchand :** acheter potions, matériaux, livre de sort, augmentation d’inventaire.
 - **Forgeron :** fabriquer chapeau/tunique/bottes contre or + matériaux.
 - **Entraînement :** combat tour par tour contre un gobelin.
 - **Qui sont-ils ? :** clin d’œil à l’artiste caché (voir Mission 6).

---

## 📂 Structure du projet

```bash
projet-red_MONJEU/
├─ README.md                # Ce fichier
├─ go.mod                   # Module Go (1.22)
├─ docs/
│  └─ gestion_projet.md     # Notes de gestion de projet
└─ src/
   ├─ main.go               # Point d'entrée
   ├─ game.go               # Menu principal
   ├─ character.go          # Création + affichage personnage
   ├─ inventory.go          # Inventaire + objets consommables/équipements
   ├─ merchant.go           # Marchand (achats, or)
   ├─ forge.go              # Craft + équipements + bonus PV
   ├─ spells.go             # Sorts + mana
   ├─ combat.go             # Combat tour par tour (gobelin)
   ├─ types.go              # Structures et constantes
   └─ utils.go              # Utilitaires (I/O, clamp, formatage)
```

---

## ✨ Fonctionnalités
**Personnage**

 - Création interactive (nom nettoyé et capitalisé ; classe Humain / Elfe / Nain).
 - PV max selon la classe (100 / 80 / 120), PV initiaux = 50% PV max.
 - Compétence de base : Coup de poing (appris d’office).
 - Or de départ : 100.

**Inventaire**

 - Capacité 10 objets (contrôle à l’ajout).
 - Augmentation d’inventaire : +10 emplacements (max 3 fois).
 - Utilisation d’objets depuis l’inventaire :
     - Potion de vie : +50 PV, plafonné à PV max.
     - Potion de poison : -10 PV/s pendant 3 s (démo d’effet sur le joueur).
     - Livre de Sort : Boule de feu : apprend le sort (une seule fois).
     - Équipements : s’équipent depuis l’inventaire (voir bonus ci-dessous).

**Marchand (avec débits d’or)**

 - Prix :
     - Potion de vie : 3 or
     - Potion de poison : 6 or
     - Livre de Sort : Boule de feu : 25 or
     - Fourrure de Loup : 4 or
     - Peau de Troll : 7 or
     - Cuir de Sanglier : 3 or
     - Plume de Corbeau : 1 or
     - Augmentation d’inventaire : 30 or
 - Vérification : or suffisant + place en inventaire.

**Forgeron (craft & équipements)**

 - Coût fixe : 5 or par craft, + matériaux requis.
 - Recettes :
     - Chapeau de l’aventurier : 1× Plume de Corbeau, 1× Cuir de Sanglier
     - Tunique de l’aventurier : 2× Fourrure de Loup, 1× Peau de Troll
     - Bottes de l’aventurier : 1× Fourrure de Loup, 1× Cuir de Sanglier
 - Équipement & bonus PV max :
     - Chapeau : +10 PV max
     - Tunique : +25 PV max
     - Bottes : +15 PV max
 - Remplacement : l’ancien équipement retourne dans l’inventaire.

**Sorts & Mana**

 - Coup de poing : 8 dégâts, 3 mana
 - Boule de feu : 18 dégâts, 7 mana
 - Blocage si mana insuffisant.

**Combat tour par tour (Entraînement)**

 - Adversaire : Gobelin d’entraînement (40 PV, 5 ATK).
 - Pattern : 
     - Chaque tour → 100% de l’attaque (5 dmg).
     - Tous les 3 tours → 200% de l’attaque (10 dmg).
 - Initiative : le premier à jouer est celui qui a la plus haute valeur d’initiative.
 - Actions disponibles :
     - Attaquer (5 dégâts).
     - Inventaire (utiliser un objet).
     - Sorts (si appris).
     - Fuir (mettre fin au combat).
     - Résurrection auto si mort : retour à 50% PV max.

**Progression (XP & niveaux)**

 - Victoire → gain d’XP (10 par gobelin).
 - Montée de niveau quand Exp >= ExpMax.
 - À chaque niveau :
     - +5 PV max
     - +2 Mana max
     - +1 Initiative
     - Soins/mana restaurés
 - `ExpMax` augmente de +10 à chaque niveau, excès conservé.

**Autres**

 - Navigation claire dans le menu (retours possibles partout).
 - Menu “Qui sont-ils ?” → artiste caché : ABBA (Mission 6).
 - Mission 5 déjà enrichie : équipements utilisables, montée en niveau, inventaire amélioré.

 ---
## 📜 Couverture des tâches

 - **Tâches 01 → 22 :** toutes implémentées.
 - **Missions 1 → 4 :** intégrées (initiative, XP/niveau, sorts, mana).
 - **Mission 5 :** enrichissements apportés (système RPG plus complet).
 - **Mission 6 :** menu “Qui sont-ils ?” → réponse : ABBA et ...

 ---

 ## 👤 Auteurs

 - Jack ...
 - Eliott ...
 - Florin HAMELIN