# 🎮 Projet RED — Jeu en ligne de commande (Go)

Petit RPG en ligne de commande conforme au sujet **Projet RED** : création de personnage, inventaire, marchand, forgeron, sorts, et combat tour par tour.

---

## ✅ Prérequis

- **Go 1.22+** (voir `go.mod`)
- Un terminal (PowerShell, cmd, bash…)

---

## 🕹️ Comment jouer (aperçu des menus)

 - Infos personnage : stats, or, XP, mana, équipements, inventaire.
 - Inventaire : utiliser potions / équiper objets / lire livre de sort.
 - Marchand : acheter potions, matériaux, livre de sort, augmentation d’inventaire.
 - Forgeron : fabriquer chapeau/tunique/bottes contre or + matériaux.
 - Entraînement : combat tour par tour contre un gobelin.
 - Qui sont-ils ? : clin d’œil à l’artiste caché (voir Mission 6).

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

Prix :

Potion de vie : 3 or

Potion de poison : 6 or

Livre de Sort : Boule de feu : 25 or

Fourrure de Loup : 4 or

Peau de Troll : 7 or

Cuir de Sanglier : 3 or

Plume de Corbeau : 1 or

Augmentation d’inventaire : 30 or

Vérification : or suffisant + place en inventaire.

Forgeron (craft & équipements)

Coût fixe : 5 or par craft, + matériaux requis.

Recettes :

Chapeau de l’aventurier : 1× Plume de Corbeau, 1× Cuir de Sanglier

Tunique de l’aventurier : 2× Fourrure de Loup, 1× Peau de Troll

Bottes de l’aventurier : 1× Fourrure de Loup, 1× Cuir de Sanglier

Équipement & bonus PV max :

Chapeau : +10 PV max

Tunique : +25 PV max

Bottes : +15 PV max

Remplacement : l’ancien équipement retourne dans l’inventaire.

Sorts & Mana

Coup de poing : 8 dégâts, 3 mana

Boule de feu : 18 dégâts, 7 mana

Blocage si mana insuffisant.

Combat tour par tour (Entraînement)

Adversaire : Gobelin d’entraînement (40 PV, 5 ATK).

Pattern : chaque tour 100% ATK ; **tous les 3 tours 200% AT