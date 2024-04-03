package enderecos

import "strings"

func TipoEndereco(endereco string) string {
	tipoValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	primeiraPalavraEndereco := strings.Split(strings.ToLower(endereco), " ")[0]

	tipoEnderecoValido := false

	for _, tipo := range tipoValidos {
		if tipo == primeiraPalavraEndereco {
			tipoEnderecoValido = true
		}
	}

	if tipoEnderecoValido {
		return strings.Title(primeiraPalavraEndereco)
	}

	return "Tipo inválido"
}
