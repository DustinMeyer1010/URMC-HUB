import type { LockoutInfo } from "@t/user";

interface LockoutState {
    LockoutInformation: LockoutInfo[]
    RefreshMsg: string
    RefreshTimeout: ReturnType<typeof setTimeout> | null
    Username: string
    RefreshLockout: () => void
}

export class LockoutStateClass implements LockoutState {
    Username: string;

    LockoutInformation: LockoutInfo[] = $state([])
    RefreshMsg: string = $state("")
    RefreshTimeout: NodeJS.Timeout | null = $state(null)
    loading: boolean = $state(false)

    constructor(username: string) {
        this.Username = username
    }



    RefreshLockout = async () => {
        if (this.RefreshTimeout != null) {
            clearTimeout(this.RefreshTimeout)
            this.RefreshMsg = ""
        }

        this.loading = true

        await fetch(`http://localhost:8000/api/user/${this.Username}/lockout`)
        .then(async (res) => {
            this.LockoutInformation = await res.json()

            // Sorts the Dates from ealiest to latest
            this.LockoutInformation.sort((a,b) => new Date(b.time).getTime() - new Date(a.time).getTime())
            this.loading = false
            this.RefreshMsg = "Refreshed"
            this.RefreshTimeout = setTimeout(() => {
                this.RefreshMsg = ""
            }, 1000)
        })
    }
}