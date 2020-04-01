package pg

import (
	"database/sql"

	"github.com/Flyewzz/group_preparation/errs"
	"github.com/Flyewzz/group_preparation/models"
	. "github.com/Flyewzz/group_preparation/models"
)

type MaterialFileControllerPg struct {
	db *sql.DB
}

func NewMaterialFileControllerPg(db *sql.DB) *MaterialFileControllerPg {
	return &MaterialFileControllerPg{
		db: db,
	}
}

// // MaterialFiles
// r.HandleFunc("/material/{id}/files", hd.AddMaterialFilesHandler).Methods("POST")
// r.HandleFunc("/material/{id}/files", hd.GetMaterialFilesHandler).Methods("GET")
// r.HandleFunc("/material/file/downloading", hd.MaterialFileDownloadHandler).Methods("GET")
// r.HandleFunc("/material/{id}/files/downloading", hd.MaterialFilesDownloadHandler).Methods("GET")

func (mfc *MaterialFileControllerPg) GetAll(materialId int) ([]MaterialFile, error) {
	rows, err := mfc.db.Query("SELECT file_id, name from materialfiles "+
		"WHERE material_id = $1", materialId)
	if err != nil {
		return nil, err
	}
	var files []MaterialFile
	for rows.Next() {
		var f MaterialFile
		err := rows.Scan(&f.Id, &f.Name)
		if err != nil {
			continue
		}
		files = append(files, f)
	}
	return files, nil
}

func (mfc *MaterialFileControllerPg) GetById(id int) (*MaterialFile, error) {
	row := mfc.db.QueryRow("SELECT file_id, name, path from materialfiles "+
		"WHERE file_id = $1", id)
	var f models.MaterialFile
	err := row.Scan(&f.Id, &f.Name, &f.Path)
	return &f, err
}

func (mfc *MaterialFileControllerPg) Add(name, path string, materialId int) (int, error) {
	var idFile int
	err := mfc.db.QueryRow("INSERT INTO materialfiles (name, path, material_id) "+
		"VALUES ($1, $2, $3) RETURNING file_id",
		name, path, materialId).Scan(&idFile)
	return idFile, err
}

func (mfc *MaterialFileControllerPg) RemoveById(id int) error {
	result, err := mfc.db.Exec("DELETE FROM materials WHERE material_id = $1", id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errs.MaterialDoesntExist
	}
	return err
}

func (mfc *MaterialFileControllerPg) RemoveAll(subjectId int) error {
	// #! Removed all the files too. Warning!
	_, err := mfc.db.Exec("DELETE FROM materials WHERE subject_id = $1", subjectId)
	return err
}
