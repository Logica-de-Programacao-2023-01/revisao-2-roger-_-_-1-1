package bonus

import "errors"

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, error) {
	if len(caminhos) == 0 {
		return "", errors.New("not implemented yet")
	}
	contador := make(map[string]int)
	if len(caminhos) == 1 {
		return caminhos[0][1], nil

	}
	for i := 0; i < len(caminhos); i++ {
		contador[string(caminhos[0][i])]++
		for j := 1; j < len(caminhos); j++ {
			contador[string(caminhos[i][j])]++
			contador[string(caminhos[len(caminhos)-1][len(caminhos)-1])]--
		}
	}
	resp := ""
	for chave, valor := range contador {
		if valor == -1 {
			resp = chave
			return chave, nil
		}
	}
	return resp, nil
}

