# Simple Rest + Docker
The "Simple Rest" can be embedded in a golang docker but it's 700MB+, with a scratch Docker is 7MB.

Command to compile go source with static dependencies:

```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```