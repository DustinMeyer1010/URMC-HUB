<script lang="ts">

	import { InfoStateClass } from "./InfoState.svelte";
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
test
DriveSimpleInfoadfa

</section>


<style>

</style>