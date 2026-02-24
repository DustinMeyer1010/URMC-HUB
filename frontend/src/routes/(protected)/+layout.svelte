<script lang="ts">
	import { isLoggedIn } from "$lib/stores/login";
	import Login from "$lib/components/form/Login.svelte";
	import Nav from "$lib/components/navigation/Nav.svelte";
	import { onMount } from "svelte";



    let { children } = $props();

    onMount(() => {
        isLoggedIn.CheckStatus();

        const interval = setInterval(() => {
            isLoggedIn.CheckStatus();
        }, 10000);

        return () => clearInterval(interval);
    });


    const login = () => {
        isLoggedIn.Login()
    }

    let IsLoggedIn: boolean = $derived.by(() => {
        return $isLoggedIn
    })


    
</script>

{#if !IsLoggedIn }
    <div>
        <Login login={login} />
    </div>
{/if}

<Nav/>
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

