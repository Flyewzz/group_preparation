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
	mock.ExpectQuery("SELECT (.+) FROM universities").
		WillReturnRows(correctedRows)
	// mock.ExpectCommit()

	defer db.Close()

	uc := NewUniversityControllerPg(5, db)
	actualUniversities, err := uc.GetAll(0)
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
