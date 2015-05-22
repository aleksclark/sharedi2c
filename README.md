# sharedi2c
embd-based i2c writer

provides a channel-based interface to an i2c bus for writing individual bytes, requires https://github.com/kidoman/embd

Example:
```
writer := sharedi2c.NewSharedWriter(port)
msg := sharedi2c.I2CMsg{Addr: addr, Value: data}
w.writer.SendMsg(msg)
```

As writers are created, they are either given a channel to an existing port, or a new channel is created if a port hasn't been used before.
