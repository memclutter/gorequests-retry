# gorequests-retry

[![Go](https://github.com/memclutter/gorequests-retry/actions/workflows/go.yml/badge.svg)](https://github.com/memclutter/gorequests-retry/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/memclutter/gorequests-retry/branch/main/graph/badge.svg?token=1IWTNCLCAQ)](https://codecov.io/gh/memclutter/gorequests-retry)

Retry middleware for [memclutter/gorequests](https://github.com/memclutter/gorequests). 
Using [hashicorp/go-retryablehttp](https://github.com/hashicorp/go-retryablehttp) middleware allows you to retry requests when network errors occur or receive specified server status codes.

## Install

Installation using the go package system

```shell
go get github.com/memclutter/gorequets-retry
```

## Use

To use, pass to the `Use()` method of the `RequestInstance`

```go
package main

import (
	"github.com/memclutter/gorequests"
	"github.com/memclutter/gorequests-retry"
	"time"
)

func main() {
	retry := &gorequests_retry.Retry{RetryMax: 5, RetryWaitMin: 200 * time.Millisecond, RetryWaitMax: 700*time.Millisecond}
	err := gorequests.Get("http://example.com").Use(retry).Exec()
	// ...
}
```