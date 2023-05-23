package bonus

import "errors"

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, error) {
	if len(caminhos) == 0 {
		return "", errors.New("not implemented yet")
	}
	contador := make(map[string]int)
	for i:= 0; i < len(contador); i++{
		for j:= 0; j < len(contador); j++{
			if caminhos[0][i] == caminhos[0][j]{
				contador[caminhos[0][i]] ++
			}
		}
	}
	for chave, valor := range contador {
		if valor == 1 {
			return chave, nil
		}
	}
	return "",nil
}
