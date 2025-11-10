<script lang="ts">
	import { copyToClip, type CopyState } from "$lib/helper/copy.svelte";

 

    let copyState: CopyState = $state({
        copied: "",
        timeout: null
    })

    let {
        value,
        label = "",
    } : {
        value: string
        label: string
    } = $props()
    

</script>



<button onclick={() => copyToClip(value, copyState)}>
    <b>{label ? label : ""}:</b>
    {@html value == copyState.copied ? "Copied" : value}
</button>




<style>
    button {
        margin: 0;
        background: transparent;
        border: none; 
        padding: 0;
        color: var(--color-text);
        text-align: left;
        font-size: 15px;
        font-family: Roboto;
    }
</style>