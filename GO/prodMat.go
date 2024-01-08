package main

// calcule le produit de deux matrices carrées de même taille
func ProdMat(taille int, A [taille][taille]int, B [taille][taille]int) (C [taille][taille]int) {

	for i := 0; i < taille; i++ {
		for j := 0; j < taille; j++ {
			C[i][j] = 0
			for k := 0; k < taille; k++ {
				C[i][j] = C[i][j] + A[i][k]*B[k][j]
			}
		}
	}
	return
}
