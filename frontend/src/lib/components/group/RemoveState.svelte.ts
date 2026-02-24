import type { Modification } from "$lib/types/resutls"
import { User } from "$lib/types/user"

export class RemoveStateClass {
    SearchValue: string = $state("")
    Users: User.CardInfo[] = $state([])
    Results: Modification.Results[] = $state([])


    Search = async (e: SubmitEvent) => {
        e.preventDefault()
        await fetch(`http://localhost:8000/api/search/users/${this.SearchValue}`)
        .then(async (res) => this.Users = await res.json())
    }


    RemoveFromGroup = async (username: string, group: string): Promise<void> => {
        this.Results = []
        await fetch(
            `http://localhost:8000/api/user/${username}/memberof`,
            {
                method: "DELETE",
                mode: "cors",
                body: JSON.stringify({
                    groups: [group]
                })
            })
        .then(async (res) => this.Results = await res.json()) 
    }

}