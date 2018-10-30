# Vitrin website for cloudltd

  **(Version Française)**

# Démarrage
Si nécessaire, installer git pour récupérer le projet: [windows](https://gitforwindows.org/)/[autre](https://git-scm.com/book/fr/v1/D%C3%A9marrage-rapide-Installation-de-Git)

## Installer les prérequis



*  [node js server](https://nodejs.org/en/)
   - Un environnement serveur permettant de récupérer et de déployer des application en javascript.
* [npm](https://www.npmjs.com/get-npm)
   - Normalement récupéré lors de l'installation de nodeJs, npm permettra de récupérer les codes extérieurs servant à l'application.
* [angular cli tools](https://cli.angular.io/)
   - Un outil en ligne de commande pour facilement construire et déployer le site.

Au lancement, le comportement par défaut de l'environnement sera de servir l'application sur http://localhost:4200 .



## Récupérer l'application



Dans un dossier et avec un outil de ligne de commande (bash, power shell) entrez la commande suivante:
```
git clone https://github.com/aurelienCastellarnau/cloudltd.git
```
Cela va créer un dossier **cloudltd/** contenant le site vitrine.

## Installer le site

Dans le dossier **clouldt/** et avec un outil de ligne de commande entrez la commande suivante pour récupérer les éléments de code extérieurs *(liste dans le fichier package.json, section dependencies)*:
```
npm install
```
Cela peut prendre quelques minutes.
Une fois l'installation des dépendances terminées, entrez la commande suivante pour lancer le server d'angular CLI:
```
ng serve --open
```
L'environnement attend une connexion d'un navigateur sur http://localhost:4200/ pour compiler et lancer le site.


# Dépendances



bootstrap@4.1.3

font-awesome@4.7.0

ng-simple-slideshow ^1.2.5
