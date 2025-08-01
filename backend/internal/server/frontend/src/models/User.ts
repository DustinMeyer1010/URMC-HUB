export type UserCardInfo = {
    Type?: string;
    Name?: string,
    Username?: string,
    NetID?: string,
    Email?: string,
    OU?: string,
}


export type UserCardComponent = {
    Name?: string,
    Username?: string,
    NetID?: string,
    Email?: string,
    OU?: string,
    Select: (Type: string, Name: string) => void,
    ItemsSelected: string[],
}


/*	
    Username           string `json:"Username"`
	Name               string `json:"Name"`
	Email              string `json:"Email"`
	RelationshipStatus string `json:"RelationshipStatus"`
	Department         string `json:"Department"`
	Title              string `json:"Title"`
	Phone              string `json:"Phone"`
	Location           string `json:"Location"`
	LastPasswordSet    string `json:"LastPasswordSet"`
	URID               string `json:"URID"`
	OU                 string `json:"OU"`
	NetID              string `json:"NetID"`
	Description        string `json:"Description"` 
*/

export type UserInformation = {
    Name?: string,
    Username?: string, 
    NetID?: string, 
    Email?: string,
    OU?: string,
    RelationshipStatus: string,
    Department?: string, 
    Title?: string,
    Phone?: string,
    Location?: string,
    LastPasswordSet?: string,
    URID?: string, 
    Descirption?: string,
}