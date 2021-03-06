// Copyright 2018 Otis Elevator Company. All rights reserved.
// Use of this source code is govered by the MIT license which
// can be found in the LICENSE file.

// Author: Jeremy Mill: jeremy.mill@otis.com

// Otis udp_rx software has been designed to utilize information
// security technology described in the Category 5 – Part 2 of the
// Commerce Control List, within Part 774 of the Export Administration
// Regulations (“EAR”)(15 CFR 774).  However, the Otis udp_rx software
// has been made publicly available in accordance with Part 742.15(b)
// of the EAR and is therefore not subject to U.S. export regulations.
// Before downloading this software, be aware that the country in which
// you are located may have restrictions related to the import, download,
// possession, use and/or reexport of encryption items.  It is your
// responsibility to comply with any applicable laws and regulations
// pertaining the import, download, possession, use and/or reexport of
// encryption items.

package certcreator

import (
	"net"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestCreateCert(t *testing.T) {
	var outputpath, caKeyPath, caCertPath string
	if isWindows() {
		outputpath = "..\\keys\\server.key"
		caKeyPath = "..\\keys\\ca.key"
		caCertPath = "..\\keys\\ca.crt"
	} else {
		outputpath = "../keys/server.key"
		caKeyPath = "../keys/ca.key"
		caCertPath = "../keys/ca.crt"
	}
	err := CreateCert("server.crt", outputpath, caKeyPath, caCertPath, "")
	if err != nil {
		log.Fatal("failed to create/sign server.crt", err)
	}
}

func TestEncryptedCert(t *testing.T) {
	var outputpath, caKeyPath, caCertPath string
	if isWindows() {
		outputpath = "..\\keys\\encrypted_keys\\server.key"
		caKeyPath = "..\\keys\\encrypted_keys\\ca.key"
		caCertPath = "..\\keys\\encrypted_keys\\ca.crt"
	} else {
		outputpath = "../keys/encrypted_keys/server.key"
		caKeyPath = "../keys/encrypted_keys/ca.key"
		caCertPath = "../keys/encrypted_keys/ca.crt"
	}
	err := CreateCert("server.crt", outputpath, caKeyPath, caCertPath, "N0y#Xr7mwy")
	if err != nil {
		log.Fatal("failed to create/sign server.crt", err)
	}
}

func TestCreateCertInMemory(t *testing.T) {
	caKeyPath := "..\\keys\\ca.key"
	caCertPath := "..\\keys\\ca.crt"
	if !isWindows() {
		caKeyPath = strings.Replace(caKeyPath, "\\", "/", -1)
		caCertPath = strings.Replace(caCertPath, "\\", "/", -1)
	}
	var ips []net.IP
	ips = append(ips, net.IPv4(8, 8, 8, 8))
	var hostnames []string
	hostnames = append(hostnames, "example.com")
	newcert, newkey, err := CreateCertInMemory(caKeyPath, caCertPath, "", ips, nil)
	if err != nil {
		log.Fatal("failed to create certs in memory")
	}
	newcertstring := string(newcert)
	newkeystring := string(newkey)
	_ = newcertstring
	_ = newkeystring
}
