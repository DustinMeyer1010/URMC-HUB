import type { Drive } from "$lib/types/drive"
import type { Group } from "$lib/types/group"

interface DriveState {
    Loading: boolean
    DrivesAccess: Drive.CardInfo[]
    GetDrives: () => void
}

export class DriveStateClass implements DriveState {
    Loading: boolean = $state(true)
    DrivesAccess: Drive.CardInfo[] = $state([])
    Groups: Group.CardInfo[] = $state([])
    Username: string;

    constructor(username: string) {
        this.Username = username
    }

    GetGroups = async () => {
        await fetch(`http://localhost:8000/api/user/${this.Username}/memberof`)
        .then(async (res) => this.Groups = await res.json())
    }

    GetDrives = async () => {
        this.Loading = true

        const names = this.Groups.map((group) => group.name)

        await fetch(
            "http://localhost:8000/api/drive/access",
            {
                method: "POST",
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(names)
            }
        )
        .then(async (res) => {
            this.DrivesAccess = await res.json()
        })
        
        this.Loading = false
    }
}