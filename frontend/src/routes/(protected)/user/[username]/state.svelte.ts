import { goto } from "$app/navigation";
import { redirect } from "@sveltejs/kit";
import { User } from "$lib/types/user";

interface UserStateInterface {
    currentSection: User.Section,
    loading: boolean,
    SwapSections: (section: User.Section) => void
    SetURL: () => void
}

export class UserStateClass implements UserStateInterface {
    currentSection: User.Section = $state("PROFILE");
    loading: boolean = $state(true)
    pageData: User.PageInfo | null = $state(null)


    SwapSections = (section: User.Section) => {
        this.currentSection = section
        this.SetURL()
    };

    SetURL() {
        goto(`?section=${this.currentSection}`, { replaceState: true, keepFocus: true, noScroll: true })
    }

    UserFullInfo = async (username: string): Promise<User.PageInfo> => {

        const response: Response = await fetch(`http://localhost:8000/api/user/${username}`);

        if (!response.ok) {
            throw redirect(301, "/search")
        }

        const data: User.PageInfo = await response.json();

        return data

    }
}