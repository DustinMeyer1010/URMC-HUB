<script lang="ts">
	import type { User } from "@t/user";
    import  disabledIcon  from '$lib/assets/disabled-color-disabled.png'
	import { copyToClip, type CopyState } from "$lib/helper/copy.svelte";
	import { onMount } from "svelte";
	import { redirect } from "@sveltejs/kit";
    const notWantedAttributes = ["member_of", "name"]



    let copyState: CopyState = $state({
        copied: "",
        timeout: null
    })

    let user: User.PageInfo | null = $state(null)
    

    let {
        username,
    } : {
        username: string
    } = $props()

    onMount(async () => {
        await GetUserFullInfo()
    })

    const GetUserFullInfo = async () => {

        const response: Response = await fetch(`http://localhost:8000/api/user/${username}`);

        if (!response.ok) {
            throw redirect(301, "/search")
        }

        user = await response.json();
    }


</script>

<section >
    {#if user != null}
        {#if user.ou.toLocaleLowerCase().includes("disabled")}
            <h1 class="disabled"><img src={disabledIcon} alt=""/>Disabled Account</h1>
        {/if}
        <ul>
            {#each Object.entries(user) as key}
                {#if !notWantedAttributes.includes(key[0].toLocaleLowerCase())}
                    <li>
                        <button onclick={() => copyToClip(key[1] ? key[1] as string : "NA", copyState)}>
                            <span><b>{key[0].toUpperCase()}: </b> </span>
                            <span>
                                {#if copyState.copied === key[1] && key[1] !== ""}
                                    Copied
                                {:else}
                                    {key[1] ? key[1] : "NA"}
                                {/if}
                            </span>
                        </button>
                    </li>
                {/if}

            {/each}

        </ul>
    {/if}
</section>



<style>
    section {
        width: 100%;
        padding: 2rem 1rem;
        font-size: 15px;
        border-radius: 10px;
        background: var(--color-surface);
        position: relative;
        animation: slideIn 0.5s forwards;
        opacity: 0;
        animation-delay: 50ms;
    }

    button {
        background: none;
        border: none;
        color: var(--color-text);
        font-size: 15px;
    }

    ul {
        display: grid;
        grid-template-columns: 1fr 1fr 1fr;
        gap: 3rem;
        list-style: none;
    }

    li {
        word-break: break-all;
    }

    h1.disabled {
        position: absolute;
        top: 0.3rem;
        left: 1rem;
        font-size: 12px;
        color: var(--color-ad-disabled);
    }

    h1.disabled img {
        margin-right: 2px;
        transform: translateY(2px);
        width: 20px;
    }

    @keyframes slideIn {
        from {
            transform: translateY(100%);
            opacity: 0;
        }
        to {
            transform: translateY(0);
            opacity: 1;
        }
    }

    @media (max-width: 1200px) {
        ul {
            grid-template-columns: 1fr 1fr;
        }
    }

    @media (max-width: 900px) {
        section {
            width: 95%;
        }
        ul {
            grid-template-columns: 1fr;
        }
    }

    @media (max-width: 600px) {
        section {
            width: 95%;
        }

        li {
            display: flex;
            flex-direction: column;
            flex-wrap: wrap-reverse;
            gap: 1rem;
        }
    }

</style>