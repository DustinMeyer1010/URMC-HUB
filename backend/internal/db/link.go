package db

import "github.com/LostProgrammer1010/URMC-HUB/internal/models"

func AddLink(link models.Link) error {
	query := "INSERT INTO links (title, link, image_path, agent) VALUES (?,?,?,?)"

	_, err := DBConnection.Exec(query, link.Name, link.URL, link.ImagePath, UserID)

	return err
}

// Removes links from database
func RemoveLink(id int) error {

	query := "DELETE FROM links WHERE id=? AND agent = ?"

	_, err := DBConnection.Exec(query, id, UserID)

	return err

}

func UpdateLink(link models.Link) error {
	query := `
	UPDATE links 
	SET 
		title = ?, 
		link = ?, 
		image_path = ?
	WHERE 
		id = ? AND agent = ?
	`

	_, err := DBConnection.Exec(query, link.Name, link.URL, link.ImagePath, link.Id, UserID)

	return err
}
