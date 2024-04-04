//TESTE UNITARIO

package enderecos

import "testing"

func TestTipoEndereco(t *testing.T) {
	enderecoTeste := "Avenida Paulista"
	tipoEnderecoEsperado := "Avenida"
	tipoEnderecoRecebido := TipoEndereco(enderecoTeste)

	if tipoEnderecoRecebido != tipoEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", tipoEnderecoEsperado, tipoEnderecoRecebido)
	}
}
