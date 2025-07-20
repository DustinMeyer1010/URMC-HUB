export const HandleCopy = (copyItem: string, copyItemValue: string , setCopied: React.Dispatch<React.SetStateAction<string>>, ) => {

    Copy(copyItemValue)

    setCopied(copyItem + " Copied")

    setTimeout(() => {
        setCopied("")
        }, 1000)
    }


export const Copy = (text: string) => {
    navigator.clipboard.writeText(text)
    .then(() => {

    })
    .catch((err) => {
        alert(err)
    })
}   