package rpc

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform/terraform"
)

func TestUIInput_impl(t *testing.T) {
	var _ terraform.UIInput = new(UIInput)
}

func TestUIInput_input(t *testing.T) {
	client, server := testClientServer(t)
	defer client.Close()

	i := new(terraform.MockUIInput)
	i.InputReturnString = "foo"

	err := server.RegisterName("UIInput", &UIInputServer{
		UIInput: i,
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	input := &UIInput{Client: client, Name: "UIInput"}

	opts := &terraform.InputOpts{
		Id: "foo",
	}

	v, err := input.Input(opts)
	if !i.InputCalled {
		t.Fatal("input should be called")
	}
	if !reflect.DeepEqual(i.InputOpts, opts) {
		t.Fatalf("bad: %#v", i.InputOpts)
	}
	if err != nil {
		t.Fatalf("bad: %#v", err)
	}

	if v != "foo" {
		t.Fatalf("bad: %#v", v)
	}
}
