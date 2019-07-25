# forth-to-forth

## Install

```
cd $GOPATH/src
git clone git@github.com:joneshf/forth-to-forth.git
cd forth-to-forth
```

## Build

```
go get
go build
echo "5 6 +" | ./forth-to-forth
# or
./forth-to-forth
10 11 -
2019/07/24 22:00:29 stack: []string{"10", "11", "-"}, output: []string{"-1"}
```

## Run Tests:

```
go test
```
