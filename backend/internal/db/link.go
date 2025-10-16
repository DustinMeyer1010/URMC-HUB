package db

import "github.com/LostProgrammer1010/URMC-HUB/internal/models"

func AddLink(link models.Link) error {
	query := "INSERT INTO links (title, link, image_path, profile_id) VALUES (?,?,?,?)"

	_, err := DBConnection.Exec(query, link.Name, link.URL, link.ImagePath, UserID)

	return err
}

func RemoveLink(id int) error {

	query := "DELETE FROM links WHERE id=?"

	_, err := DBConnection.Exec(query, id)

	return err

}

func UpdateLink() {

}
