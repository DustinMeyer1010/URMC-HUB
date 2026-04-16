export const readableDN = (dn: string): string => {

    if (dn.length == 0) {
        return ""
    }

    return dn.replaceAll("CN=", "").replaceAll("OU=", "").replaceAll(",DC=", ".").replaceAll(",", "*").split("*").reverse().join(" | ").replace(/\\+$/, "")
}

export const pullNameFromDN = (dn: string): string => {
    if (dn.length == 0) {
        return ""
    }

    return dn.replaceAll("CN=", "").split(",")[0]
}