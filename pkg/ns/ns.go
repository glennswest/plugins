package ns

import (
        "os"
)

type netNS struct {
        file   *os.File
        closed bool
}


// Returns an object representing the namespace referred to by @path
func GetNS(nspath string) (NetNS, error) {
        err := IsNSorErr(nspath)
        if err != nil {
                return nil, err
        }

        fd, err := os.Open(nspath)
        if err != nil {
                return nil, err
        }

        return &netNS{file: fd}, nil
}

