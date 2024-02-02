export {affiche_main, affiche_tapis};

const affiche_tapis = function(tapis, num_joueur) {
    let l = 0
    console.log("> tapis joueur " + num_joueur + " :")
    while (l < tapis.length) {
        console.log(tapis[l])
        l += 1
    }
}

const affiche_main = function(main, num_joueur) {
    if (main[main.length-1] == "fin du mot") {
        main.pop()
    }
    console.log("> main joueur " + num_joueur + " :")
    console.log(main)
    main.push("fin du mot")
}