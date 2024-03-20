package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: go run main.go <pasta_origem> <pasta_destino>")
		return
	}

	origem := os.Args[1]
	destino := os.Args[2]

	copiados, err := copiarArquivosJava(origem, destino)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Arquivos Java copiados com sucesso:")
	for _, arquivo := range copiados {
		fmt.Println(arquivo)
	}
	fmt.Printf("Total de arquivos copiados: %d\n", len(copiados))
}

func copiarArquivosJava(origem, destino string) ([]string, error) {
	// Cria a pasta de destino se ela não existir
	err := os.MkdirAll(destino, os.ModePerm)
	if err != nil {
		return nil, err
	}

	var copiados []string

	// Percorre recursivamente a pasta de origem
	err = filepath.Walk(origem, func(caminho string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Verifica se é um arquivo Java
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".java") {
			// Calcula o caminho relativo do arquivo em relação à pasta de origem
			caminhoRelativo, err := filepath.Rel(origem, caminho)
			if err != nil {
				return err
			}

			// Cria o caminho completo do arquivo no destino
			destinoArquivo := filepath.Join(destino, caminhoRelativo)

			// Cria a pasta de destino se ela não existir
			pastaDestino := filepath.Dir(destinoArquivo)
			err = os.MkdirAll(pastaDestino, os.ModePerm)
			if err != nil {
				return err
			}

			// Abre o arquivo de origem
			arquivoOrigem, err := os.Open(caminho)
			if err != nil {
				return err
			}
			defer arquivoOrigem.Close()

			// Cria o arquivo de destino
			arquivoDestino, err := os.Create(destinoArquivo)
			if err != nil {
				return err
			}
			defer arquivoDestino.Close()

			// Copia o conteúdo do arquivo de origem para o arquivo de destino
			_, err = io.Copy(arquivoDestino, arquivoOrigem)
			if err != nil {
				return err
			}

			// Adiciona o nome do arquivo copiado à lista de arquivos copiados
			copiados = append(copiados, caminhoRelativo)
			fmt.Println("Arquivo copiado:", caminhoRelativo)
		}

		return nil
	})

	return copiados, err
}
