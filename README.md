# Govershell
Govershell is Reverse shell in golang
--
Change the ip and port on line 23 to yours.

![test](https://user-images.githubusercontent.com/68773572/211536224-37b8a194-1710-43f0-9be1-db4af23ff15c.png)
--
This Reverse Shell is compatible with windows only.


build:
```bash
GOOS=windows GOARCH=amd64 go build -o revsh.exe --ldflags "-H=windowsgui" main.go
```
