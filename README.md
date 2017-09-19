# Lockpick

Commandline tool for web application brute-force attack. 

## Usage
Project use [golang/dep](https://github.com/golang/dep) to manage dependencies. 

    > dep ensure
    
    > go run main.go --address http://localhost:8080/login.php \
    --dictionary /tmp/dict.txt \
    --username admin \
    --message "User or password incorrect" \
    --payload "{\"user\": \"{{username}}\", \"pass\": \"{{password}}\", \"Login\": \"Login\"}"

## Flags
*  -a, --address Full address to password form
*  -u, --username Payload template.
*  -d, --dictionary Path to dictionary file.
*  -p, --payload Payload template. 
    * Default "{\"username\": \"{{username}}\", \"password\": \"{{password}}\", \"Login\": \"Login\"}"
* -m, --message  Message after unsuccesful login attempt.
    * Default "Login failed"
    
    