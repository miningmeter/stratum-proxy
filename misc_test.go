package main

import (
	"testing"
)

func TestValidateIPV4(t *testing.T) {

	type testpair struct {
		ipv4   string
		result bool
	}

	var tests = []testpair{
		{"", false},
		{"a.0.0.1", false},
		{"127.0.0", false},
		{"127.0.0.1.1", false},
		{"-1.0.0.1", false},
		{"127.-1.0.1", false},
		{"127.0.-1.1", false},
		{"127.0.0.-1", false},
		{"256.0.0.1", false},
		{"127.256.0.1", false},
		{"127.0.256.1", false},
		{"127.0.0.256", false},
		{"127.0.0.1", true},
	}

	for _, pair := range tests {
		r := ValidateIPV4(pair.ipv4)
		if r != pair.result {
			t.Error("For", pair.ipv4, "expected", pair.result, "got", r)
		}
	}
}

func TestValidateDNS(t *testing.T) {
	type testpair struct {
		dns    string
		result bool
	}

	var tests = []testpair{
		{"", false},
		{"example", false},
		{".example", false},
		{"example.", false},
		{".example.", false},
		{".", false},
		{"..", false},
		{"000", false},
		{"example-", false},
		{"-example", false},
		{"-example-", false},
		{"-", false},
		{"--", false},
		{"-.", false},
		{".-.", false},
		{"-.-", false},
		{"-.-.", false},
		{"example.com", true},
		{"en.multipool.us", true},
	}

	for _, pair := range tests {
		r := ValidateDNS(pair.dns)
		if r != pair.result {
			t.Error("For", pair.dns, "expected", pair.result, "got", r)
		}
	}
}

func TestValidatePort(t *testing.T) {

	type testpair struct {
		port   string
		result bool
	}

	var tests = []testpair{
		{"", false},
		{"a", false},
		{"111111", false},
		{"000001", false},
		{"0", false},
		{"99999", false},
		{"1", true},
		{"65535", true},
		{"01", true},
		{"001", true},
		{"0001", true},
		{"00001", true},
	}

	for _, pair := range tests {
		r := ValidatePort(pair.port)
		if r != pair.result {
			t.Error("For", pair.port, "expected", pair.result, "got", r)
		}
	}
}

func TestValidateAddr(t *testing.T) {

	type testpair struct {
		addr   string
		canDNS bool
		result bool
	}

	var tests = []testpair{
		{"a:a", false, false},
		{"aaa", false, false},
		{"a:a:a", false, false},
		{"aaa.0.0.1:1", false, false},
		{"999.0.0.1:1024", false, false},
		{"127.0.0.1:99999", false, false},
		{"127.0.0.1:1024", false, true},
		{":1024", true, false},
		{"example:1024", true, false},
		{".example:1024", true, false},
		{"example.:1024", true, false},
		{".example.:1024", true, false},
		{".:1024", true, false},
		{"..:1024", true, false},
		{"000:1024", true, false},
		{"example-:1024", true, false},
		{"-example:1024", true, false},
		{"-example-:1024", true, false},
		{"-:1024", true, false},
		{"--:1024", true, false},
		{"-.:1024", true, false},
		{".-.:1024", true, false},
		{"-.-:1024", true, false},
		{"-.-.:1024", true, false},
		{"example.com:1024", true, true},
		{"en.multipool.us:1024", true, true},
		{"127.0.0.1:1024", true, true},
	}

	for _, pair := range tests {
		r := ValidateAddr(pair.addr, pair.canDNS)
		if r != pair.result {
			t.Error("For", pair.addr, "expected", pair.result, "got", r)
		}
	}
}

func TestValidateHexString(t *testing.T) {

	type testpair struct {
		str    string
		result bool
	}

	var tests = []testpair{
		{"", false},
		{"/", false},
		{"jklm", false},
		{"0123456789abcdefABCDEF", true},
	}

	for _, pair := range tests {
		r := ValidateHexString(pair.str)
		if r != pair.result {
			t.Error("For", pair.str, "expected", pair.result, "got", r)
		}
	}
}
