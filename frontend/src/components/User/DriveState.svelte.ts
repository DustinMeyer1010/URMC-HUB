import type { DriveSimpleInfo } from "@t/drive"
import type { GroupSimpleInfo } from "@t/group"

interface DriveState {
    Loading: boolean
    DrivesAccess: DriveSimpleInfo[] | undefined
    GetDrives: (groups: GroupSimpleInfo[]) => void
}

export class DriveStateClass implements DriveState {
    Loading: boolean = $state(false)
    DrivesAccess: DriveSimpleInfo[] = $state([])

    GetDrives = async (groups: GroupSimpleInfo[]) => {
        this.Loading = true

        const names = groups.map((group) => group.name)

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