package gostl

import "testing"

func TestStack_Push(t *testing.T) {
	type fields struct {
		elements []StackT
	}
	type args struct {
		val StackT
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &Stack{
				elements: tt.fields.elements,
			}
			self.Push(tt.args.val)
		})
	}
}
