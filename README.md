govindalu-client
----------------
Go client for vindalu

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
