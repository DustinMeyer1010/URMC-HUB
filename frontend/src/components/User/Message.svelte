<script lang="ts">
	import type { ModifyResults } from "@t/resutls";


    let {
        result
    } : { 
        result: ModifyResults
    } = $props()

    let message: string = $derived.by(() => {
        if (result.message.includes("Entry Already Exists")) {
            return "Already Memeber of Group"
        } else if (result.message.includes("Insufficient Access Rights")) {
            return "Insufficient Access Rights"
        } else if (result.message.includes("Unwilling To Perform")) {
            return "Not Apart of Group"
        } else {
            return result.message
        }
    })

</script>


<div class:failed={!result.successful}>
    <span>{result.group} | </span>
    <span>{message}</span>
</div>



<style>
    div {
        display: flex;
        flex-direction: row;
        gap: 10px;
        position: relative;
        border: 2px solid var(--color-success);
        color: var(--text-color);
        padding: 1rem;
        border-radius: 10px;
        font-weight: bold;
    }


    div.failed {
        border: 2px solid var(--color-error);
    }
</style>