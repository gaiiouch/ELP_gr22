import random from 'random'
import { jarnac } from './jarnac.js'
import { affiche_main, affiche_tapis } from './affichage.js'
import { affiche_fin } from './fin.js'
import { jouer_tour } from './tour.js'

//FONCTIONS DU JEU
const piocher_x_lettres = function(x, sac) {
    let i = 0
    let new_lettres = []
    while (i < x) {
        let number = random.int(0,sac.length-1)
        let lettre = sac[number][0]
        sac[number][1] = (sac[number][1])-1
        if (sac[number][1] === 0){
            sac.splice(number,1)
        }
        new_lettres.push(lettre)
        i = i + 1
    }
    return new_lettres
}

const playGame = async (main1, main2, tapis1,tapis2) => {
    let end = false;
    let tour = 0;
    let tapis;
    let main;
    let num;

    while (!end) {

        console.log("\n << Tour du joueur " + (tour % 2 + 1)+" >>\n")

        console.log("--------------- PLATEAU ---------------")
        affiche_main(main1, 1)
        affiche_main(main2, 2)
        affiche_tapis(tapis1, 1);
        affiche_tapis(tapis2, 2);

        // DEBUT TOUR
        if (tour > 0) {
            await jarnac(tapis1, tapis2, main1, main2, tour)
        }

        if (tour % 2 == 0) {
            tapis = tapis1
            main = main1
            num = 1
        } else {
            tapis = tapis2
            main = main2
            num = 2
        }
    
        if (tour > 1) {
            main.pop()
            main = main.concat(piocher_x_lettres(1, sac))
            main.push("fin du mot")
        }

        await jouer_tour(tapis, main, num, tour)

        if (tour % 2 === 0){
            main1 = main
        } else {
            main2 = main  
        }

        console.log("Fin du tour du joueur " + (tour%2+1))

        tour ++

        if (tapis1.length === 8 || tapis2.length === 8) {
            end = true;
            affiche_tapis(tapis1, 1);
            affiche_tapis(tapis2, 2);
        }
        
    }
};

//------------------------------------------------------------------------------------------------------------------------

//INITIALISATION DU JEU
let sac =  [["A", 14],["B", 4],["C", 7],["D", 5],["E", 19],["F", 2],["G", 4],["H", 2],["I", 11],["J", 1],["K", 1],["L", 6],["M", 5],["N" , 9],["O" , 8],["P" , 4],["Q" , 1],["R" , 10],["S", 7],["T", 9],["U", 8],["V", 2],["W" , 1],["X" , 1],["Y" , 1],["Z", 2]]

let tapis1 = []
let tapis2 = []

//DEBUT DU JEU
let main1 = piocher_x_lettres(6, sac)
let main2 = piocher_x_lettres(6, sac)
main1.push("fin du mot")
main2.push("fin du mot")

await playGame(main1, main2, tapis1, tapis2)

affiche_fin(tapis1,tapis2)
