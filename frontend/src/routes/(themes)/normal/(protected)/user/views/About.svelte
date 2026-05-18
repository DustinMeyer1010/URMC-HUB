<script lang="ts">
	import { browser } from "$app/environment";
	import { goto } from "$app/navigation";
	import CopyAllButton from "$lib/components/buttons/CopyAllButton.svelte";
	import CopyButton from "$lib/components/buttons/CopyButton.svelte";
	import Card from "$lib/components/cards/Card.svelte";
	import PageLoading from "$lib/components/loading/PageLoading.svelte";
	import { Icons } from "$lib/managers/icons";
	import { readableDN } from "$lib/parsers/ou";
	import { convertToReadableTime } from "$lib/parsers/time";
	import { onMount } from "svelte";


    let {userDN} : {userDN: string} = $props()

    let userInfo: Promise<any> = $state(new Promise(() => {}))
    let timer: number = $state(3)
    const attributes: string[] = ["cn","username","phone","department","dn","urid","netid","relationship","pwdlastset","email","description","title", "memberof"]

    function checkPassword(passwordExpiredTime: string): boolean {
        const val = BigInt(passwordExpiredTime);
        if (val === 0n) return true;
        const windowsEpochOffset = 116444736000000000n;
        const millisecondsSinceEpoch = (val - windowsEpochOffset) / 10000n;

        const passwordDate = new Date(Number(millisecondsSinceEpoch));
        const oneYearAgo = new Date();
        oneYearAgo.setFullYear(oneYearAgo.getFullYear() - 1);

        return passwordDate < oneYearAgo;
    }

    function checkIdleAccount(groups: string): boolean {
        return groups.includes("IDM_IdleAccounts_URMC")
    }
    
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
        <Card>
            {#snippet header()}
                <CopyAllButton icon={Icons.COPY} copyTemplate={"test"} />
                {#if checkPassword(u.pwdlastset[0]) }
                    <span>password expired</span>
                {/if}
                {#if checkIdleAccount(u.memberof.join())}
                    <span>Has IDLE Account</span>
                {/if}
            {/snippet}

            {#snippet body()}
                <div>
                    <CopyButton value={u.cn.length > 0 ? u.cn.join() : "NA"} category="title"/>
                    <CopyButton label="USERNAME" value={u.username.length > 0 ? u.username.join() : "NA"} category="list-item"/>
                    <CopyButton label="NETID" value={u.netid.length > 0 ? u.netid.join() : "NA"} category="list-item"/>
                    <CopyButton label="URID" value={u.urid.length > 0 ? u.urid.join() : "NA"} category="list-item"/>
                    <CopyButton label="EMAIL" value={u.email.length > 0 ? u.email.join() : "NA"} category="list-item"/>
                    <CopyButton label="PHONE" value={u.phone.length > 0 ? u.phone.join() : "NA"} category="list-item"/>
                    <CopyButton label="DEPARTMENT" value={u.department.length > 0 ?  u.department.join() : "NA"} category="list-item"/>
                    <CopyButton label="RELATIONSHIP" value={u.relationship.length > 0 ? u.relationship.join(" | ") : "NA"} category="list-item"/>
                    <CopyButton label="LAST PASSWORD SET" value={convertToReadableTime(u.pwdlastset[0] ?? "0")} category="list-item"/>    
                    <CopyButton label="OU" value={readableDN(u.dn[0]) ?? "NA"} category="list-item"/>
                </div>
            {/snippet}
                
        
        </Card>
    </section>
{/snippet}


<style> 

    section {
        height: 100%;
    }

    div {
        display: flex;
        flex-direction: column;
        gap: 1rem;
        background: var(--color-surface);
        padding: 10px;
        height: 100%;
        border-radius: 10px;
    }

    @media (max-width: 700px) {
        div {
            gap: 0.3rem;
        }
    }

</style>