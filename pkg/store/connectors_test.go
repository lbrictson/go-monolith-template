package store

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestConnect(t *testing.T) {
	type args struct {
		opts SqliteAdapterOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test ConnectSQLITE3 memory database",
			args: args{
				opts: SqliteAdapterOptions{
					InMemory: true,
				},
			},
			wantErr: false,
		},
		{
			name: "Test ConnectSQLITE3 file database",
			args: args{
				opts: SqliteAdapterOptions{
					InMemory: false,
					FileName: fmt.Sprintf("../../local/test_dbs/test_db_%v.db", uuid.New().String()),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := ConnectSQLITE3(tt.args.opts, true); got == nil {
				t.Errorf("ConnectSQLITE3() = %v, want not nil", got)
				if (err != nil) != tt.wantErr {
					t.Errorf("ConnectSQLITE3() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}
