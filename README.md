# fxtabs

Go package to collect open tabs on all Mozilla Firefox windows (from the same
_profile_).

Tabs are collected from `recovery.jsonlz4` file, where Firefox uses as a
persistent backup of open tabs, back and forward button pages, cookies, forms,
and other session data.

This file is written almost in real time (there will be only some seconds delay)
whenever there is a browsing/tabs action.

## Usage

Import package in your project:

```
go get github.com/zanardo/fxtabs-go
```

Collect open tabs:

```go
package main

import "github.com/zanardo/fxtabs-go"

func main() {
    tabs, err := fxtabs.OpenTabs("/path/to/sessionstore-backups/recovery.jsonlz4")
    if err != nil {
        panic(err)
    }
    for _, tab := range tabs {
        fmt.Printf("title: %s\n", tab.Title)
        fmt.Printf("url: %s\n", tab.URL)
    }
}
```
