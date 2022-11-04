
# Gobuster 
Gobuster é uma ferramenta usada para força bruta:

- URIs (diretórios e arquivos) em sites.
- Subdomínios DNS (com suporte a curingas).
- Nomes de hosts virtuais em servidores web de destino.
- Abra buckets do Amazon S3

## Modos disponíveis

- dir - o modo de força bruta de diretório clássico
- dns - Modo de força bruta do subdomínio DNS
- s3 - Enumerar buckets abertos do S3 e procurar listagens de existência e bucket
- gcs - Enumerar buckets abertos do Google Cloud
- vhost - modo de força bruta de host virtual (não é o mesmo que DNS!)
- fuzz - algum fuzzing básico, substitui a palavra-chave `FUZZ`

<details>
 <summary><h3>continuação</h3></summary>


<details>
  <summary><h3>Termux:</h3></summary>

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
 ```text
   
  
  <details>
  <summary><h3>Debian:</h3></summary>
          
 ```text
apt-get update -y
apt-get upgrade -y 
apt-get install golang git -y 
git clone https://github.com/OJ/gobuster 
mv gobuster /data/data/com.termux/files/usr/lib/go/src/ 
go mod init gobuster 
go get github.com/OJ/gobuster/v3/cli/cmd 
go run gobuster
```    
   
   <details>
  <summary><h3>Ubuntu:</h3></summary>

   ```    
Mais Informação embreve
    ```    

    
    
    

