<script lang="ts">
	import { page } from "$app/state";
	import Profile from "$lib/components/user/Profile.svelte";
	import type { User } from "$lib/types/user";
	import { onMount } from "svelte";

    let username: string = $state("")
    let ou: string = $state("")
    let section: string = $state("")
    


    async function fetchUser(): Promise<User.PageInfo> {
        ou = page.url.searchParams.get("ou") ?? "NO_OU"
        const res = await fetch(`http://localhost:8000/api/user?ou=${encodeURIComponent(ou)}`);
        if (!res.ok) throw new Error('Failed to load user');
        return res.json();
    }



    onMount(async () => {
        username = page.url.searchParams.get("username") ?? "NO_USERNAME"
        section = page.url.searchParams.get("section") ?? "PROFILE"
    })

    let userPromise: Promise<User.PageInfo>  = fetchUser()

</script>

<!-- * Waits for the response for the server to render content -->
 <main>
    {#await userPromise}
        <p>Search for user in {ou}</p>
    {:then user} 
        {@render Content(user)}
    {:catch error}
        <p>{error.message}</p>
    {/await}
</main>

{#snippet Content(user: User.PageInfo)}
    <Profile username={user.username}></Profile>
    
{/snippet}
