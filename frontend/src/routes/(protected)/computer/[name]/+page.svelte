<script lang="ts">
	import type { ComputerPageInfo } from "@t/computer";
	import { onMount } from "svelte";

    let { data } : {data: {name: string}} = $props();

    let computer: ComputerPageInfo | null = $state(null)


    onMount(async () => {
        let res = await fetch(`http://localhost:8000/api/computer/info/${data.name}`)

        computer = await res.json()
    })
</script>



<h1>{computer?.computer_info.name}</h1>
<h1>{computer?.computer_info.ou}</h1>
<h1>{computer?.computer_info.operating_system}</h1>
<h1>{computer?.is_online}</h1>
<h1>{computer?.ping_results}</h1>


<style>

</style>
