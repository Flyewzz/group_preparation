package pg

import (
	"database/sql"

	. "github.com/Flyewzz/group_preparation/models"
)

type UniversityStoreControllerPg struct {
	itemsPerPage int
	db           *sql.DB
}

func NewUniversityStoreControllerPg(itemsPerPage int, db *sql.DB) *UniversityStoreControllerPg {
	return &UniversityStoreControllerPg{
		itemsPerPage: itemsPerPage,
		db:           db,
	}
}

func (usc *UniversityStoreControllerPg) GetAll() ([]University, error) {
	rows, err := usc.db.Query("SELECT university_id, name from universities")
	var universities []University
	if err != nil {
		// 'universities' is an empty struct
		return universities, err
	}
	for rows.Next() {
		var u University
		rows.Scan(&u.Id, &u.Name)
		universities = append(universities, u)
	}
	defer rows.Close()
	return universities, nil
}

// func (sc *StoreController) GetByPage(page int) ([]Request, error) {
// 	requests := sc.GetAll()
// 	itemsPerPage := sc.itemsPerPage
// 	start := (page - 1) * itemsPerPage
// 	stop := start + itemsPerPage

// 	if start > (len(requests)-1) || start < 0 {
// 		return nil, errors.New("The incorrect page number.")
// 	}

// 	if stop > len(requests) {
// 		stop = len(requests)
// 	}

// 	return requests[start:stop], nil
// }

// func (sc *StoreController) GetById(id int) (*Request, error) {
// 	sc.mtx.Lock()
// 	Request, ok := sc.requests[id]
// 	sc.mtx.Unlock()
// 	if !ok {
// 		return nil, errors.New("A Request was not found")
// 	}
// 	return Request, nil
// }

// func (sc *StoreController) Add(Request *Request) int {
// 	sc.mtx.Lock()
// 	sc.currentId++
// 	sc.requests[sc.currentId] = Request
// 	sc.mtx.Unlock()
// 	return sc.currentId
// }

// func (sc *StoreController) RemoveById(id int) error {
// 	sc.mtx.Lock()
// 	_, ok := sc.requests[id]
// 	sc.mtx.Unlock()
// 	if !ok {
// 		return errors.New("A Request was not found")
// 	}
// 	sc.mtx.Lock()
// 	delete(sc.requests, id)
// 	sc.mtx.Unlock()
// 	return nil
// }

// func (sc *StoreController) RemoveAll() {
// 	sc.mtx.Lock()
// 	for id := range sc.requests {
// 		delete(sc.requests, id)
// 	}
// 	sc.mtx.Unlock()
// }
