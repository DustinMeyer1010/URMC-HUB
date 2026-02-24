
export namespace Copy {

    export const EMPTY_COPY_STATE: State =  {
        copied: "",
        timeout: null
    }

    export type State = {
        copied: string,
        timeout: ReturnType<typeof setTimeout> | null
    }

    export const ToClipboard = async (text: string, state: State ): Promise<void> => {
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
}

