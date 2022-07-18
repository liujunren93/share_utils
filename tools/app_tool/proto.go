package main

import (
	"bufio"
	"bytes"
	"os"
	"regexp"
	"strings"
)

type Proto struct {
	Pkg     string
	GoPkg   string
	Imports []Proto
	Service
}
type Service struct {
	Name  string
	Funcs []string
}

func parseProto(p string) Proto {
	var proto Proto
	pOpen, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	re := bufio.NewReader(pOpen)
	for {
		buf, _, err := re.ReadLine()
		if err != nil {
			break
		}
		ok, err := regexp.Match("^package.*;", buf)
		if err != nil {
			panic(err)
		}
		if ok {
			proto.Pkg = GetPkg(buf)
		}

		ok, err = regexp.Match("option go_package.*;", buf)
		if err != nil {
			panic(err)
		}
		if ok {
			proto.GoPkg = GetGoPackage(buf)
		}
		if err != nil {
			panic(err)
		}
		ok, err = regexp.Match("import.*;", buf)
		if err != nil {
			panic(err)
		}
		if ok {
			proto.Imports = append(proto.Imports, GetImport(buf))
		}
		// services
		ok, err = regexp.Match(`service.*\{`, buf)
		if err != nil {
			panic(err)
		}
		var ser Service
		if ok {
			ser.Name = GetSerName(buf)
		}

	}
	return proto
}
func GetPkg(buf []byte) string {
	return strings.Trim(string(bytes.Split(buf, []byte(" "))[1]), ";")

}

func GetGoPackage(buf []byte) string {
	return strings.Trim(string(bytes.Split(buf, []byte(";"))[1]), "\"")
}

func GetImport(buf []byte) Proto {
	//import "proto/base.proto";
	ok, err := regexp.Match(".*\"", buf)
	if err != nil {
		panic(err)
	}
	if ok {
		buf = bytes.Trim(buf, "\"")
	}
	return parseProto(string(buf))
}
func GetSerName(buf []byte) string {
	return string(bytes.Trim(bytes.Split(buf, []byte(" "))[1], "{"))
}

func GetSerFunc(buf []byte) string {
	//  rpc List(ConfigListReq)returns(proto.DefaultRes);
	strings.Replace(s string, old string, new string, n int)
}
