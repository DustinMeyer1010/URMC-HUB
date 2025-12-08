
import { goto } from "$app/navigation"
import type { GroupSimpleInfo } from "@t/group"


const EMPTY_GROUP_INFO: GroupSimpleInfo = {
    name: "",
    description: "",
    information: "",
    ou: ""
}

interface InfoState {
    GroupName: string
    GroupInfo: GroupSimpleInfo
    GetGroupInfo: () => void
}


export class InfoStateClass implements InfoState {
    GroupName: string;
    GroupInfo: GroupSimpleInfo = $state(EMPTY_GROUP_INFO)


    constructor(groupName: string) {
        this.GroupName = groupName
    }


    GetGroupInfo = async () => {
        await fetch(`http://localhost:8000/api/group/${this.GroupName}`)
        .then(async (res) => {
            if (res.status == 404) {
                goto("/", {keepFocus: false, invalidateAll: true, replaceState: true})
            }
            this.GroupInfo = await res.json()
        }).catch(() => {
            goto("/", {keepFocus: false, invalidateAll: true, replaceState: true})
        })

    }
}