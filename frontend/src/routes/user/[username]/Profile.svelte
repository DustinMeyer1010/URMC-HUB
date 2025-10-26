<script lang="ts">
	import type { UserFullInfo } from "@t/user";
	import { fly } from "svelte/transition";
    import  disabledIcon  from '$lib/assets/disabled-color-disabled.png'

    const duration = 200;
    const delay = 250;
    const notWantedAttributes = ["member_of", "name"]

    

    let {
        user 
    } : {
        user: UserFullInfo
    } = $props()



</script>

<section in:fly={{delay: delay, duration: duration, x: -200, y: 200}} out:fly={{duration: duration, x: 200, y: 200}}>
    {#if user.ou.toLocaleLowerCase().includes("disabled")}
        <h1 class="disabled"><img src={disabledIcon} alt=""/>Disabled Account</h1>
    {/if}
    <ul>
        {#each Object.entries(user) as key}
            {#if !notWantedAttributes.includes(key[0].toLocaleLowerCase())}
                <li>
                    <span><b>{key[0].toUpperCase()}: </b> </span>
                    <span>{key[1] ? key[1] : "NA"}</span>
                </li>
            {/if}

        {/each}

    </ul>
</section>



<style>
    section {
        width: 100%;
        padding: 2rem 1rem;
        font-size: 15px;
        border-radius: 10px;
        background: var(--color-surface);
        position: relative;
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