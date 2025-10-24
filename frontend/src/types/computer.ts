/*
type ComputerSimpleInfo struct {
	Name            string `json:"name"`
	OU              string `json:"ou"`
	OperatingSystem string `json:"operating_system"`
}

*/

export type ComputerSimpleInfo = {
    name: string,
    ou: string,
    operating_system: string
}   