FROM golang:latest AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o portfolio

FROM scratch
WORKDIR /app
COPY --from=builder /app/portfolio /portfolio
COPY ./index.html ./index.html
COPY ./style.css ./style.css
COPY ./portfolio.js ./portfolio.js
COPY ./w-logo.svg ./w-logo.svg
COPY ./og-card.png ./og-card.png
COPY ./robots.txt ./robots.txt
COPY ./sitemap.xml ./sitemap.xml
COPY ./llms.txt ./llms.txt
COPY ./llms-full.txt ./llms-full.txt

ENTRYPOINT ["/portfolio"]
