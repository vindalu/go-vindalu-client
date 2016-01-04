govindalu-client
----------------
Go vindalu client


Requirements
------------

    Godep
    Go >= 1.4.2


Installing
----------
    
    go get github.com/vindalu/go-vindalu-client
    cd $GOPATH/src/github.com/vindalu/go-vindalu-client
    godep restore


Usage
-----
Start by creating a credentials file in your home directory under `~/.vindalu/credentials` similar to the contents shown below.

    {
        "auth": {
            "username": "...",
            "password": "..."
        }
    }

You can now start using the api.

Example:

    import github.com/vindalu/go-vindalu-client

    // Setup client
    c, _ := vindalu.NewClient("http://localhost:5454")

    // Get specific resource
    item, err := c.Get("vserver", "testid1")
    ...

    
    // List based on query/filter
    options := map[string]string{"from":"0","size":"500"}
    query := map[string]interface{"cpu_count": ">=4", "version": 234}
    items, err := c.List("vserver", options, query)
    ...

