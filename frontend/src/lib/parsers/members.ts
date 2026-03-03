
export function seperateMemberOfGroups(groups: any[]): string[] {
    let filter: string[] = []
    groups.forEach(group => {
        let temp = group.split(",")[0]
        temp = temp.replaceAll("CN=", "")
        filter.push(temp)
    });

    return filter as string[]
}
