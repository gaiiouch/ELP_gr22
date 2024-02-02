export {affiche_fin}

const compte_pts = function (tapis) {
    let pts = 0
    for (let i = 0; i < tapis.length; i++){
        pts += (tapis[i].length)*(tapis[i].length)
    }
    return pts
}

const affiche_fin = function(tapis1, tapis2) {
    let j1_pts = compte_pts(tapis1)
    let j2_pts = compte_pts(tapis2)
    console.log("FIN DU JEU")
    if (j1_pts > j2_pts) {
        console.log("Le joueur 1 a gagné avec un total de "+j1_pts+" points !")
        console.log("Le joueur 2 a perdu avec un score de "+j2_pts+" points !")
    }
    if (j1_pts < j2_pts) {
        console.log("Le joueur 2 a gagné avec un total de "+j2_pts+" points !")
        console.log("Le joueur 1 a perdu avec un score de "+j1_pts+" points !")
    }

    if (j1_pts == j2_pts) {
        console.log("Egalité entre les deux joueurs avec "+j1_pts+" points chacun !")
    }
}