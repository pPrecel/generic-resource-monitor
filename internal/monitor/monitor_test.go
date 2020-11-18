package monitor

import (
	"errors"
	"github.com/pPrecel/raspberrypi-file-monitor/internal/scanner"
	"reflect"
	"testing"
)

func TestMonitor_FireEntries(t *testing.T) {
	type fields struct {
		Metrics      Files
		readFile     func(string) ([]string, error)
		extractFloat func([]string, int, int) (float64, error)
	}
	tests := []struct {
		name    string
		fields  fields
		wantNil bool
	}{
		{
			name: "should return channel",
			fields: fields{
				Metrics: Files{{}, {}},
				readFile: func(s string) (strings []string, err error) {
					return []string{}, nil
				},
				extractFloat: func(strings []string, i int, i2 int) (f float64, err error) {
					return 0.0, nil
				},
			},
			wantNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Monitor{
				Metrics:      tt.fields.Metrics,
				readFile:     tt.fields.readFile,
				extractFloat: tt.fields.extractFloat,
			}
			if got := m.FireEntries(); (got != nil) != tt.wantNil {
				if got == nil {
				}
				t.Errorf("FireEntries() = %v, want (want nil?) %v", got, tt.wantNil)
			}
		})
	}
}

func TestMonitor_entryFunc(t *testing.T) {
	type fields struct {
		Metrics      Files
		readFile     func(string) ([]string, error)
		extractFloat func([]string, int, int) (float64, error)
	}
	type args struct {
		channel chan Channel
		entry   File
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		wantFloat float64
	}{
		{
			name: "should return channel with value",
			args: args{
				channel: make(chan Channel),
				entry: File{
					Metrics: []Metric{{}},
				},
			},
			fields: fields{
				readFile: func(s string) (strings []string, err error) {
					return []string{}, nil
				},
				extractFloat: func(strings []string, i int, i2 int) (f float64, err error) {
					return 12, nil
				},
			},
			wantFloat: 12,
			wantErr:   false,
		},
		{
			name: "should return channel with error when readFile throw error",
			args: args{
				channel: make(chan Channel),
				entry: File{
					Metrics: []Metric{{}},
				},
			},
			fields: fields{
				readFile: func(s string) (strings []string, err error) {
					return []string{}, errors.New("test error")
				},
				extractFloat: func(strings []string, i int, i2 int) (f float64, err error) {
					return 12, nil
				},
			},
			wantFloat: 0,
			wantErr:   true,
		},
		{
			name: "should return channel with error when extractFloat throw error",
			args: args{
				channel: make(chan Channel),
				entry: File{
					Metrics: []Metric{{}},
				},
			},
			fields: fields{
				readFile: func(s string) (strings []string, err error) {
					return []string{}, nil
				},
				extractFloat: func(strings []string, i int, i2 int) (f float64, err error) {
					return 12, errors.New("test error")
				},
			},
			wantFloat: 0,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Monitor{
				readFile:     tt.fields.readFile,
				extractFloat: tt.fields.extractFloat,
			}

			go m.entryFunc(tt.args.channel, tt.args.entry)

			event := <-tt.args.channel
			if event.Float != tt.wantFloat {
				t.Errorf("entryFunc() = %v, want %v", event.Float, tt.wantFloat)
			} else if tt.wantErr != (event.Error != nil) {
				t.Errorf("entryFunc() = %v, want nil", event.Error)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		files Files
	}
	tests := []struct {
		name string
		args args
		want *Monitor
	}{
		{
			name: "should return Monitor struct",
			args: args{
				files: Files{{}, {}, {}},
			},
			want: &Monitor{
				Metrics:      Files{{}, {}, {}},
				readFile:     scanner.ReadLines,
				extractFloat: scanner.ExtractFloat,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.files)
			if !reflect.DeepEqual(got.Metrics, tt.want.Metrics) {
				t.Errorf("New() = %v, want %v", got.Metrics, tt.want.Metrics)
			}
		})
	}
}
