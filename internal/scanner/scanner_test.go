package scanner

import (
	"io/ioutil"
	"testing"
)

var exampleFileData = []byte(`adssada 12 12 42 54
fddsdfs 52 45 65 1
sdasd 2 1 4 69
666
3333 2 21.66 1`)

func TestExtractVal(t *testing.T) {
	tmpFile, err := ioutil.TempFile("/tmp", "go-test-")
	failOnError(err)
	defer tmpFile.Close()
	_, err = tmpFile.Write(exampleFileData)
	failOnError(err)

	type args struct {
		filename string
		line     int
		elem     int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "should scan right number",
			args: args{
				filename: tmpFile.Name(),
				line:     1,
				elem:     2,
			},
			wantErr: false,
			want:    45,
		},
		{
			name: "throw error when file doesn't exist",
			args: args{
				filename: "",
				line:     0,
				elem:     0,
			},
			wantErr: true,
		},
		{
			name: "throw error when value is not float64",
			args: args{
				filename: tmpFile.Name(),
				line:     0,
				elem:     0,
			},
			wantErr: true,
		},
		{
			name: "throw error when file has not enough lines",
			args: args{
				filename: tmpFile.Name(),
				line:     666,
				elem:     0,
			},
			wantErr: true,
		},
		{
			name: "throw error when file has not enough elements",
			args: args{
				filename: tmpFile.Name(),
				line:     0,
				elem:     666,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractVal(tt.args.filename, tt.args.line, tt.args.elem)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractVal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractVal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}
