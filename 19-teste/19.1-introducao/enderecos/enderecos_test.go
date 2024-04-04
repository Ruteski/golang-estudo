//TESTE UNITARIO
// comando para testar(na raiz onde tem o package main) -> go test ./..
// go test -v(verbose)
// t.Parallel() -> roda os testes em paralelo

// go test --cover -> para ver a cobertura de teste do projeto
// go test --converprofile <nomearquivo>.txt
// go tool cover --func=<nomearquivo>.txt (le o arquivo e envia o resultado no terminal)
// go tool cover --html=<nomearquivo>.txt (le o arquivo e envia o resultado no terminal)  **************

package enderecos

import "testing"

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoEndereco(t *testing.T) {
	t.Parallel()

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

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste quebrou")
	}
}
