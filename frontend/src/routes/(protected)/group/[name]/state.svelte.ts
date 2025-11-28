import type { GroupPageInfo } from "@t/group";

export type Section = "INFO" | "ADD" | "REMOVE" | "MEMBERS"
export const Sections: Section[] = ["INFO", "ADD", "REMOVE", "MEMBERS"]

interface GroupPageState {
    loading: boolean,
    PageInfo: GroupPageInfo | null,
    section: Section,
    retrievePageData: (group: string) => void
    swapSection: (section: Section) => void

}


export class GroupPageStateClass implements GroupPageState {
    loading = $state(false)
    PageInfo: GroupPageInfo | null = $state(null)
    section: Section = $state("INFO")


    retrievePageData = (group: string) => {

    }

    swapSection = (newSection: Section) => {
        this.section = newSection
    }
}