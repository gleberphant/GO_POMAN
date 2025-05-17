package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath" // Para manipulação segura de caminhos de arquivo
)

// frontendDir é o caminho para a pasta de build do seu app React (ex: "meu-app-react/dist")
// Ajuste este caminho conforme a estrutura do seu projeto.
// Se o seu backend Go estiver em uma pasta e o frontend em outra (ex: ../meu-app-react/dist)
const frontendDir = "meu-app-react/dist" // IMPORTANTE: Ajuste este caminho!
const indexFile = "index.html"

func main() {
	fmt.Println("Iniciando servidor Go...")

	// 1. Endpoints da API
	http.HandleFunc("/api/mensagem", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"texto": "Olá do backend Go!"}`)
	})

	// 2. Servir arquivos estáticos do React (JS, CSS, imagens, etc.)
	// http.FileServer espera um http.FileSystem. http.Dir usa o sistema de arquivos local.
	// fileServer := http.FileServer(http.Dir(frontendDir))

	// Servir arquivos de assets. Se seus assets (JS/CSS) são referenciados como /assets/* no index.html
	// e estão em frontendDir/assets/*
	// Se o Vite (ou CRA) coloca os assets diretamente na raiz do 'dist', você pode não precisar de /assets/
	// Verifique a estrutura da sua pasta 'dist' e como o index.html referencia os arquivos.
	// Por exemplo, se o Vite gera <script type="module" crossorigin src="/assets/index-DR5S5pQ5.js">
	// então você precisa servir a pasta 'assets' que está dentro de 'frontendDir'
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(filepath.Join(frontendDir, "assets")))))

	// Servir outros arquivos estáticos da raiz da pasta de build do React (ex: favicon.ico, manifest.json)
	// Isso é um exemplo, você pode precisar adicionar mais ou nenhum, dependendo do seu build.
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(frontendDir, "favicon.ico"))
	})
	// Adicione outros arquivos específicos se necessário (ex: manifest.json, robots.txt)

	// 3. Servir o index.html do React para a rota raiz e qualquer outra rota não tratada (para o React Router)
	// Isso permite que o React Router controle a navegação no lado do cliente.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Verifica se o caminho solicitado corresponde a um arquivo existente na pasta de build.
		// Se for um arquivo (ex: um asset que não foi pego pelo Handle de /assets/), sirva-o.
		// Se não for um arquivo existente (ex: /produtos/123), sirva o index.html principal.
		requestedPath := filepath.Join(frontendDir, r.URL.Path)
		stat, err := os.Stat(requestedPath)

		// Se o arquivo não existe OU se o caminho é um diretório (para evitar listar o diretório)
		// sirva o index.html principal.
		if os.IsNotExist(err) || stat.IsDir() {
			http.ServeFile(w, r, filepath.Join(frontendDir, indexFile))
			return
		}

		// Se o arquivo existir e não for um diretório, sirva o arquivo solicitado.
		// Isso pode ser útil se alguns assets não estão sob /assets/ mas são referenciados diretamente.
		if err == nil {
			http.ServeFile(w, r, requestedPath)
			return
		}

		// Fallback final (embora o caso acima deva cobrir)
		http.ServeFile(w, r, filepath.Join(frontendDir, indexFile))
	})

	// Use uma porta diferente da do servidor de desenvolvimento do React
	port := ":80"
	fmt.Printf("Servidor Go rodando em http://localhost%s\n", port)
	fmt.Printf("Servindo arquivos do React de: %s\n", frontendDir)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
