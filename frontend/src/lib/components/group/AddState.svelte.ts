import type { Modification } from "$lib/types/resutls";
import type { User } from "$lib/types/user";

export class AddStateClass {
    SearchValue: string = $state("")
    Users: User.CardInfo[] = $state([])
    Results: Modification.Results[] = $state([])

    Search = async (e: SubmitEvent) => {
        e.preventDefault()
        await fetch(`http://localhost:8000/api/search/users/${this.SearchValue}`)
        .then(async (res) => this.Users = await res.json())
    }

    AddToGroup = async (username: string, group: string): Promise<void> => {
        this.Results = []
        const res = await fetch(
            `http://localhost:8000/api/user/${username}/memberof`, 
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                    groups: [group]
                })
            })
        
        this.Results = await res.json()

        setTimeout(() => {
            this.Results = []
        }, 5000);
    }
}