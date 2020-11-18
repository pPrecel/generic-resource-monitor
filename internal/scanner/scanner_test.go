package scanner

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestExtractFloat(t *testing.T) {
	fileBody := []string{
		"1 sda 2 121      fsd  s",
		"",
		"123  122      5   99",
	}
	type args struct {
		lines  []string
		line   int
		column int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "should return right value",
			args: args{
				lines:  fileBody,
				line:   2,
				column: 3,
			},
			want:    99,
			wantErr: false,
		},
		{
			name: "should return error when element is not a float",
			args: args{
				lines:  fileBody,
				line:   0,
				column: 1,
			},
			wantErr: true,
		},
		{
			name: "should return error when given line is higher than len of input strings",
			args: args{
				lines:  fileBody,
				line:   100,
				column: 0,
			},
			wantErr: true,
		},
		{
			name: "should return error when given column is higher than len of columns",
			args: args{
				lines:  fileBody,
				line:   0,
				column: 100,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractFloat(tt.args.lines, tt.args.line, tt.args.column)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractFloat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

var exampleFileData = []byte(`adssada                            12 12 42 54
fddsdfs          52          45 65 1
sdasd 2           1   4 69
666
3333 2 21.66 1`)

var exampleStringsData = []string{
	"adssada                            12 12 42 54",
	"fddsdfs          52          45 65 1",
	"sdasd 2           1   4 69",
	"666",
	"3333 2 21.66 1",
}

func TestReadLines(t *testing.T) {
	tmpFile, err := ioutil.TempFile("/tmp", "go-test-")
	failOnError(t, err)
	defer tmpFile.Close()
	_, err = tmpFile.Write(exampleFileData)
	failOnError(t, err)

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "should scan right number",
			args: args{
				filename: tmpFile.Name(),
			},
			wantErr: false,
			want:    exampleStringsData,
		},
		{
			name: "throw error when file doesn't exist",
			args: args{
				filename: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadLines(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadLines() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func failOnError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
		os.Exit(1)
	}
}
