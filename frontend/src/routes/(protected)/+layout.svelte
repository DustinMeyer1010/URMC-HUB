<script lang="ts">
	import Login from "@components/Login.svelte";
    import { onMount } from "svelte";


    let isLoggedIn: boolean = $state(true);

    let { children } = $props();
    onMount(async () => {
        const res = await fetch("http://localhost:8000/api/verify")

        if (res.status != 200) {
            isLoggedIn = false
            return
        }

        isLoggedIn = true


    })
</script>

{#if !isLoggedIn}
    <div>
        <Login/>
    </div>
{/if}
{@render children?.()}

<style>
    div {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        width: 100%;
        height: 100%;
        backdrop-filter: blur(5px);
        z-index: 3;
    }
</style>

