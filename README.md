# serialnum

Utility to translate serial number between different formats.

# Install

    go install github.com/borud/serialnum/cmd/serialnum@latest

# Usage examples

From integer serial number

    $ serialnum -i 158916359710
	uint64 = 158916359710
    string = 000.037.2020.108.00039
    bytes  = {25 27 7e4 6c}

From xxx.xxx.xxxx.xxx.xxxxx format:

    $  bin/serialnum -s 000.037.2020.108.00039
    uint64 = 158916359710
    string = 000.037.2020.108.00039
    bytes  = {25 27 7e4 6c}