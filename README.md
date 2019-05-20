Overview of application.

1. For solving the problem, the data structure used for representing parking lot is a mix of Heap and maps to solve all of the commands in optimal manner as possible, assuming space is not an issue.
Min Heap is used to store the empty slots and fetch the nearest slot in optimal time. 

2. No external testing library has been used as golang provides internal "testing" library. All test cases have been written using it.

3. script is written in `bin/setup` to install dependencies,run test case and build package. It will generate `bin/parking_lot` binary to run the application.

Follow this step to run application in Golang

1. To run this application, install Go.
2. Place the package in /go/src/github.com/ where GOPATH(/go/src).
3. Navigate to path $GOPATH/github.com/parking_lot/
3. Run `/bin/setup` script to install dependencies,run test case and build package.
4. To run application, run binary `/bin/parking_lot`.