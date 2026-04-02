<script lang="ts">
	import { browser } from "$app/environment";
	import { goto } from "$app/navigation";
	import CopyAllButton from "$lib/components/buttons/CopyAllButton.svelte";
	import CopyButton from "$lib/components/buttons/CopyButton.svelte";
	import PageLoading from "$lib/components/loading/PageLoading.svelte";
	import { Icons } from "$lib/managers/icons";
	import { readableDN } from "$lib/parsers/ou";
	import { convertToReadableTime } from "$lib/parsers/time";
	import { onMount } from "svelte";


    let {userDN} : {userDN: string} = $props()

    let userInfo: Promise<any> = $state(new Promise(() => {}))
    let timer: number = $state(3)
    const attributes: string[] = ["cn","username","phone","department","dn","urid","netid","relationship","pwdlastset","email","description","title"]
    
    function startTimer() {
        setInterval(() => {
            timer -= 1;
        }, 1000);
    }

    async function fetchUser(): Promise<any> {

        const res = await fetch(`http://localhost:8000/api/user?dn=${encodeURIComponent(userDN)}&attributes=${attributes.join(",")}`);
        if (!res.ok && browser) {
            setTimeout(() => {
                goto("/normal/search")
            }, 3000);

            startTimer()

            throw new Error(`Failed to find User with ${userDN}`);
        }
        return res.json();
    }

    onMount(async () => {
        userInfo = fetchUser()
    })

</script>

<section>
    {#await userInfo}
        <PageLoading/>
    {:then user} 
        {@render Content(user)}
    {:catch error}
        <h1>{error.message} Redirecting in {timer}s</h1>
    {/await}
</section>



{#snippet Content(u: any)}
    <section>
        <div>
            <CopyAllButton icon={Icons.COPY} copyTemplate={"test"} />
            <CopyButton value={u.cn.length > 0 ? u.cn.join() : "NA"} category="title"/>
            <CopyButton label="USERNAME" value={u.username.length > 0 ? u.username.join() : "NA"} category="list-item"/>
            <CopyButton label="NETID" value={u.netid.length > 0 ? u.netid.join() : "NA"} category="list-item"/>
            <CopyButton label="URID" value={u.urid.length > 0 ? u.urid.join() : "NA"} category="list-item"/>
            <CopyButton label="PHONE" value={u.phone.length > 0 ? u.phone.join() : "NA"} category="list-item"/>
            <CopyButton label="DEPARTMENT" value={u.department.length > 0 ?  u.department.join() : "NA"} category="list-item"/>
            <CopyButton label="RELATIONSHIP" value={u.relationship.length > 0 ? u.relationship.join(" | ") : "NA"} category="list-item"/>
            <CopyButton label="LASTPASSWORDSET" value={convertToReadableTime(u.pwdlastset[0] ?? "0")} category="list-item"/>    
            <CopyButton label="OU" value={readableDN(u.dn[0]) ?? "NA"} category="list-item"/>
        </div>
    </section>
{/snippet}