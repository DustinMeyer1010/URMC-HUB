<script lang="ts">
	import PageLoading from "@components/Loading-Animations/PageLoading.svelte";
	import type { ComputerPageInfo } from "@t/computer";
    import ComputerOnlineIcon from "$lib/assets/computer-online.png"
    import ComputerOfflineIcon from "$lib/assets/computer-offline.png"
	import { onMount } from "svelte";
    import DisabledIcon from "$lib/assets/disabled-computer.png"
	import Ping from "@components/Computer/Ping.svelte";

    let { data } : {data: {name: string}} = $props();

    let computer: ComputerPageInfo | null = $state(null)
    let disabled: boolean | undefined = $derived.by(() => {
        return computer?.computer_info.ou.toUpperCase().includes("DISABLED")
    })


    onMount(async () => {
        let res = await fetch(`http://localhost:8000/api/computer/${data.name}/info`)

        computer = await res.json()
    })
</script>

<section>
{#if computer}
    <div class="header">
        <h1>{computer?.computer_info.name}</h1>
        <sup>
            {#if computer.is_online}
                <span>ONLINE <img src={ComputerOnlineIcon} alt=""></span>
            {:else}
                <span title="Could be offline due to user being remote">OFFLINE <img src={ComputerOfflineIcon} alt=""></span>
            {/if}
        </sup>
    </div>

    <div class="main">
            {#if disabled} 
                <span class="disabled"><img src={DisabledIcon} alt="">Computer Disabled </span>
            {/if}
        <ul>
            <li><b>OU: </b> {computer.computer_info.ou}</li>
            <li><b>OS:</b> {computer.computer_info.operating_system}</li>
        </ul>
    </div>
    <Ping isOnline={computer.is_online} pingResults={computer.ping_results}/>
{:else}

<PageLoading/>

{/if}

</section>


<style>

    span.disabled {
        color: var(--color-ad-disabled)
    }

    ul {
        list-style: none;
        display: flex;
        flex-direction: column;
        text-align: left;
        gap: 1rem;
        font-size: 15px;
        word-break: break-all;
    }

    section {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }
    

    div.main {
        background: var(--color-surface);
        width: 75%;
        justify-self: center;
        padding: 1rem;
        border-radius: 20px;
    }


    div.header {
        position: relative;
        display: flex;
        justify-content: center;
        padding-top: 30px;
    }

    span {
        display: flex;
        align-items: center;
        gap: 5px;
        font-weight: bold;
        font-size: 12px;
    }

    div.header sup {
        padding-top: 5px;
    }

    div.header h1 {
        font-size: 30px;
    }

    img {
        width: 20px;
    }

    @media (max-width: 800px) {
        div.main ul {
            padding: 1rem;
            margin: 0;
        }
    }

</style>
