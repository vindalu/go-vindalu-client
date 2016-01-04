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

    c, _ := vindalu.NewClient("http://localhost:5454")

    item, err := c.Get("vserver", "testid1")
    ...
