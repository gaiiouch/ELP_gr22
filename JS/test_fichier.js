import fs from 'fs'

export const lecture_fichier = async()=>{
    const cheminFichier = './liste_francais.txt';

    let mots = fs.readFile(cheminFichier, 'utf8', (err, data) => {
        if (err) {
            console.error('Erreur de lecture du fichier:', err);
            return;
        }

        const mots = data.split(/\s+/);
        for(let i = 0; i < mots.length; i++){
            mots[i] = mots[i].toLowerCase()
        }
        console.log(mots)
        return mots        
    });
    return mots
}

let words 
console.log(words =  await lecture_fichier())
console.log("coucou")