## Hyperledger Fabric Library

### 1. npm install

- /fabric-book/library/javascript

```bash
cd fabric-book/library/javascript
npm install
```

### 2. Enroll Admin & Register User

- in 'fabric-book/library'
  - because we'll make 'Wallet' on 'fabric-book/library'

```bash
cd ..
node javasciprt/enrollAdmin.js
node javascript/registerUser.js
```

### 3. Start Library Network

```bash
./startLibrary.sh
```

### 4. Start Server

- You need to change ip

```bash
node application/server.js
```

