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
		for j := 1; j < len(caminhos); j++ {
			if caminhos[0][i] == caminhos[0][j] {
				for chave := range caminhos {
					contador[string(chave)]++
				}
			}
		}
	}

	resp := ""
	for chave, valor := range contador {
		if valor == 1 {
			resp = chave
			return chave, nil
		}
	}
	return resp, nil
}

