# [Golang.hr] Platform configuration package
Base configuration package for [Golang.hr Platform]. Package is designed to be used with [etcd].

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

### Prerequisites
You will need to have fully functional [etcd] service setup and accessible by www.

### Installation
TBD

### Examples
Bellow you can find some useful examples of how to use this package

#### Example 1 - Initializing struct

```go
manager, err := NewManager(map[string]interface{}{
	"env": "environment_name", // Useful for if you wish to have "sandbox", "production" or any other
	"etcd": map[string]interface{}{
	    // We use it as project name. Or so, golanghr
		"folder":                     "folder_name",
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
