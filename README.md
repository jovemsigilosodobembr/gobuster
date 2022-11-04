
# Gobuster 
Gobuster é uma ferramenta usada para força bruta:

- URIs (diretórios e arquivos) em sites.
- Subdomínios DNS (com suporte a curingas).
- Nomes de hosts virtuais em servidores web de destino.
- Abra buckets do Amazon S3

## Tags, Statuses, etc

[![Build Status](https://travis-ci.com/OJ/gobuster.svg?branch=master)](https://travis-ci.com/OJ/gobuster) [![Backers on Open Collective](https://opencollective.com/gobuster/backers/badge.svg)](https://opencollective.com/gobuster) [![Sponsors on Open Collective](https://opencollective.com/gobuster/sponsors/badge.svg)](https://opencollective.com/gobuster)


## Love this tool? Back it!

If you're backing us already, you rock. If you're not, that's cool too! Want to back us? [Become a backer](https://opencollective.com/gobuster#backer)!

[![Backers](https://opencollective.com/gobuster/backers.svg?width=890)](https://opencollective.com/gobuster#backers)

All funds that are donated to this project will be donated to charity. A full log of charity donations will be available in this repository as they are processed.

# Changes

## 3.3

- Support TLS client certificates / mtls
- support loading extensions from file
- support fuzzing POST body, HTTP headers and basic auth
- new option to not canonicalize header names

## 3.2

- Use go 1.19
- use contexts in the correct way
- get rid of the wildcard flag (except in DNS mode)
- color output
- retry on timeout
- google cloud bucket enumeration
- fix nil reference errors

## 3.1

- enumerate public AWS S3 buckets
- fuzzing mode
- specify HTTP method
- added support for patterns. You can now specify a file containing patterns that are applied to every word, one by line. Every occurrence of the term `{GOBUSTER}` in it will be replaced with the current wordlist item. Please use with caution as this can cause increase the number of requests issued a lot.
- The shorthand `p` flag which was assigned to proxy is now used by the pattern flag

## 3.0

- New CLI options so modes are strictly separated (`-m` is now gone!)
- Performance Optimizations and better connection handling
- Ability to enumerate vhost names
- Option to supply custom HTTP headers

# License

See the LICENSE file.

# Manual

## Modos disponíveis

- dir - o modo de força bruta de diretório clássico
- dns - Modo de força bruta do subdomínio DNS
- s3 - Enumerar buckets abertos do S3 e procurar listagens de existência e bucket
- gcs - Enumerar buckets abertos do Google Cloud
- vhost - modo de força bruta de host virtual (não é o mesmo que DNS!)
- fuzz - algum fuzzing básico, substitui a palavra-chave `FUZZ`



```text
As opções padrão com códigos de status desabilitados são assim:
```

### Examples


### Examples

```text
gobuster s3 -w bucket-names.txt
```






<details>
  <summary><h3>Termux:</h3></summary>
  <details>

```text
termux-setup-storage 
termux-change-repo 
apt update -y
apt upgrade -y 
pkg install golang git -y 
git clone https://github.com/OJ/gobuster 
mv gobuster /data/data/com.termux/files/usr/lib/go/src/ 
go mod init gobuster 
go get github.com/OJ/gobuster/v3/cli/cmd 
go run gobuster
 ```

