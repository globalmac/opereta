# Opereta
Console cross-platform utility to encrypt / decrypt phone numbers (and possibly other important information) using AES algorithm on GoLang

System requirements: **GoLang >= 1.10**

### Usage:

`Some help information about usage:`

```bash
./opereta_{OS} -h
```

`Encrypt phone number (arg encrypt - only INTEGER):`

```bash
./opereta_{OS} --encrypt=79012345678
```

`Decrypt phone number:`

```bash
./opereta_{OS} --decrypt=wCwrKYcT1552nZ3u690wV-PV7Kwo83cYYm05soyWIdM
```

`Stdout in JSON:`

```bash
./opereta_{OS} --encrypt/decrypt={VALUE} -json
```

### Notice!

> To simplify processing, the utility always returns status "1" when errors occur, in other cases, the cipher or decryption with status "0"

### Custom build:

**{GOOS} - OS:**

* Mac os - darwin
* Windows - windows
* Linux - linux
* FreeBSD - freebsd

**{GOARCH} - architecture:**

* x86_64 - amd64
* x86 - 386
* ARM - arm  (linux only)

**{APP_NAME}** - name of your app (like, app_os):


Example for Linux (amd64):

```bash
$ env GOOS=linux GOARCH=amd64 go build -o SUPER_APP main.go
```
