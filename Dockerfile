# Use la última imagen oficial de Golang
FROM golang:latest

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia los archivos go.mod y go.sum a /app
COPY go.mod ./
COPY go.sum ./

# Descarga las dependencias usando go mod
RUN go mod download

# Copia todo el contenido de tu proyecto, incluyendo el directorio cmd
COPY . .

# Compila la aplicación Go en un archivo ejecutable llamado server
RUN go build -o server .

# Expone el puerto 8080
EXPOSE 8080

# Inicia la aplicación al ejecutar el archivo server
CMD [ "./server" ]