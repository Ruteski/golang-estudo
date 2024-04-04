//TESTE UNITARIO

package enderecos

import "testing"

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoEndereco(t *testing.T) {
	cenariosTeste := []cenarioTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça dos Polacos", "Tipo inválido"},
		{"Estrada das Ortigueiras", "Estrada"},
		{"RUA DAS RUAS", "Rua"},
		{"rua do chisto", "Rua"},
		{"", "Tipo inválido"},
	}

	for _, cenario := range cenariosTeste {
		tipoEnderecoRecebido := TipoEndereco(cenario.enderecoInserido)
		if tipoEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava '%s' e recebeu '%s'", cenario.retornoEsperado, tipoEnderecoRecebido)
		}
	}
}
