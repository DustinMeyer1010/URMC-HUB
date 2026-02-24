<script lang="ts">

	import CopyButton from "../buttons/CopyButton.svelte";
    import { InfoStateClass } from "./states/InfoState.svelte";
	import { onMount } from "svelte";

    let {
        group
    } : {
        group: string
    } = $props();


    let InfoState: InfoStateClass = new InfoStateClass(group)

    onMount(async () => {
        await InfoState.GetGroupInfo()
    })


</script>

<section>
    <CopyButton value={InfoState.GroupInfo.information == "" ? "NA" : InfoState.GroupInfo.information} label="Information"/>
    <CopyButton value={InfoState.GroupInfo.description == "" ? "NA" : InfoState.GroupInfo.description} label="Description"/>
    <CopyButton value={InfoState.GroupInfo.ou == "" ? "NA" : InfoState.GroupInfo.ou} label="OU"/>
</section>


<style>

    section {
        display: flex;
        flex-direction: column;
        max-width: 80%;
        gap: 10px;
        padding: 20px;
        align-self: center;
        text-align: center;
        border-radius: 10px;
        background: var(--color-surface);
    }

</style>

