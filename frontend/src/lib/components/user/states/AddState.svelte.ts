import type { Group } from "$lib/types/group"
import type { Modification } from "$lib/types/resutls"

interface AddState {
    Search: string
    Groups: Group.CardInfo[]
    Results: Modification.Results[]
    SearchGroup: (e: SubmitEvent) => void
    AddGroup: (username: string, group: string) => void
}


export class AddStateClass implements AddState {
    Search: string = $state("")
    Groups: Group.CardInfo[] = $state([])
    Results: Modification.Results[] = $state([])

    SearchGroup = async (e: SubmitEvent) => {
        e.preventDefault()
        await fetch(`http://localhost:8000/api/search/groups/${this.Search}`)
        .then(async (res) => {
            this.Groups = await res.json()
        })
    }

    AddGroup = async (username: string, group: string) => {
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

    }
}