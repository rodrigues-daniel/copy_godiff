package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	// Verifica se o número de argumentos é adequado
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run script.go pasta_origem pasta_destino")
		os.Exit(1)
	}

	// Captura os argumentos da linha de comando
	origem := os.Args[1]
	destino := os.Args[2]

	// Verifica se a pasta de origem existe
	if _, err := os.Stat(origem); os.IsNotExist(err) {
		fmt.Printf("A pasta de origem '%s' não existe.\n", origem)
		os.Exit(1)
	}

	// Verifica se a pasta de destino existe
	if _, err := os.Stat(destino); os.IsNotExist(err) {
		// Se a pasta de destino não existe, tenta criá-la
		if err := os.MkdirAll(destino, 0755); err != nil {
			fmt.Println("Erro ao criar a pasta de destino:", err)
			os.Exit(1)
		}
	}

	// Copia o conteúdo da pasta de origem para a pasta de destino
	err := filepath.WalkDir(origem, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Calcula o caminho de destino
		destPath := filepath.Join(destino, path[len(origem):])

		if d.IsDir() {
			// Se for um diretório, tenta criar na pasta de destino
			if err := os.MkdirAll(destPath, d.Type().Perm()); err != nil {
				return err
			}
		} else {
			// Se for um arquivo, copia para a pasta de destino
			if err := copyFile(path, destPath, d.Type().Perm()); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Erro ao copiar o conteúdo da pasta de origem para a pasta de destino:", err)
		os.Exit(1)
	}

	fmt.Println("Cópia concluída com sucesso!")
}

// Copia um arquivo de origem para um destino
func copyFile(src, dst string, perm fs.FileMode) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		return err
	}

	// Define as permissões do arquivo
	if err := out.Chmod(perm); err != nil {
		return err
	}

	return nil
}
