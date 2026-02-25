<script lang="ts">
	import type { User } from "$lib/types/user";
    import  disabledIcon  from '$lib/assets/disabled-color-disabled.png'
	import { Copy } from "$lib/types/copy";
	import { onMount } from "svelte";
	import { redirect } from "@sveltejs/kit";
    import expired_password from "$lib/assets/expired-password.png"
    const notWantedAttributes = ["member_of", "name"]



    let copyState: Copy.State = $state(Copy.EMPTY_COPY_STATE)

    let user: User.PageInfo | null = $state(null)

    let passwordExpired: boolean = $derived.by(() => {
        if (user) {
            const oneYearAgo = new Date();
            oneYearAgo.setFullYear(oneYearAgo.getFullYear() - 1);
            return new Date(user.last_password_set) < oneYearAgo
        }
        return false
    })
    

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
                        <button onclick={() => Copy.ToClipboard(key[1] ? key[1] as string : "NA", copyState)}>
                            {#if key[0].toUpperCase() === "LAST_PASSWORD_SET" && passwordExpired}
                                <img id="expired" src={expired_password} alt="">
                                <span id="expired"><b>{key[0].toUpperCase()}: </b> </span>
                            {:else}
                                <span><b>{key[0].toUpperCase()}: </b> </span>
                            {/if}
                            
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


    span#expired {
        color: var(--color-ad-disabled)
    }


    img#expired {
        position: absolute;
        top: -15px;
        left: -15px;
        width: 20px;
    }

    button {
        position: relative;
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