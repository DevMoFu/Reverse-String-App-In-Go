# Reverse-String-App-In-Go
This is a reverse string app made with Go

***Prerequisites:***
- Have [Go](https://golang.org/doc/install) installed
- Have [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) installed

***Instructions:***
1. git clone https://github.com/DevMoFu/Reverse-String-App-In-Go.git
2. `go build ReverseStringApp.go`
3. `ReverseStringApp -String "<String to reverse>"`

***Args***
```
Usage of ./ReverseStringApp:
  -String string
        String to mod (default "Defualt Value")
```

***Sample Output***
```
[]# ./ReverseStringApp -String "this is sparta"

Reversed String: 'atraps si siht'
```

***TODO:***
- (feature) Add paindrome detection
- (feature) Capitalize case for every other letter