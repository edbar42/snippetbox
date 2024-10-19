# Snippetbox
Snippetbox is a web application designed to store and manage user-submitted code snippets. It leverages Go's templating engine for dynamic HTML rendering and utilizes MySQL for data storage. 

At the moment, the server can only handle HTTPS requests using self-signed certificates.

## ✨ Usage
### Pre-requisites
You need [Go](https://go.dev/dl/) and [MySQL](https://dev.mysql.com/downloads/installer/) (or equivalently, [MariaDB](https://mariadb.com/kb/en/where-to-download-mariadb/)) installed on your machine in order to run this project locally.

### Running the app
Clone this repository:
```bash
git clone git@github.com:edbar42/snippetbox.git
```

Go to the directory where you cloned the app in and run:
```bash
go mod tidy
```
Run the `make_cert.sh` script to generate self-signed TLS certificates:

```bash
./make_cert.sh
```
> [!WARNING]
> Since these certificates are self-signed, your browser will give you a warning when trying to access the website. You can tell your browser to proceed to the website.

Finally, run the files under `cmd/web/` using the `go` toolchain:
```bash
go run cmd/web/*
```
Your server should now be running on [`https://localhost/42069`](https://localhost:42069/) (nice).

## 🤝 Contributing
You can contribute by forking this repo and opening pull requests with your suggestion. 

All pull requests should be submitted to the main branch.

## ⚖️ License
This project is licensed under the MIT License - see the LICENSE file for details.

## :pray: Acknowledgements
Thanks to [Alex Edwards](https://www.alexedwards.net/) for illustrating the development process for this app in his book [Let's Go!](https://lets-go.alexedwards.net/)
