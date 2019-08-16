
# f7k

A golang framework

### Prerequisites

golang 1.12

### Installation

**Installation of sources**

    go get github.com/dev2choiz/f7k

**Installation of binary**

    go get -u github.com/dev2choiz/f7k/cmd/f7k


**Check the installation**

    f7k version

### Usage

**Create a new application**

    f7k create [-n | --name] [-i | --import-path] [-p | --port]



|&#45;n|&#45;&#45;name|name of the new project|
|:-------|:-------|:-------|
|&#45;p|&#45;&#45;import-path|example : github.com/dev2choiz/f7k|
|&#45;p|&#45;&#45;port|port number|

```bash
cd name # 'name' the name of the project
```

**installation of dependencies**  
```bash
go mod vendor
```

**Regenerate the cache**    
Some components need to regenerate the cache before they are executed.  
This is the case of the component github.com/dev2choiz/f7k/controllers which must generate in the cache the controllers referenced in the file ./conf/routes.yaml  
  
```bash
f7k cache --verbose
```
  
**Run the application**  
run with f7k in development : 
```bash
f7k run --verbose
```
this command line generate cache then execute main.go

run natively :
```bash
go run main.go --verbose
```

check :
```bash
curl localhost:8080
```

output :  
>Hello world !
