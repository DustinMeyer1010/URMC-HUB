import axios from "axios";
import type { UserCardInfo } from "../models/User";
import type { ComputerCardInfo } from "../models/Computer";
import type { GroupCardInfo } from "../models/Group";

// TODO
export const ComputerSearch = async (searchValue: string): Promise<ComputerCardInfo[]> => {
    let computers: ComputerCardInfo[] = [];
    
    try {
        const response = await axios.get<ComputerCardInfo[]>(`http://localhost:8080/search/users/${searchValue}`, {});
        computers = response.data
    } catch (error) {
        if (axios.isAxiosError(error)){
            if (error.response){

            console.log(error.response.data)
            }
            else {
                console.log(error)
            }
        }
        console.log(error)
    }

    return computers

}


export const UserSearch = async (searchValue: string): Promise<UserCardInfo[]> => {
        let users: UserCardInfo[] = [];
        try {
            const response = await axios.get<UserCardInfo[]>(`http://localhost:8080/search/users/${searchValue}`, {});
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

    let test = encodeURIComponent(searchValue)
    console.log(test)

    try {
        const response = await axios.get<GroupCardInfo[]>(`http://localhost:8080/search/groups/${searchValue}`, {});
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
