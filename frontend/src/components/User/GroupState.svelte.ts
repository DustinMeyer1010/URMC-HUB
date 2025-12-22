import type { CopyState } from "$lib/helper/copy.svelte";
import type { GroupSimpleInfo } from "@t/group";
import type { ModifyResults } from "@t/resutls";

interface GroupState {
    CopyState: CopyState
    Filter: string
    FilteredGroups: GroupSimpleInfo[]
    Results: ModifyResults[]
    RemoveGroup: (username: string, group: string) => void
    GetGroupForUser: (username: string) => void
}

export class GroupStateClass implements GroupState {
    CopyState: CopyState = $state({copied: "", timeout: null});
    Filter: string = $state("");
    Groups: GroupSimpleInfo[] = $state([]);
    Results: ModifyResults[] = $state([]);

    FilteredGroups: GroupSimpleInfo[] = $derived.by(() => {
        if (this.Filter == "") {
            return this.Groups
        }

        return this.Groups.filter((group) => group.name.toLowerCase().includes(this.Filter.toLowerCase()))
    })


    GetGroupForUser = async (username: string) => {
        await fetch(`http://localhost:8000/api/user/${username}/memberof`)
        .then(async (res) => this.Groups = await res.json()) 
    }


    RemoveGroup = async (username: string, group: string): Promise<void> =>  {
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
        await this.GetGroupForUser(username)
    }
}