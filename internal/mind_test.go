package internal

import "testing"

func TestPmindToMarkdown(t *testing.T) {
	type args struct {
		jsonData []byte
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
				jsonData: []byte(""),
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PmindToMarkdown(tt.args.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("PmindToMarkdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PmindToMarkdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
