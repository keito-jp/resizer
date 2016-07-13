package option_test

import (
	"github.com/go-microservices/resizer/option"
	"testing"
)

func TestFlags(t *testing.T) {
	o, err := option.New([]string{
		"--id", "AAAA",
		"--bucket", "BBBB",
		"--json", "CCCC",
		"--dbuser", "DDDD",
		"--dbpassword", "EEEE",
		"--dbprotocol", "FFFF",
		"--dbaddress", "GGGG",
		"--dbname", "HHHH",
		"--host", "IIII",
		"--host", "JJJJ",
	})
	if err != nil {
		t.Fatal(err)
	}
	if o.ProjectID != "AAAA" {
		t.Error("wrong ProjectID:", o.ProjectID)
	}
	if o.Bucket != "BBBB" {
		t.Error("wrong Bucket:", o.Bucket)
	}
	if o.JSON != "CCCC" {
		t.Error("wrong JSON:", o.JSON)
	}
	if o.DBUser != "DDDD" {
		t.Error("wrong DBUser:", o.DBUser)
	}
	if o.DBPassword != "EEEE" {
		t.Error("wrong DBPassword:", o.DBPassword)
	}
	if o.DBProtocol != "FFFF" {
		t.Error("wrong DBProtocol:", o.DBProtocol)
	}
	if o.DBAddress != "GGGG" {
		t.Error("wrong DBAddress:", o.DBAddress)
	}
	if o.DBName != "HHHH" {
		t.Error("wrong DBName:", o.DBName)
	}
	if o.Hosts[0] != "IIII" {
		t.Error("wrong Hosts[0]:", o.Hosts[0])
	}
	if o.Hosts[1] != "JJJJ" {
		t.Error("wrong Hosts[1]:", o.Hosts[0])
	}
}
