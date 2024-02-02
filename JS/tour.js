import inquirer from 'inquirer'
import { affiche_main, affiche_tapis } from './affichage.js'

export const jouer_tour = async (tapis, main, num, tour) => {

    console.log("--------------- TON TOUR ---------------")
    affiche_main(main, num)
    affiche_tapis(tapis, num);

    let question_lign = [
        {
            type : 'list',
            name : 'ligne',
            message : 'Sur quelle ligne veux-tu écrire un mot ?',
            choices: Array.from({ length: tapis.length + 1 }, (_, index) => index + 1), //les choix sont les lignes déjà existantes ou écrire une nouvelle ligne
            lign_choice(val) {
                return val;
            },
        },
    ];

    // Use await to wait for player input before moving on
    let chosen_lign = await inquirer.prompt(question_lign).then((answers) => {
        let chosen_lign = answers["ligne"]
        return chosen_lign
    });

    if (chosen_lign-1 < tapis.length) {
        main.pop()
        main = main.concat(tapis[chosen_lign-1])
        main.push("fin du mot")
        tapis[chosen_lign-1] = []
    }
    
    let question_letter = [
        {
            type : 'list',
            name : 'lettre',
            message : 'Ecris ton mot lettre par lettre :',
            choices: main, //main du joueur + lettres déjà sur la ligne
            filter(val) {
                return val;
            },
        },
    ];
    let word = false
    let letters

    while(!word) {
        let end_word = false
        letters = []
        while (!end_word) {
            letters = await inquirer.prompt(question_letter).then((answers) => {
                if (answers["lettre"] === "fin du mot") {
                    end_word = true
                } else {
                    letters.push(answers['lettre'])
                    let index = main.indexOf(answers['lettre']);
                    if (index !== -1) {
                        main.splice(index, 1);
                    }
                }
                return letters
            });
            
        }

        let mot = ((letters.toString()).replace(/,/g, '')).toLowerCase();
        word = true
    }

    if (chosen_lign === tapis.length+1) {
        tapis.push([]);
    }

    for (let i = 0; i < letters.length; i++){
        tapis[chosen_lign-1].push(letters[i])
    }

    for (let i = 0; i < tapis.length; i++){
        if (tapis[i].length === 0){
            tapis.splice(i, 1)
        }
    }
}