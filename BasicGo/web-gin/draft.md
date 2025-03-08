
go mod init rathandev/web-gin 

go get .

curl http://localhost:8080/albums

curl http://localhost:8080/albums --request "POST" --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

curl -X POST "http://localhost:8080/albums" -d '{"id": "1", "title": "My Album", "artist": "Artist Name", "price": 9.99}'


