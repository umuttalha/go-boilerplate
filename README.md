For start the server

1. go install github.com/cosmtrek/air@latest

2. export PATH=$PATH:$(go env GOPATH)/bin

3. source ~/.zshrc


4. air

5. make db-start 


6. go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


7. export PATH=$PATH:$(go env GOPATH)/bin


8. make migrate-create


9. make migrate-up


10. make migrate-down


11. make migrate-version


12. make db-start

