<script lang="ts">
	import { goto } from "$app/navigation";
	import { page } from "$app/state";
	import { onMount } from "svelte";

    let groupName: string;
    let ou: string;



    onMount(() => {
        groupName = page.url.searchParams.get("name") ?? ""
        ou = page.url.searchParams.get("ou") ?? ""

        if (ou.length == 0 || groupName.length == 0 ){
            goto("/search")
        }

    })


</script>


<main>
    <h1>{groupName}</h1>
    {ou}
</main>





<style>

</style>