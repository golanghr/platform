# [Golang.hr] Platform configuration package
Base configuration package for [Golang.hr Platform]. Package is designed to be used with [etcd].

**NOTICE: Package under active development. No ETA atm...**

#### Why Etcd?

[etcd] we will just copy out what they said without much of complications...

```
etcd is a distributed, consistent key-value store for shared configuration and service discovery, with a focus on being:

Simple: curl'able user facing API (HTTP+JSON)
Secure: optional SSL client cert authentication
Fast: benchmarked 1000s of writes/s per instance
Reliable: properly distributed using Raft
```

You can see more at  [etcd github]

### Key Functionalities
- Configuration manager based on environment and etcd directory.
- Saving and updating configuration key/values directly on [etcd] instance
- Auto reload configuration once it's changed (watching)

**NOTICE: This package wrapper is designed to manage configuration. If you wish
to manage roles or anything else you will need to use custom logic. You can retrieve
Etcd client by ```manager.Etcd()```**

### Prerequisites
You will need to have fully functional [etcd] service setup and accessible by www.

### Installation
TBD

### Examples
Bellow you can find some useful examples of how to use this package

#### Example 1 - Initializing struct

```go
manager, err := NewManager(map[string]interface{}{
	 // Useful for if you wish to have "sandbox", "production" or any other
	"env": "environment_name",

	// We use it as project name. Or so, golanghr
	"folder":                     "folder_name",

	// Do we want to auto sync existing configuration from the etcd or not
	"auto_sync": true,

  // The recommended sync interval is 10 seconds to 1 minute, which does
  // not bring too much overhead to server and makes client catch up the
  // cluster change in time.
	"auto_sync_interval": 10 * time.Second,

	"etcd": map[string]interface{}{
		// list of etcd endpoints separated by comma
		"endpoints":                  []string{"127.0.0.1:2379"},
    // Transport is used by the Client to drive HTTP requests. If not
    // provided, DefaultTransport will be used.
		"transport":                  etcdc.DefaultTransport,
		// Username specifies the user credential to add as an authorization header
		"username":                   "",
    // Password is the password for the specified user to add as an authorization header
    // to the request.
		"password":                   "",
		// set timeout per request to fail fast when the target endpoint is unavailable
		"header_timeout_per_request": time.Second,
	},
})

if err != nil {
  panic(err)
}


```


[Golang.hr]: <https://github.com/golanghr>
[Golang.hr Platform]: <https://github.com/golanghr/platform>
[etcd]: <https://coreos.com/etcd/>
[etcd github]: <https://github.com/coreos/etcd>
