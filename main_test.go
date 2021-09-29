package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	inStream := strings.NewReader(
		`%7B%22ETag%22:%20%226b223fb4296eb82082def6d431124381%22%2C%20%22Owner%22:%20%22minio%22%2C%20%22StorageClass%22:%20%22STANDARD%22%7D`)
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{inStream: inStream, outStream: outStream, errStream: errStream}
	args := strings.Split("", " ")

	if status := cli.Run(args); status != 0 {
		t.Errorf("expected %d to eq %d", status, 0)
	}

	expected := `{"ETag": "6b223fb4296eb82082def6d431124381", "Owner": "minio", "StorageClass": "STANDARD"}`
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestEncode(t *testing.T) {
	inStream := strings.NewReader(
		`{"ETag": "6b223fb4296eb82082def6d431124381", "Owner": "minio", "StorageClass": "STANDARD"}`)
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{inStream: inStream, outStream: outStream, errStream: errStream}
	args := strings.Split("ped -e", " ")

	if status := cli.Run(args); status != 0 {
		t.Errorf("expected %d to eq %d", status, 0)
	}

	expected :=
		`%7B%22ETag%22:%20%226b223fb4296eb82082def6d431124381%22%2C%20%22Owner%22:%20%22minio%22%2C%20%22StorageClass%22:%20%22STANDARD%22%7D`
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}
