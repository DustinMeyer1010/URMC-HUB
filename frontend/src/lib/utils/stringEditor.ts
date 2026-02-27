export const readableOU = (ou: string): string => {

    console.log(ou)

    if (ou.length == 0) {
        return ""
    }

    return ou.replaceAll("CN=", "").replaceAll("OU=", "").replaceAll(",DC=", ".").replaceAll(",", "*").split("*").reverse().join(" | ").replace(/\\+$/, "")
}