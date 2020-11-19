package usecase

import (
	"orders-srv-go/mocks"
	"orders-srv-go/models"
	"orders-srv-go/repository"
	"reflect"
	"testing"
	"time"
)

func Test_offlineOrders_GetAll(t *testing.T) {
	storeMock := mocks.NewStore(t)
	type fields struct {
		store repository.OrderStore
	}
	type args struct {
		limit          int
		offset         int
		createDateFrom time.Time
		createDateTo   time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.OfflineOrders
		want1   int
		wantErr bool
	}{
		{
			name: "invalid",
			fields: fields{
				store: storeMock,
			},
			args: args{
				limit:          0,
				offset:         0,
				createDateFrom: time.Time{},
				createDateTo:   time.Time{},
			},
			want:    nil,
			want1:   0,
			wantErr: true,
		},
		{
			name: "valid",
			fields: fields{
				store: storeMock,
			},
			args: args{
				limit:          0,
				offset:         0,
				createDateFrom: time.Now(),
				createDateTo:   time.Now(),
			},
			want:    models.OfflineOrders{models.OfflineOrder{Id: 1}},
			want1:   1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := offlineOrders{
				store: tt.fields.store,
			}
			got, got1, err := u.GetAll(tt.args.limit, tt.args.offset, tt.args.createDateFrom, tt.args.createDateTo)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetAll() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_offlineOrders_ByClient(t *testing.T) {
	storeMock := mocks.NewStore(t)
	type fields struct {
		store repository.OrderStore
	}
	type args struct {
		limit    int
		offset   int
		clientID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.OfflineOrders
		want1   int
		wantErr bool
	}{
		{
			name: "invalid",
			fields: fields{
				store: storeMock,
			},
			args: args{
				limit:    0,
				offset:   0,
				clientID: 0,
			},
			want:    nil,
			want1:   0,
			wantErr: true,
		},
		{
			name: "valid",
			fields: fields{
				store: storeMock,
			},
			args: args{
				limit:    0,
				offset:   0,
				clientID: 1,
			},
			want:    models.OfflineOrders{models.OfflineOrder{Id: 1}},
			want1:   1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := offlineOrders{
				store: tt.fields.store,
			}
			got, got1, err := u.ByClient(tt.args.limit, tt.args.offset, tt.args.clientID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByClient() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ByClient() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
