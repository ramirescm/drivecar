package folders

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ramirescm/drivecar/internal/files"
)

func (h *handler) GetAll(rw http.ResponseWriter, r *http.Request) {
	c, err := GetRootFolderContent(h.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	fc := FolderContent{Folder: Folder{Name: "root"}, Content: c}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(fc)
}

func getRootSubFolders(db *sql.DB) ([]Folder, error) {
	stmt := `select * from "folders" where "parent_id" is null and "deleted" = false`
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}

	f := make([]Folder, 0)
	for rows.Next() {
		var folder Folder
		err := rows.Scan(&folder.ID, &folder.ParentID, &folder.Name, &folder.CreatedAt, &folder.ModifiedAt, &folder.Deleted)
		if err != nil {
			continue
		}

		f = append(f, folder)
	}

	return f, nil
}

func GetRootFolderContent(db *sql.DB) ([]FolderResource, error) {
	subfolders, err := getRootSubFolders(db)
	if err != nil {
		return nil, err
	}

	fr := make([]FolderResource, 0, len(subfolders))
	for _, subfolder := range subfolders {
		r := FolderResource{
			ID:         subfolder.ID,
			Name:       subfolder.Name,
			Type:       "directory",
			CreatedAt:  subfolder.CreatedAt,
			ModifiedAt: subfolder.ModifiedAt,
		}

		fr = append(fr, r)
	}

	folderFiles, err := files.GetAllRootFiles(db)
	if err != nil {
		return nil, err
	}

	for _, f := range folderFiles {
		r := FolderResource{
			ID:         f.ID,
			Name:       f.Name,
			Type:       f.Type,
			CreatedAt:  f.CreatedAt,
			ModifiedAt: f.ModifiedAt,
		}

		fr = append(fr, r)
	}

	return fr, nil
}
