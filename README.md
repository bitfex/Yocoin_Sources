# yocoin
# Preparing Ubuntu

As root,

```bash
add-apt-repository ppa:gophers/archive
apt update
apt install golang-1.9-go
```

Add user for builds:

```bash
useradd -m dev
```

# Unpacking sources

As dev,

```bash
mkdir -p ~/go/src/github.com`
```

Put Yocoin15 to `/home/dev/go/src/github.com`

# Building sources

```bash
su - dev
cd ~/go/src/github.com/Yocoin15/Yocoin_Sources
/usr/lib/go-1.9/bin/go build -o ~/test cmd/yocoin/*.go
```

Looks like:

```bash
dev@ubuntu-s-1vcpu-1gb-blr1-01:~/go/src/github.com/Yocoin15/Yocoin_Sources$ /usr/lib/go-1.9/bin/go build -o ~/test cmd/yocoin/*.go
```

## Starting node

1. Create new account(-s)

```bash
./yocoin account new
```
 
2. Initialize private net

```bash
./yocoin init genesis.json
```

3. Run new node with console mode

```bash
./yocoin  --networkid 13 console
```
 


4. Chech mining base address

```bash
eth.coinbase
```

5. Start mining to generate DAG

```bash
miner.start()
```

5. For Exchanges

https://github.com/ethereum/wiki/wiki/JSON-RPC#json-rpc-methods
