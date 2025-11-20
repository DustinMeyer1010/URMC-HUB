import { goto } from "$app/navigation";
import type { Section } from "@t/section";

interface UserStateInterface {
    currentSection: Section,
    loading: boolean,
    SwapSections: (section: Section) => void
    SetURL: () => void
}

export class UserStateClass implements UserStateInterface {
    currentSection: Section = $state("PROFILE");
    loading: boolean = $state(true)


    SwapSections = (section: Section) => {
        this.currentSection = section
        this.SetURL()
    };

    SetURL() {
        goto(`?section=${this.currentSection}`, { replaceState: true, keepFocus: true, noScroll: true })
    }
}