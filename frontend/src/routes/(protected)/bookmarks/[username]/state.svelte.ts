import { goto } from "$app/navigation";
import type { BookmarkT } from "$lib/types/bookmark";

export class BookmarkStateClass {
    Bookmarks: BookmarkT[] = $state([])
    Filter: string = $state("")
    SelectedAgentBookmarks: string = $state("")
    LoggedInAgent: string = $state("")
    AgentsWithBookmarks: string[] = $state([])
    AddFormVisiable: boolean = $state(false)

    EditMode: boolean = $state(false)
    Owner: boolean = $derived(this.LoggedInAgent == this.SelectedAgentBookmarks)
    FilteredBookmarks: BookmarkT[] = $derived.by(() => {
        if (this.Filter === "") {
            return this.Bookmarks
        }

        const lowerFilter = this.Filter.toLocaleLowerCase()
        return this.Bookmarks.filter((bookmark) => {
            return bookmark?.name?.toLocaleLowerCase().includes(lowerFilter) || bookmark?.description?.toLocaleLowerCase().includes(lowerFilter)
        })
    })


    Username: string;
    constructor(usernaem: string) {
        this.Username = usernaem
    }

    UpdateLocation = async () => {
        if (this.SelectedAgentBookmarks == "Generic") {
            await goto("/bookmarks", {replaceState: true, invalidateAll: true})
        }
        await goto(`/bookmarks/${this.SelectedAgentBookmarks}`, {replaceState: true, invalidateAll: true})
        this.Username = this.SelectedAgentBookmarks
        await this.UpdateBookmarks()
        this.EditMode = false
    }

    UpdateBookmarks = async () => {
        await fetch(`http://localhost:8000/api/bookmarks/${this.Username}`).then(async (res) => {
            this.Bookmarks = await res.json()
        })

    }

    UpdateAgentBookmarksList = async () => {
        await fetch("http://localhost:8000/api/bookmarks/all/agents").then(async (res) => {
            this.AgentsWithBookmarks = await res.json()
            this.AgentsWithBookmarks = this.AgentsWithBookmarks.filter((agentB) => agentB != this.LoggedInAgent)
        })
    }
    ShowForm = () => {
        this.AddFormVisiable = true
    }
    
    HideForm = () => {
        this.AddFormVisiable = false
    }
}