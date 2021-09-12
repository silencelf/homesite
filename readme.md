# experimental project: Saving Management

## post savings:
curl -d '{"desc": "hello", "target": 100}' http://localhost:8080/savings -v -H "Content-Type: application/json"