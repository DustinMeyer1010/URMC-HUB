export function convertToReadableTime(time: number | string): string {
    const ticks = typeof time === 'string' ? BigInt(time) : BigInt(time);
    if (ticks === 0n) return "Must change at next logon";
    const epochOffset = 11644473600000n;
    const ms = ticks / 10000n;
    const unixMs = Number(ms - epochOffset);
    const date = new Date(unixMs);
    return date.toLocaleDateString('en-US', {
        month: 'long',
        day: 'numeric',
        year: 'numeric'
    });
}


function seperateMemberOfGroups(groups: any[]): string[] {
    let filter: string[] = []
    groups.forEach(group => {
        let temp = group.split(",")[0]
        temp = temp.replaceAll("CN=", "")
        filter.push(temp)
    });

    return filter as string[]
}
