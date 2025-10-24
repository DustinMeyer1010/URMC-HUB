
export type CopyState = {
    copied: string,
    timeout: ReturnType<typeof setTimeout> | null
}


export async function copyToClip(text: string, state: CopyState ) {
    if (state.timeout != null) {
            clearTimeout(state.timeout)
    }
    try {
        await navigator.clipboard.writeText(text);
        console.log("Copied to clipboard:", text);
    } catch (err) {
        console.error("Failed to copy:", err);
    }

    state.copied = text
    state.timeout = setTimeout(() => {
        state.copied = ""
        state.timeout = null
    }, 500)
}