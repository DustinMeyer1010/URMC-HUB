import { goto } from "$app/navigation";
import { redirect } from "@sveltejs/kit";
import type { Section } from "@t/section";
import type { UserFullInfo } from "$lib/types/user";

interface UserStateInterface {
    currentSection: Section,
    loading: boolean,
    SwapSections: (section: Section) => void
    SetURL: () => void
}

export class UserStateClass implements UserStateInterface {
    currentSection: Section = $state("PROFILE");
    loading: boolean = $state(true)
    pageData: UserFullInfo | null = $state(null)


    SwapSections = (section: Section) => {
        this.currentSection = section
        this.SetURL()
    };

    SetURL() {
        goto(`?section=${this.currentSection}`, { replaceState: true, keepFocus: true, noScroll: true })
    }

    UserFullInfo = async (username: string): Promise<UserFullInfo> => {

        const response: Response = await fetch(`http://localhost:8000/api/user/${username}`);

        if (!response.ok) {
            throw redirect(301, "/search")
        }

        const data: UserFullInfo = await response.json();

        return data

    }
}