package pg

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/Flyewzz/group_preparation/models"
)

func getTestUniversities() []University {

	return []University{
		{
			Id:   1,
			Name: "МГТУ им Н.Э.Баумана",
		},
		{
			Id:   2,
			Name: "Финансовый университет при Правительстве РФ",
		},
		{
			Id:   3,
			Name: "МГУ им М.В.Ломоносова",
		},
		{
			Id:   4,
			Name: "ВШЭ",
		},
		{
			Id:   5,
			Name: "МФТИ",
		},
	}
}

// func TestStoreController_Add(t *testing.T) {
// 	type fields struct {
// 		currentId   int
// 		Universitys map[int]*University
// 	}
// 	type args struct {
// 		University *University
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   int
// 	}{
// 		{
// 			name: "Expected 4",
// 			fields: fields{
// 				currentId:   3,
// 				Universitys: make(map[int]*University),
// 			},
// 			args: args{
// 				University: NewUniversity(nil, nil),
// 			},
// 			want: 4,
// 		},
// 		{
// 			name: "Expected 1",
// 			fields: fields{
// 				currentId:   0,
// 				Universitys: make(map[int]*University),
// 			},
// 			args: args{
// 				University: NewUniversity(nil, nil),
// 			},
// 			want: 1,
// 		},
// 		{
// 			name: "Expected 2",
// 			fields: fields{
// 				currentId:   1,
// 				Universitys: make(map[int]*University),
// 			},
// 			args: args{
// 				University: NewUniversity(nil, nil),
// 			},
// 			want: 2,
// 		},
// 		{
// 			name: "Expected 9999",
// 			fields: fields{
// 				currentId:   9998,
// 				Universitys: make(map[int]*University),
// 			},
// 			args: args{
// 				University: NewUniversity(nil, nil),
// 			},
// 			want: 9999,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sc := newMockStoreController(
// 				2,
// 				tt.fields.currentId,
// 				tt.fields.Universitys,
// 			)
// 			if got := sc.Add(tt.args.University); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("StoreController.AddNew() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error("Cannot creating mock database\n")
	}
	columns := []string{
		"university_id",
		"name",
	}
	// mock.ExpectBegin()
	expectedUniversities := getTestUniversities()
	correctedRows := sqlmock.NewRows(columns)
	for _, u := range expectedUniversities {
		correctedRows.AddRow(u.Id, u.Name)
	}
	mock.ExpectQuery("SELECT (.+) from universities").
		WillReturnRows(correctedRows)
	// mock.ExpectCommit()

	defer db.Close()

	usc := NewUniversityStoreControllerPg(5, db)
	actualUniversities, err := usc.GetAll()
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
	if len(actualUniversities) != len(expectedUniversities) {
		t.Errorf("Expected %d users, but got %d\n", len(expectedUniversities), len(actualUniversities))
	}
	for i, u := range actualUniversities {
		if u.Id != expectedUniversities[i].Id {
			t.Errorf("Expected id '%d', but got '%d'\n", expectedUniversities[i].Id, u.Id)
		}
		if u.Name != expectedUniversities[i].Name {
			t.Errorf("Expected name '%s', but got '%s'\n", expectedUniversities[i].Name, u.Name)
		}
	}
}

// testUniversities := getTestUniversities()
// mapUniversitys := map[int]*University{
// 	1: &testUniversitys[0],
// 	2: &testUniversitys[1],
// 	3: &testUniversitys[2],
// 	4: &testUniversitys[3],
// }
// tests := []struct {
// 	name   string
// 	fields fields
// 	want   []University
// }{
// 	{
// 		name: "First adding",
// 		fields: fields{
// 			currentId:   4,
// 			Universitys: mapUniversitys,
// 		},
// 		want: testUniversitys,
// 	},
// }
// for _, tt := range tests {
// 	t.Run(tt.name, func(t *testing.T) {
// 		sc := newMockStoreController(
// 			2,
// 			tt.fields.currentId,
// 			tt.fields.Universitys,
// 		)
// 		if got := sc.GetAll(); !reflect.DeepEqual(got, tt.want) {
// 			t.Errorf("StoreController.GetAll() = %v, want %v", got, tt.want)
// 		}
// 	})
// }
// }

// func TestStoreController_GetById(t *testing.T) {
// 	testUniversitys := getTestUniversitys()
// 	mapUniversitys := map[int]*University{
// 		1: &testUniversitys[0],
// 		2: &testUniversitys[1],
// 		3: &testUniversitys[2],
// 		4: &testUniversitys[3],
// 	}
// 	type fields struct {
// 		currentId   int
// 		Universitys map[int]*University
// 	}
// 	type args struct {
// 		id int
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *University
// 		wantErr bool
// 	}{
// 		{
// 			name: "get 3",
// 			fields: fields{
// 				currentId:   4,
// 				Universitys: mapUniversitys,
// 			},
// 			args: args{
// 				id: 3,
// 			},
// 			want:    mapUniversitys[3],
// 			wantErr: false,
// 		},

// 		{
// 			name: "get 2",
// 			fields: fields{
// 				currentId:   4,
// 				Universitys: mapUniversitys,
// 			},
// 			args: args{
// 				id: 2,
// 			},
// 			want:    mapUniversitys[2],
// 			wantErr: false,
// 		},

// 		{
// 			name: "get 4",
// 			fields: fields{
// 				currentId:   4,
// 				Universitys: mapUniversitys,
// 			},
// 			args: args{
// 				id: 4,
// 			},
// 			want:    mapUniversitys[4],
// 			wantErr: false,
// 		},

// 		{
// 			name: "get 0",
// 			fields: fields{
// 				currentId:   4,
// 				Universitys: mapUniversitys,
// 			},
// 			args: args{
// 				id: 0,
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},

// 		{
// 			name: "get -5",
// 			fields: fields{
// 				currentId:   4,
// 				Universitys: mapUniversitys,
// 			},
// 			args: args{
// 				id: -5,
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sc := newMockStoreController(
// 				2,
// 				tt.fields.currentId,
// 				tt.fields.Universitys,
// 			)
// 			got, err := sc.GetById(tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("StoreController.GetById() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("StoreController.GetById() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestStoreController_RemoveById(t *testing.T) {
// 	testUniversitys := getTestUniversitys()
// 	mapUniversitys := map[int]*University{
// 		1: &testUniversitys[0],
// 		2: &testUniversitys[1],
// 		3: &testUniversitys[2],
// 		4: &testUniversitys[3],
// 	}
// 	type fields struct {
// 		currentId   int
// 		Universitys map[int]*University
// 	}
// 	type args struct {
// 		id int
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "First",
// 			fields: fields{
// 				currentId:   4,
// 				Universitys: mapUniversitys,
// 			},
// 			args: args{
// 				id: 2,
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Second",
// 			fields: fields{
// 				currentId:   4,
// 				Universitys: mapUniversitys,
// 			},
// 			args: args{
// 				id: 15,
// 			},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sc := newMockStoreController(
// 				2,
// 				tt.fields.currentId,
// 				tt.fields.Universitys,
// 			)
// 			if err := sc.RemoveById(tt.args.id); (err != nil) != tt.wantErr {
// 				t.Errorf("StoreController.RemoveById() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			if _, ok := tt.fields.Universitys[tt.args.id]; ok {
// 				t.Errorf("StoreController.RemoveById() expected 'ok' value: %v, but got %v", !ok, ok)
// 			}
// 		})
// 	}
// }

// func TestStoreController_GetByPage(t *testing.T) {
// 	// Items (Universitys) per page for testing
// 	itemsPerPage := 2
// 	type fields struct {
// 		itemsPerPage int
// 		currentId    int
// 		Universitys  map[int]*University
// 	}
// 	type args struct {
// 		page int
// 	}
// 	testUniversitys := getTestUniversitys()

// 	mapUniversitys := map[int]*University{
// 		1: &testUniversitys[0],
// 		2: &testUniversitys[1],
// 		3: &testUniversitys[2],
// 		4: &testUniversitys[3],
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    []University
// 		wantErr bool
// 	}{
// 		{
// 			name: "1st page",
// 			fields: fields{
// 				itemsPerPage: itemsPerPage,
// 				currentId:    4,
// 				Universitys:  mapUniversitys,
// 			},
// 			args: args{
// 				page: 1,
// 			},
// 			want:    testUniversitys[:itemsPerPage],
// 			wantErr: false,
// 		},

// 		{
// 			name: "2nd page",
// 			fields: fields{
// 				itemsPerPage: itemsPerPage,
// 				currentId:    4,
// 				Universitys:  mapUniversitys,
// 			},
// 			args: args{
// 				page: 2,
// 			},
// 			want:    testUniversitys[itemsPerPage : itemsPerPage+itemsPerPage],
// 			wantErr: false,
// 		},
// 		{
// 			name: "3rd page",
// 			fields: fields{
// 				itemsPerPage: itemsPerPage,
// 				currentId:    4,
// 				Universitys:  mapUniversitys,
// 			},
// 			args: args{
// 				page: 3,
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},

// 		{
// 			name: "0 (zero) page",
// 			fields: fields{
// 				itemsPerPage: itemsPerPage,
// 				currentId:    4,
// 				Universitys:  mapUniversitys,
// 			},
// 			args: args{
// 				page: 0,
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},

// 		{
// 			name: "-1 page",
// 			fields: fields{
// 				itemsPerPage: itemsPerPage,
// 				currentId:    4,
// 				Universitys:  mapUniversitys,
// 			},
// 			args: args{
// 				page: -1,
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sc := newMockStoreController(
// 				tt.fields.itemsPerPage,
// 				tt.fields.currentId,
// 				tt.fields.Universitys,
// 			)
// 			got, err := sc.GetByPage(tt.args.page)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("StoreController.GetByPage() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("StoreController.GetByPage() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestStoreController_RemoveAll(t *testing.T) {
// 	type fields struct {
// 		itemsPerPage int
// 		currentId    int
// 		Universitys  map[int]*University
// 	}

// 	testUniversitys := getTestUniversitys()

// 	mapUniversitys := map[int]*University{
// 		1: &testUniversitys[0],
// 		2: &testUniversitys[1],
// 	}

// 	tests := []struct {
// 		name   string
// 		fields fields
// 	}{
// 		{
// 			name: "Non-empty map",
// 			fields: fields{
// 				itemsPerPage: 2,
// 				currentId:    2,
// 				Universitys:  mapUniversitys,
// 			},
// 		},
// 		{
// 			name: "Empty map",
// 			fields: fields{
// 				itemsPerPage: 2,
// 				currentId:    0,
// 				Universitys:  make(map[int]*University),
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sc := newMockStoreController(
// 				tt.fields.itemsPerPage,
// 				tt.fields.currentId,
// 				tt.fields.Universitys,
// 			)
// 			sc.RemoveAll()
// 			if len(sc.Universitys) != 0 {
// 				t.Errorf("Expected 0 Universitys, but got: %d\n", len(sc.Universitys))
// 			}
// 		})
// 	}
// }
