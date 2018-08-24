// Copyright 2018 Otis Elevator Company. All rights reserved.
// Use of this source code is govered by the MIT license which
// can be found in the LICENSE file.

// Author: Jeremy Mill: jeremy.mill@otis.com

// Otis udp_rx software has been designed to utilize information
// security technology described Part 774 of the EAR Category 5 Part 2
// but has been made publicly available in accordance with Part 742.15(b)
// and is therefore not subject to U.S. export regulations.
// Before you download this software be aware that the country in which you
// are located may have restrictions related to the import, possession, use
// and/or reexport of encryption items.  It is your responsibility to comply
// with any applicable laws and regulations pertaining the import, possession,
// use and/or reexport of encryption items.
package main

import "fmt"

type connTimeoutError struct {
	err string
}

func (e *connTimeoutError) Error() string {
	return fmt.Sprintf("connection has timed out")
}
