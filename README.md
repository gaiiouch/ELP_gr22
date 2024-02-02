# ELP_gr22
Projet 3TC ELP

Membres du groupe : Anne-Gaëlle Mauger 3TC2, Louise Marc 3TC2, Ilhême Maliki 3TC3

## Projet GO

Nous avons choisi de traiter le problème du produit de deux matrices carrées. 

Nous avons commencé par écrire un code qui calculait le produit des matrices données en paramètre de façon séquentielle. 
Pour accélérer le processus de calcul, nous avons ajouté les go routines de façon à ce que chaque go routine calcule une ligne de la matrice résultat. 
Les lignes résultantes sont échangées grâce à un channel et assemblées en une seule matrice Res de manière séquentielle pour éviter les problèmes de mémoire partagée. 
En troisième étape, nous avons ajouté un système de client qui envoie les matrices générées aléatoirement à un serveur. Le serveur calcule la matrice produit et la renvoie au client, qui écrit dans des fichiers textes les trois matrices.

Au final, il y a deux versions du code, les légères différences n'influant pas sur le fonctionnement global des go routines et du système client-serveur TCP. 
- Dans la première version, il faut changer la valeur de la constante "taille" des matrices voulues à la main dans les fichiers tcpclient.go et tcpserver.go.
- Dans la deuxième version, l'utilisateur peut décider de la taille des matrices lors de l'exécution de tcpclient.go grâce à un système de slices.

Pour lancer les fichiers, les commandes à écrire dans les terminaux sont notées en commentaire en haut des fichiers tcpclient.go et tcpserver.go. 
Il faut lancer le serveur avant le client.

## Projet ELM

Pour télécharger ELM en fonction de son OS : https://guide.elm-lang.org/install/elm.html

Packages à télécharger pour le bon fonctionnement du code avec elm install :
- elm/http
- elm/json
- elm/random
- elm-community/list-extra

Le but de ce projet était de créer une page web qui affiche une définition d'un mot choisi aléatoirement dans une liste. L'utilisateur doit deviner ce mot grâce à la définition.
En fonction de la réponse donnée par l'utilisateur dans la zone de réponse, la page affiche si le mot est correct ou non. En cas d'échec, l'utilisateur peut choisir d'afficher la réponse ou de changer la définition pour jouer avec un autre mot.

## Projet JavaScript

Pour lancer le code, il faut commencer par installer node JS, ainsi que des packages grâce aux commandes suivantes  :
- npm install random
- npm install inquirer
- npm install fs

Le but de ce projet était d'implémenter le jeu de société Jarnac. Notre version se lance et se joue directement dans le terminal avec la commande *node main.js* si vous êtes déjà dans le dossier JS.
!! En fonction du dossier dans lequel vous avez ouvert le programme, le chemin pour lire le fichier liste_francais.txt qui contient la liste des mots autorisés peut être faux. Si une erreur apparaît lors du lancement, modifiez le chemin ligne 129 du fichier main.js en fonction de votre répertoire.
Pour une partie du code, nous nous sommes inspirées du code suivant : https://github.com/SBoudrias/Inquirer.js/blob/master/packages/inquirer/examples/pizza.js .
