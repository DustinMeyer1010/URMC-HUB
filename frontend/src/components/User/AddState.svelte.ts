import type { GroupSimpleInfo } from "@t/group"
import type { ModifyResults } from "@t/resutls"

interface AddState {
    Search: string
    Groups: GroupSimpleInfo[]
    Results: ModifyResults[]
    SearchGroup: (e: SubmitEvent) => void
    AddGroup: (username: string, group: string) => void
}


export class AddStateClass implements AddState {
    Search: string = $state("")
    Groups: GroupSimpleInfo[] = $state([])
    Results: ModifyResults[] = $state([])

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