export const Copy = (text: string) => {
    navigator.clipboard.writeText(text)
    .then(() => {

    })
    .catch((err) => {
        alert(err)
    })
}   