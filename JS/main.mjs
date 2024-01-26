import random from 'random'

//initialisation du jeu

// 14 A, 4 B, 7 C, 5 D, 19 E, 2 F, 4 G, 2 H, 11 I, 1 J, 1 K, 6 L, 5 M, 9 N, 8 O, 4 P, 1 Q, 10 R, 7 S, 9 T, 8 U, 2 V, 1 W, 1 X, 1 Y et 2 Z.
let sac =  [["A", 14],["B", 4],["C", 7],["D", 5],["E", 19],["F", 2],["G", 4],["H", 2],["I", 11],["J", 1],["K", 1],["L", 6],["M", 5],["N" , 9],["O" , 8],["P" , 4],["Q" , 1],["R" , 10],["S", 7],["T", 9],["U", 8],["V", 2],["W" , 1],["X" , 1],["Y" , 1],["Z", 2]]

let tapis1 = [["B","O","N"],["J","E","U","X"]]
let tapis2 = [["H","E","L","L","O"]]

let main1 = []
let main2 = []

//pas sûre que ce soit utile
let alphabet = ["A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"]


//fonctions du jeu
const piocher_x_lettres = function(x, sac) {
    let i = 0
    let new_lettres = []
    while (i < x) {
        let number = random.int(0,25)
        let lettre = sac[number][0]
        sac[number][1] = (sac[number][1])-1
        new_lettres.push(lettre)
        i = i + 1
    }
    return new_lettres
}

const affiche_tapis = function(tapis) {
    let l = 0
    while (l < tapis.length) {
        console.log(tapis[l])
        l += 1
    }
}

//splice(start, deleteCount, item1, item2, /* …, */ itemN)

const poser_lettre = function(lettre, tapis, ligne, place_dans_la_ligne) {
    tapis[ligne].splice(place_dans_la_ligne, 0, lettre)
} 


//début du jeu

let lettres_piochees = piocher_x_lettres(6, sac)
console.log(lettres_piochees)

console.log("tapis joueur 1 :")
affiche_tapis(tapis1)
console.log("tapis joueur 2 :")
affiche_tapis(tapis2)




