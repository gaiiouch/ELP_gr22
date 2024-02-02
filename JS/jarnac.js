import inquirer from 'inquirer'
import { affiche_main, affiche_tapis } from './affichage.js'

export const jarnac = async (tapis1, tapis2, main1, main2, tour) => {
    let main;
    let tapis;
    let num;
    
    let jarnac = [
        {
            type : 'list',
            name : 'Jarnac',
            message : 'Veux-tu dire Jarnac ?',
            choices: ['Oui', 'Non'],
            filter(val) {
                return val;
            },
        },
    ];

    let play_jarnac = await inquirer.prompt(jarnac).then((answers) => {
        let play_jarnac = answers["Jarnac"]
        return play_jarnac
    });

    if (play_jarnac === 'Oui'){
        if (tour % 2 == 0) {
            tapis = tapis2
            main = main2
            num = 2
            console.log("Coup de Jarnac du joueur 1 sur le joueur 2 !")
        } else {
            tapis = tapis1
            main = main1
            num = 1
            console.log("Coup de Jarnac du joueur 2 sur le joueur 1 !")
        }

        console.log("--------------- JARNAC ---------------")
        affiche_main(main, num)
        affiche_tapis(tapis, num);

        let question_lign = [
            {
                type : 'list',
                name : 'ligne',
                message : 'Sur quelle ligne veux-tu écrire un mot ?',
                choices: Array.from({ length: tapis.length + 1}, (_, index) => index + 1), //les choix sont les lignes déjà existantes ou écrire une nouvelle ligne
                lign_choice(val) {
                    return val;
                },
            },
        ];

        let chosen_lign = await inquirer.prompt(question_lign).then((answers) => {
            let chosen_lign = answers["ligne"]
            return chosen_lign
        });

        main.pop()
        main = main.concat(tapis[chosen_lign-1])
        main.push("fin du mot")
        tapis.splice(chosen_lign-1, 1)

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

        let end_word = false
        let letters = []
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

        if (tour % 2 == 0) {
            tapis = tapis1
        } else {
            tapis = tapis2
        }
        
        tapis.push([]);
        for (let i = 0; i < letters.length; i++){
            tapis[tapis.length-1].push(letters[i])
        }

    }
}