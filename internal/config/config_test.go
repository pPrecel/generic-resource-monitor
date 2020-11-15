package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"reflect"
	"testing"
)

var exampleConfigFilename = "./example_input/config.yaml"

func TestReadConfig(t *testing.T) {
	tmpFile, err := ioutil.TempFile("/tmp", "go-test-")
	failOnError(err)
	defer tmpFile.Close()
	config := fixExampleConfig()
	bytes, err := yaml.Marshal(&config)
	failOnError(err)
	_, err = tmpFile.Write(bytes)
	failOnError(err)

	incorrectTmpFile, err := ioutil.TempFile("/tmp", "go-test-")
	failOnError(err)
	defer incorrectTmpFile.Close()
	_, err = incorrectTmpFile.Write([]byte(":)"))
	failOnError(err)

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{
			name:    "open example config and parse it correctly",
			args:    args{filename: tmpFile.Name()},
			wantErr: false,
			want:    config,
		},
		{
			name:    "throw error when file has incorrect data",
			args:    args{filename: incorrectTmpFile.Name()},
			wantErr: true,
			want:    Config{},
		},
		{
			name:    "throw error when file doesn't exist",
			args:    args{filename: "BAD_PATH"},
			wantErr: true,
			want:    Config{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadConfig(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func fixExampleConfig() Config {
	return Config{
		ServerPort:    ":4000",
		ServerAddress: "/metrics",
		Entries: []MonitorEntry{
			{
				MetricInfo: MetricInfo{
					Name:        "example_prometheus_entry",
					Description: "This is example prometheus entry",
					Timestamp:   5,
				},
				FileInfo: FileInfo{
					Filename:   "./example_input/thermal",
					LineNumber: 1,
					ElemNumber: 2,
				},
			},
		},
	}
}
