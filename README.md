# goutils
Utility methods from go

## Creating a module for the project  
```$go mod init github.com/om-ganesh/goutils    ```

## Importing a module in your new project  
- Import the module in your go file  
``` import "github.com/om-ganesh/goutils" ```
- Use the desired public method(s)  
``` fmt.Println(goutils.Hello()) ```
- Add the module to your project from github  
(Note, go.mod, go.sum file will be created)  
```$go get github.com/om-ganesh/goutils ```
- Run your program  
```$go build ```

## Updating the module package
- Test what are the dependency packages are in your project  
```$go list -m all ```
- Check the list of versions available in remote repo  
```$go list -m -versions github.com/om-ganesh/goutils ```
- Get the upgrade/downgrade specific version  
```$go get github.com/om-ganesh/goutils@v0.2.0 ```  
(Alternatively, to update to the latest version)  
```$go get github.com/om-ganesh/goutils@latest ```
- Clean the stale version in go.sum  
```$go mod tidy ```

## Reference: Publishing a module
https://go.dev/doc/modules/publishing  
https://stackoverflow.com/questions/43716691/how-to-publish-a-go-package  
https://www.digitalocean.com/community/tutorials/how-to-distribute-go-modules  