#/bin/sh
ID=$1

if [[ -z $ID ]]; then
    echo "Usage: $0 <id>"
    exit 1
fi

curl -X POST -v http://localhost:8080/api/v1/user/$ID/unlock
