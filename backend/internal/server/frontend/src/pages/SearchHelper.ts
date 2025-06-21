import axios from "axios";
import type { UserCardInfo } from "../models/User";
import type { ComputerCardInfo } from "../models/Computer";
import type { GroupCardInfo } from "../models/Group";



export const UserSearch = async (searchValue: string): Promise<UserCardInfo[]> => {
        let users: UserCardInfo[] = [];

        let valueEncoded = encodeURIComponent(searchValue)

        try {
            const response = await axios.get<UserCardInfo[]>(`http://localhost:8080/search/users/${valueEncoded}`, {});
                users = response.data
          } catch (error) {
            if (axios.isAxiosError(error)){
                if (error.response) {
                    console.log(error.response.data)
                }
            } else {
                console.error(error)
            }
            
          }
          return users 
}

export const GroupSearch = async (searchValue: string): Promise<GroupCardInfo[]> => {
    let groups: GroupCardInfo[] = []

    let valueEncoded = encodeURIComponent(searchValue)

    try {
        const response = await axios.get<GroupCardInfo[]>(`http://localhost:8080/search/groups/${valueEncoded}`, {});
        groups = response.data

    } catch (error) {
        if (axios.isAxiosError(error)){
                if (error.response) {
                    console.log(error.response.data)
                }
            } else {
                console.error(error)
            }
    }
    return groups
}


export const ComputerSearch = async (searchValue: string): Promise<ComputerCardInfo[]> => {
    let computers: ComputerCardInfo[] = []

    let valueEncoded = encodeURIComponent(searchValue) 

    try {
        const response = await axios.get<GroupCardInfo[]>(`http://localhost:8080/search/computers/${valueEncoded}`, {});
        computers = response.data

    } catch (error) {
    if (axios.isAxiosError(error)){
        if (error.response) {
            console.log(error.response.data)    
        } else {
            console.error(error)
        }
   }
    }
    return computers
}
