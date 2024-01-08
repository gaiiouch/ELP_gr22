package main

func ProduitMatrices(a [][]int, b [][]int) [][]int {

	// produit de matrices carrées donc le résultat à la même taille que les matrices servant à effectuer le calcul
	n := len(a)
	resultat := make([][]int, n)

	for i := 0; i < n; i++ {
		resultat[i] = make([]int, n)

		for j := 0; j < n; j++ {
			resultat[i][j] = 0

			for k := 0; k < n; k++ {
				resultat[i][j] = resultat[i][j] + a[i][k]*b[k][j]
			}
		}
	}

	return resultat
}
