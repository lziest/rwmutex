# Subtleties in Go's Mutex Usage

In this package, you will find a `GoodCounter` which works as expected with multiple goroutines. And there is a `BadCounter` that only looks like a good mutex protected structure. Why? Please google "golang lock by value". :)
