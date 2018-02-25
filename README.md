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

admin.addPeer("enode://5c41098d930079d2880b6f6ff48b6f5a2cd0315d60d4fadf1b3f05049772593050e4c88cb1d6a878a5e1c5ae67fd010e7156f9f04711b5243b8a50337dfdf0d9@67.205.183.161:30303")
admin.addPeer("enode://850789000e6a5dba7962f48df9922e24a40dc59a025bdb35953355bbf69721f856fd4e3f029e125a26153d62e15416471981c3ab8aeb274efb1763ddede790e3@67.205.162.228:30301")
admin.addPeer("enode://4ae29c3a038c5912b54ee961478483065dcc27a4331c23f77593e98d835ef6f4c7c8e651c4934299d916f036da7639c07177ea79e94bd2818a90acf1d50aae74@165.227.92.171:30301")
admin.addPeer("enode://0f3e72f6f5ab71f2154d394ab8b7a6d9eefef9fc4f4ebf2d40027f8802ada8260b069ba8ed7412c3e42d1491d3a66431e3bcd4ca233ba8eac46786859e2d9e69@165.227.92.67:30301")


4. Chech mining base address
```
    > eth.coinbase
```
4. Start mining to generate DAG
```
    > miner.start()
```
