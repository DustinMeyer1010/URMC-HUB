<script lang="ts">
	import { page } from "$app/state";
	import CopyAllButton from "$lib/components/buttons/CopyAllButton.svelte";
	import CopyButton from "$lib/components/buttons/CopyButton.svelte";
	import PageLoading from "$lib/components/loading/PageLoading.svelte";
	import { convertToReadableTime }  from "$lib/parsers/time";
	import { seperateMemberOfGroups } from "$lib/parsers/members";
	import { readableOU } from "$lib/parsers/ou";
	import { onMount } from "svelte";
	import { goto } from "$app/navigation";
    import { browser } from "$app/environment";

    let ou: string = $state("")
    let section: string = $state("")
    const attributes = ["cn","username","phone","department","dn","urid","netid","relationship","pwdlastset","email","description","title", "memberof"]
    let userPromise: Promise<any> = $state(new Promise(() => {}));
    let timer: number = $state(3)
    
    function startTimer() {
        setInterval(() => {
            timer -= 1;
        }, 1000);
    }

    async function fetchUser(): Promise<any> {
        ou = page.url.searchParams.get("dn") ?? ""
        if (ou == "") goto("/search"); 
        const res = await fetch(`http://localhost:8000/api/user?dn=${encodeURIComponent(ou)}&attributes=${attributes.join(",")}`);
        if (!res.ok && browser) {
            setTimeout(() => {
                goto("/search")
            }, 3000);

            startTimer()

            throw new Error(`Failed to find User with ${ou}`);
        }
        return res.json();
    }

    onMount(async () => {
        section = page.url.searchParams.get("section") ?? "PROFILE"
        userPromise = fetchUser()
    })




</script>

<!-- * Waits for the response for the server to render content -->
 <main>
    {#await userPromise}
        <PageLoading/>
    {:then user} 
        {@render Content(user)}
    {:catch error}
        <h1>{error.message} Redirecting in {timer}s</h1>
    {/await}
</main>

{#snippet Content(u: any)}
    <section>
        <div>
            <CopyAllButton copyTemplate={"test"} />
            <CopyButton value={u.cn[0]} category="title"/>
            <CopyButton label="USERNAME" value={u.username[0] ?? "NA"} category="list-item"/>
            <CopyButton label="NETID" value={u.netid[0] ?? "NA"} category="list-item"/>
            <CopyButton label="URID" value={u.urid[0] ?? "NA"} category="list-item"/>
            <CopyButton label="PHONE" value={u.phone[0] ?? "NA"} category="list-item"/>
            <CopyButton label="DEPARTMENT" value={u.department[0] ?? "NA"} category="list-item"/>
            <CopyButton label="RELATIONSHIP" value={u.relationship[0] ?? "NA"} category="list-item"/>
            <CopyButton label="LASTPASSWORDSET" value={convertToReadableTime(u.pwdlastset[0] ?? "0")} category="list-item"/>    
            <CopyButton label="OU" value={readableOU(u.dn[0]) ?? "NA"} category="list-item"/>
            <li>Memberships</li>
                {#each seperateMemberOfGroups(u.memberof) as group }
                    <li>{group}</li>
                {/each}
        </div>
    </section>
{/snippet}

<style>
    main {
        width: 80%;
        color: var(--text);
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    div {
        width: 500px;
        height: fit-content;
        padding: 2rem;
        box-shadow: 10px 0 10px 5px rgba(0,0,0,0.3);
        border-radius: 0px 0px 40px 40px;
        background: black;
    }

    section {
        display: flex;
        position: relative;
        flex-direction: column;
        align-items: center;
        overflow: hidden;
        width: 500px;
        height: 500px;
        border-radius: 40px 40px 40px 40px;
        background: white;
    }
</style>
