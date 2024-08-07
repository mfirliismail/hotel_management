# Gunakan image Go resmi sebagai base image
FROM golang:1.22

# Set direktori kerja di dalam container
WORKDIR /app

# Salin go.mod dan go.sum (jika ada) dan instal dependensi
COPY go.mod go.sum ./
RUN go mod download

# Salin kode sumber aplikasi ke dalam container
COPY . .

# Build aplikasi
RUN go build -o hotel_management

# Tentukan command yang dijalankan saat container berjalan
CMD ["./hotel_management"]
