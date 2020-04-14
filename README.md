# Netsharer
## CLI for sharing files in a intranet

### Install

```
$ git clone https://github.com/sriosdev/netsharer
$ cd netsharer
$ go build -o netsharer
```

### Usage

#### Default (port 3000):

```
./netsharer -f path/to/file/or/folder
```

#### Specific port:
```
./netsharer -f path/to/file/or/folder -p 4040
```

#### Output:
```
  Your file is waiting at:
  - Download link: 192.168.0.150:3000

  Remember you can only access link once.
```
Copy the link given into your browser's location bar. You'll be able to download the file to any device connected to the same intranet.

### Arguments

* **-f**
   
   File or folder path. Must be a text string and it is required.
* **-p**

   Server port. Number of port. Opcional parameter (default 3000).
