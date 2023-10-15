package sql

import (
	"context"
	"database/sql"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestCloseStmt(t *testing.T) {
	mockDb := initTestDatabase()

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "ok",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, _ := mockDb.Follower().BeginTx(context.Background(), "test", TxOptions{})
			stmt, _ := tx.Prepare("test", "SELECT * FROM farm")
			err := stmt.Close()
			if (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSelectStmt(t *testing.T) {
	mockDb := initTestDatabase()
	type farm struct {
		Id int64
	}
	farmRes := []farm{}

	type args struct {
		name string
		dest interface{}
		args interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				name: "testing select",
				dest: &farmRes,
				args: int64(1),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, _ := mockDb.Follower().BeginTx(context.Background(), "test", TxOptions{})
			stmt, _ := tx.Prepare("test", "SELECT id FROM farm WHERE id = ?")
			err := stmt.Select(tt.args.name, tt.args.dest, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetStmt(t *testing.T) {
	mockDb := initTestDatabase()
	var resCount int64

	type args struct {
		name string
		dest interface{}
		args interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				name: "testing select",
				dest: &resCount,
				args: int64(1),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, _ := mockDb.Follower().BeginTx(context.Background(), "test", TxOptions{})
			stmt, _ := tx.Prepare("test", "SELECT COUNT(*) FROM action WHERE id = ?")
			err := stmt.Get(tt.args.name, tt.args.dest, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestQueryRowStmt(t *testing.T) {
	mockDb := initTestDatabase()

	type args struct {
		name string
		args interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *sqlx.Row
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				name: "testing select",
				args: int64(1),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, _ := mockDb.Follower().BeginTx(context.Background(), "test", TxOptions{})
			stmt, _ := tx.Prepare("test", "SELECT COUNT(*) FROM action WHERE id = ?")
			_, err := stmt.QueryRow(tt.args.name, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryRow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestQueryStmt(t *testing.T) {
	mockDb := initTestDatabase()

	type args struct {
		name string
		args interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *sqlx.Rows
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				name: "testing select",
				args: int64(1),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, _ := mockDb.Follower().BeginTx(context.Background(), "test", TxOptions{})
			stmt, _ := tx.Prepare("test", "SELECT COUNT(*) FROM action WHERE id = ?")
			_, err := stmt.Query(tt.args.name, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestExecStmt(t *testing.T) {
	mockDb := initTestDatabase()

	type args struct {
		name string
		args interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    sql.Result
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				name: "testing select",
				args: int64(1),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, _ := mockDb.Follower().BeginTx(context.Background(), "test", TxOptions{})
			stmt, _ := tx.Prepare("test", "DELETE FROM FARM WHERE id = ?")
			_, err := stmt.Exec(tt.args.name, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tx.Rollback()
		})
	}
}
