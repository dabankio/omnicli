Omnicore cli wrapper
----

Control omnicored omnicore-cli from golang.


Zero dependency.

## Currently not support windows
Im developing under macOS and linux(ubuntu), not enough time to deal with windows.

## How to use?
- env variable `OMNI_BIN_PATH` to your [omnicore bin] path (see init.go
- `init.go`
- `func.go`
- other funcs normally start with name which same as file prefix (eg:func in `cli_wrap.go` usually like cliXxx)

Notice:
- Some behaviors may be effected by your local bitcoin.conf.



## LiCENSE
BSD 3-Clause License

