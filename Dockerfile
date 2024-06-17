FROM golang:1.19

WORKDIR /go/src

ENV PATH="/go/bin:${PATH}"

RUN go install github.com/golang/mock/mockgen@v1.5.0
RUN go install github.com/spf13/cobra-cli@latest
# RUN go install github.com/spf13/cobra@latest

RUN apt-get update && apt-get install -y sqlite3

# Configurando o usuário e permissões
RUN usermod -u 1000 www-data
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go /var/www/.cache /go/src

# Copiando o script de entrada
COPY entrypoint.sh /entrypoint.sh

USER www-data

# Configurando o ponto de entrada
ENTRYPOINT ["/entrypoint.sh"]
