import type { UserSimpleInfo } from "@t/user";

export class MembersStateClass {
    Members: UserSimpleInfo[] = $state([])
    Page: number = $state(0)
    MembersLength: number = $derived(this.Members.length)
    Loading: boolean = $state(false)
    PagedMembers: UserSimpleInfo[] = $derived.by(() => {
        if (this.Filter.length <= 3) {
            if (this.Page == this.MembersLength) {
                return this.Members.slice(this.MembersLength-11, this.MembersLength-1)
            }
            return this.Members.slice(this.Page, this.Page+10)
        }
        return this.FilteredMembers() 
    })
    Group: string;

    Filter: string = $state("")


    constructor(group: string) {
        this.Group = group
    }

    FilteredMembers = (): UserSimpleInfo[] => {
        return this.Members.filter((member) => 
            member.name.toLowerCase().includes(this.Filter.toLocaleLowerCase()) 
        || member.username.toLowerCase().includes(this.Filter.toLowerCase())
        || member.email.replaceAll("_", " ").toLowerCase().includes(this.Filter.toLowerCase()))
    }

    GetMembers = async () => {
        this.Loading = true
        await fetch(`http://localhost:8000/api/group/${this.Group}/members`).then( async (res) => {
            this.Members = await res.json()
        })
        this.Members.sort((a,b) => a.name.localeCompare(b.name))
        this.Loading = false
    }

    NextPage = () => {
        this.Page = Math.min(this.Page+10, this.MembersLength)
        window.scrollTo(0,0)
    }

    PrevPage = () => {
        this.Page = Math.max(0, this.Page-10)
        window.scrollTo(0,0)
    }
}