package internal

import "testing"

func TestMdToPMind(t *testing.T) {
	type args struct {
		markdownStr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "demo1",
			args: args{
				markdownStr: `
思维导图
    新建节点1
        node1
    新建节点2
        node2
            n3
`,
			},
			want:    ``,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MdToPMind(tt.args.markdownStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("MdToPMind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MdToPMind() = %v, want %v", got, tt.want)
			}
		})
	}
}
