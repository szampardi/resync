# resync

`sync.Once` with `Reset()` and `Error()`

  * See [sync.Once](http://golang.org/pkg/sync/#Once)

## Example

The following example examines how `resync.Once` could be used in a HTTP server situation.

```go
// use it just like sync.Once
var once resync.Once

// handle a web request
func handleRequest(w http.ResponseWriter, r *http.Request) {
	once.Do(func(){
		// load templates or something
	})
	// TODO: respond
}

// handle some request that indicates things have changed
func handleResetRequest(w http.ResponseWriter, r *http.Request) {
	once.Reset() // call Reset to cause initialisation to happen again above
}
```