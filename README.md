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
A client can immediately be created which will give the user access to the read functions of the API.

Example:

    import github.com/vindalu/go-vindalu-client

    // Setup client
    c, _ := vindalu.NewClient("http://localhost:5454")

    // Get specific resource
    item, err := c.Get("vserver", "testid1", 0)
    ...

    
    // List based on query/filter
    options := map[string]string{"from":"0","size":"500"}
    query := map[string]interface{"cpu_count": ">=4", "version": 234}
    items, err := c.List("vserver", options, query)
    ...

In order to write to the API (Create/Update/Delete) there are two options to set client credentials.

1. Create a credentials file in your home directory under `~/.vindalu/credentials` containing the json below with your applicable username and password. This file will be read when your client is initialized.

        {
            "auth": {
                "username": "...",
                "password": "..."
            }
        }

1. The client itself has a SetCredentials call that can be made to attach your auth information to an existing client. This will override any credentials which may have been set via the credentials file outlined above.

    Example:

        // Setup client
        c, _ := vindalu.NewClient("http://localhost:5454")

        // Set credentials
        c.SetCredentials(user, pass)