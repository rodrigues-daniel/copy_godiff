package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: go run main.go <pasta_origem> <arquivo_destinos>")
		return
	}

	origem := os.Args[1]
	arquivoDestinos := os.Args[2]

	destinos, err := lerDestinos(arquivoDestinos)
	if err != nil {
		log.Fatal(err)
	}

	for _, destino := range destinos {
		err := copiarArquivosJava(origem, destino)
		if err != nil {
			log.Printf("Erro ao copiar arquivos para %s: %v\n", destino, err)
		} else {
			fmt.Printf("Arquivos Java copiados com sucesso para %s!\n", destino)
		}
	}
}

func lerDestinos(arquivoDestinos string) ([]string, error) {
	file, err := os.Open(arquivoDestinos)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var destinos []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		destino := strings.TrimSpace(scanner.Text())
		if destino != "" {
			destinos = append(destinos, destino)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return destinos, nil
}

func copiarArquivosJava(origem, destino string) error {
	// Cria a pasta de destino se ela não existir
	err := os.MkdirAll(destino, os.ModePerm)
	if err != nil {
		return err
	}

	//var copiados []string

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
			//copiados = append(copiados, caminhoRelativo)
			fmt.Println("Arquivo copiado:", caminhoRelativo)
		}

		return nil
	})

	return err
}
