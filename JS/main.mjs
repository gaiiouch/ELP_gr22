import random from 'random'
import inquirer from 'inquirer'

//github de la pizza : https://github.com/SBoudrias/Inquirer.js/tree/master/packages/inquirer/examples

//INITIALISATION DU JEU

// 14 A, 4 B, 7 C, 5 D, 19 E, 2 F, 4 G, 2 H, 11 I, 1 J, 1 K, 6 L, 5 M, 9 N, 8 O, 4 P, 1 Q, 10 R, 7 S, 9 T, 8 U, 2 V, 1 W, 1 X, 1 Y et 2 Z.
let sac =  [["A", 14],["B", 4],["C", 7],["D", 5],["E", 19],["F", 2],["G", 4],["H", 2],["I", 11],["J", 1],["K", 1],["L", 6],["M", 5],["N" , 9],["O" , 8],["P" , 4],["Q" , 1],["R" , 10],["S", 7],["T", 9],["U", 8],["V", 2],["W" , 1],["X" , 1],["Y" , 1],["Z", 2]]

//let tapis1 = [["B","O","N"],["J","E","U","X"]]
//let tapis2 = [["H","E","L","L","O"]]
let tapis1 = []
let tapis2 = []

//pas sûre que ce soit utile
//let alphabet = ["A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"]


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

//------------------------------------------------------------------------------------------------------------------------

//DEBUT DU JEU
let main1 = piocher_x_lettres(6, sac)
let main2 = piocher_x_lettres(6, sac)
main1.push("end word")
main2.push("end word")

const playGame = async () => {
    let end = false;
    let tour = 0;
    let tapis;
    let main;

    while (!end) {

        console.log("Tour du joueur " + (tour % 2 + 1))

        console.log("main du joueur 1")
        console.log(main1)
        console.log("main du joueur 2")
        console.log(main2)
        affiche_tapis(tapis1, 1);
        affiche_tapis(tapis2, 2);

        // DEBUT TOUR
        if (tour > 0) {

            console.log()



            let jarnac = [
                {
                    type : 'list',
                    name : 'Jarnac',
                    message : 'Do you want to say Jarnac ?',
                    choices: ['Yes', 'No'],
                    filter(val) {
                        return val;
                    },
                },
            ];

            let play_jarnac = await inquirer.prompt(jarnac).then((answers) => {
                // récapitulation de la réponse donnée par le joueur et on l'affiche,
                // mise à jour du tapis et de la main du joueur
                console.log('\nPlaying Jarnac');
                console.log(JSON.stringify(answers, null, '  '));
                let play_jarnac = answers["Jarnac"]
                return play_jarnac
            });

            if (play_jarnac === 'Yes'){
                if (tour % 2 == 0) {
                    tapis = tapis2
                    main = main2
                    console.log("Coup de Jarnac du joueur 1 sur le joueur 2 !")
                } else {
                    tapis = tapis1
                    main = main1
                    console.log("Coup de Jarnac du joueur 2 sur le joueur 1 !")
                }

                console.log("main de l'autre joueur ")
                console.log(main)
                affiche_tapis(tapis1, 1);
                affiche_tapis(tapis2, 2);

                let question_lign = [
                    {
                        type : 'list',
                        name : 'lign',
                        message : 'On which lign do you want to make changes ?',
                        choices: Array.from({ length: tapis.length}, (_, index) => index + 1), //les choix sont les lignes déjà existantes ou écrire une nouvelle ligne
                        lign_choice(val) {
                            return val;
                        },
                    },
                ];

                let chosen_lign = await inquirer.prompt(question_lign).then((answers) => {
                    // récapitulation de la réponse donnée par le joueur et on l'affiche,
                    // mise à jour du tapis et de la main du joueur
                    console.log('\nLign number');
                    console.log(JSON.stringify(answers, null, '  '));
                    let chosen_lign = answers["lign"]
                    return chosen_lign
                });

                main.pop()
                main = main.concat(tapis[chosen_lign-1])
                main.push("end word")
                console.log(main)
                tapis.splice(chosen_lign-1, 1)

                let question_letter = [
                    {
                        type : 'list',
                        name : 'letter',
                        message : 'Write your word letter per letter :',
                        choices: main, //main du joueur + lettres déjà sur la ligne
                        filter(val) {
                            return val;
                        },
                    },
                ];

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

                if (tour % 2 == 0) {
                    tapis = tapis1
                } else {
                    tapis = tapis2
                }
                
                tapis.push([]);
                for (let i = 0; i < letters.length; i++){
                    //poser_lettre(letters[i], tapis1, chosen_lign-1, tapis1[chosen_lign-1].length - 1); 
                    tapis[tapis.length-1].push(letters[i])
                }

            }

        }

        if (tour % 2 == 0) {
            tapis = tapis1
            main = main1
        } else {
            tapis = tapis2
            main = main2
        }

        if (tour > 1) {
            main.pop()
            main = main.concat(piocher_x_lettres(1, sac))
            main.push("end word")
        }

        console.log("main du joueur")
        console.log(main)
        affiche_tapis(tapis1, 1);
        affiche_tapis(tapis2, 2);

        let question_lign = [
            {
                type : 'list',
                name : 'lign',
                message : 'On which lign do you want to make changes ?',
                choices: Array.from({ length: tapis.length + 1 }, (_, index) => index + 1), //les choix sont les lignes déjà existantes ou écrire une nouvelle ligne
                lign_choice(val) {
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

        if (chosen_lign-1 < tapis.length) {
            main = main.concat(tapis[chosen_lign-1])
            tapis[chosen_lign-1] = []
        }
        
        let question_letter = [
            {
                type : 'list',
                name : 'letter',
                message : 'Write your word letter per letter :',
                choices: main, //main du joueur + lettres déjà sur la ligne
                filter(val) {
                    return val;
                },
            },
        ];

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

        if (tour % 2 === 0){
            main1 = main
            console.log("Fin du tour du joueur " + (tour%2+1))
        } else {
            main2 = main
            console.log("Fin du tour du joueur " + (tour%2+1))
        }

        tour ++

        if (tour == 4) {
            end = true;
            affiche_tapis(tapis1, 1);
            affiche_tapis(tapis2, 2);
        }
        
    }
};

playGame()