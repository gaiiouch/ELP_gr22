import random from 'random'
import inquirer from 'inquirer'

//github de la pizza : https://github.com/SBoudrias/Inquirer.js/tree/master/packages/inquirer/examples

//INITIALISATION DU JEU

// 14 A, 4 B, 7 C, 5 D, 19 E, 2 F, 4 G, 2 H, 11 I, 1 J, 1 K, 6 L, 5 M, 9 N, 8 O, 4 P, 1 Q, 10 R, 7 S, 9 T, 8 U, 2 V, 1 W, 1 X, 1 Y et 2 Z.
let sac =  [["A", 14],["B", 4],["C", 7],["D", 5],["E", 19],["F", 2],["G", 4],["H", 2],["I", 11],["J", 1],["K", 1],["L", 6],["M", 5],["N" , 9],["O" , 8],["P" , 4],["Q" , 1],["R" , 10],["S", 7],["T", 9],["U", 8],["V", 2],["W" , 1],["X" , 1],["Y" , 1],["Z", 2]]

let tapis1 = [["B","O","N"],["J","E","U","X"]]
let tapis2 = [["H","E","L","L","O"]]

//pas sûre que ce soit utile
let alphabet = ["A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"]


//FONCTIONS DU JEU
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

const affiche_tapis = function(tapis, num_joueur) {
    let l = 0
    console.log("tapis joueur " + num_joueur + " :")
    while (l < tapis.length) {
        console.log(tapis[l])
        l += 1
    }
}

//splice(start, deleteCount, item1, item2, /* …, */ itemN)

const poser_lettre = function(lettre, tapis, ligne, place_dans_la_ligne) {
    tapis[ligne].splice(place_dans_la_ligne, 0, lettre)
} 

const list_choix1 = function (tapis) {
    let choice1 = []
    for (let p = 0; p < tapis.length; p++) {
        choice1.push(p+1)
    }
    choice1.push("new lign")
    return choice1
}

//let choice1 = list_choix1(tapis1)

let main1 = piocher_x_lettres(6, sac)
let main2 = piocher_x_lettres(6, sac)
main1.push("end word")
main2.push("end word")

//DEBUT DU JEU

const playGame = async () => {
    let end = false;
    let i = 0;
    let tapis;
    let main;

    while (!end) {
        if (i % 2 === 0){
            tapis = tapis1
            main = main1
            console.log("Tour du joueur 1")
        } else {
            tapis = tapis2
            main = main2
            console.log("Tour du joueur 2")
        }

        console.log(main)
        affiche_tapis(tapis1, 1);
        affiche_tapis(tapis2, 2);

        let question_lign = [
            {
                type : 'list',
                name : 'lign',
                message : 'On which lign do you want to make changes ?',
                choices: Array.from({ length: tapis.length + 1 }, (_, index) => index + 1), //les choix sont les lignes déjà existante ou écrire une nouvelle ligne
                lign_choice(val) {
                    return val;
                    // on récupère le numéro de ligne pour mettre à jour les choix de la question d'après
                },
            },
        ];
        
        let question_letter = [
            {
                type : 'list',
                name : 'letter',
                message : 'Write your word letter per letter :',
                choices: main, //main du joueur + lettres déjà sur la ligne
                filter(val) {
                    //il faut une boucle while sur ces choix, CA VA ÊTRE COMPLIQUE JE PENSE
                    //et les lettres déjà posées sont enlevées de la liste choices
                    return val;
                },
            },
        ];

        // Use await to wait for player input before moving on
        let chosen_lign = await inquirer.prompt(question_lign).then((answers) => {
            // récapitulation de la réponse donnée par le joueur et on l'affiche,
            // mise à jour du tapis et de la main du joueur
            console.log('\nLign number');
            console.log(JSON.stringify(answers, null, '  '));
            let chosen_lign = answers["lign"]
            return chosen_lign
        });

        let end_word = false
        let letters = []
        while (!end_word) {
            letters = await inquirer.prompt(question_letter).then((answers) => {
                // récapitulation de la réponse donnée par le joueur et on l'affiche,
                // mise à jour du tapis et de la main du joueur

                if (answers["letter"] === "end word") {
                    end_word = true
                    console.log('\nTurn summary:');
                    console.log(JSON.stringify(letters, null, '  '));
                } else {
                    letters.push(answers['letter'])
                    let index = main.indexOf(answers['letter']);
                    if (index !== -1) {
                        main.splice(index, 1);
                    }
                }
                return letters
            });
            
        }

        if (chosen_lign === tapis.length+1) {
            tapis.push([]);
        }

        for (let i = 0; i < letters.length; i++){
            //poser_lettre(letters[i], tapis1, chosen_lign-1, tapis1[chosen_lign-1].length - 1); 
            tapis[chosen_lign-1].push(letters[i])
        }

        main.pop()
        main = main.concat(piocher_x_lettres(6-main.length, sac))
        main.push("end word")
        if (i % 2 === 0){
            main1 = main
            console.log("Fin du tour du joueur 1")
        } else {
            main2 = main
            console.log("Fin du tour du joueur 2")
        }

            
        i++;

        if (i == 4) {
            end = true;
            affiche_tapis(tapis1, 1);
            affiche_tapis(tapis2, 2);
        }
        
    }
};

playGame()