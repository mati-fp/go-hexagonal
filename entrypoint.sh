#!/bin/bash

# Garantir que o diretório de trabalho tem as permissões corretas
chown -R www-data:www-data /go/src

# Verificar se o arquivo go.mod já existe
if [ -f "go.mod" ]; then
  echo "go.mod already exists. Skipping go mod init."
else
  echo "Running go mod init..."
  go mod init github.com/mati-fp/go-hexagonal || true

  # Verifica se o arquivo go.mod foi criado
  if [ -f "go.mod" ]; then
    echo "go.mod successfully created."
  else
    echo "Failed to create go.mod."
  fi
fi

# Mantém o container rodando
tail -f /dev/null
