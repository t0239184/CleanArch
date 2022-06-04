#/bin/sh
ID=$1
PASSWORD=$2

if [[ -z $ID || -z $PASSWORD ]]; then
    echo "Usage: $0 <id> <password>"
    exit 1
fi

curl -X POST -v http://localhost:8080/api/v1/user/$ID -H 'Content-Type: application/json' -d  "{\"id\":\"$ID\",\"password\":\"$PASSWORD\"}"
