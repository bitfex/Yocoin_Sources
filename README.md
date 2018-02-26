# yocoin
# Preparing Ubuntu

As root,

> add-apt-repository ppa:gophers/archive
> apt update
> apt install golang-1.9-go

Add user for builds:

> useradd -m dev

# Unpacking sources

As dev,

> mkdir -p ~/go/src/github.com`

Put Yocoin15 to /home/dev/go/src/github.com`

# Building sources


> su - dev
> cd ~/go/src/github.com/Yocoin15/Yocoin_Sources
> /usr/lib/go-1.9/bin/go build -o ~/test cmd/yocoin/*.go


Looks like:
> dev@ubuntu-s-1vcpu-1gb-blr1-01:~/go/src/github.com/Yocoin15/Yocoin_Sources$ /usr/lib/go-1.9/bin/go build -o ~/test cmd/yocoin/*.go



## Starting node

1. Create new account(-s)
```
    ./yocoin account new
```
 
2. Initialize private net
```
    ./yocoin init genesis.json
```

3. Run new node with console mode
```
     ./yocoin  --networkid 13 console
```
 


4. Chech mining base address
```
    > eth.coinbase
```
4. Start mining to generate DAG
```
    > miner.start()
```
