package usecase

import (
	"orders-srv-go/mocks"
	"orders-srv-go/models"
	"orders-srv-go/repository"
	"reflect"
	"testing"
)

func Test_onlineOrders_ByClient(t *testing.T) {
	storeMock := mocks.NewStore(t)
	type fields struct {
		store repository.OrderStore
	}
	type args struct {
		clientID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.OnlineOrders
		wantErr bool
	}{
		{
			name: "invalid",
			fields: fields{
				store: storeMock,
			},
			args: args{
				clientID: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid",
			fields: fields{
				store: storeMock,
			},
			args: args{
				clientID: 1,
			},
			want:    models.OnlineOrders{models.OnlineOrder{ID: 1}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := onlineOrders{
				store: tt.fields.store,
			}
			got, err := u.ByClient(tt.args.clientID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}
