govindalu-client
----------------
Go client for vindalu


Example:

    import github.com/vindalu/go-vindalu-client

    c, _ := vindalu.NewClient("http://localhost:5454")

    item, err := c.Get("vserver", "testid1")
    ...