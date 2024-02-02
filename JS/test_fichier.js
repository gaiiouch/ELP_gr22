import fs from 'fs'

// Utilisez le module fs pour manipuler les fichiers
//const fs = require('fs');

// Spécifiez le chemin du fichier à lire
const cheminFichier = './JS/liste_francais.txt';

// Lisez le contenu du fichier de manière asynchrone
fs.readFile(cheminFichier, 'utf8', (err, data) => {
    if (err) {
        console.error('Erreur de lecture du fichier:', err);
        return;
    }

    // Divisez le contenu en mots (séparés par des espaces)
    const mots = data.split(/\s+/);
    console.log(typeof mots)
    // Affichez les mots
    console.log('Mots dans le fichier:', mots);
});