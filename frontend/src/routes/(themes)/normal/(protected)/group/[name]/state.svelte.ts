import type { Group } from "$lib/types/group";

export type Section = "ADD" | "REMOVE" | "MEMBERS"
export const Sections: Section[] = [ "ADD", "REMOVE", "MEMBERS"]

interface GroupPageState {
    loading: boolean,
    PageInfo: Group.PageInfo | null,
    section: Section,
    retrievePageData: (group: string) => void
    swapSection: (section: Section) => void

}


export class GroupPageStateClass implements GroupPageState {
    loading = $state(false)
    PageInfo: Group.PageInfo | null = $state(null)
    section: Section = $state("ADD")


    retrievePageData = async (group: string) => {
        await fetch(`http://localhost:8000/api/group/${group}`)
        .then(async (res) => {
            this.PageInfo = await res.json()
        })
    }

    swapSection = (newSection: Section) => {
        this.section = newSection
    }
}