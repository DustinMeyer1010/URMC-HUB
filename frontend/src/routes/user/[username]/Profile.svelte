<script lang="ts">
	import type { UserFullInfo } from "@t/user";
	import { slide } from "svelte/transition";
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

<section in:slide={{delay: delay, duration: duration}} out:slide={{duration: duration}}>
    {#if user.ou.toLocaleLowerCase().includes("disabled")}
        <h1 class="disabled"><img src={disabledIcon} alt=""/>Disabled Account</h1>
    {/if}
    <ul>
        {#each Object.entries(user) as key}
            {#if !notWantedAttributes.includes(key[0].toLocaleLowerCase())}
            {console.log(key[0])}
                <li>
                    <b>{key[0].toUpperCase()}: </b> 
                    {key[1] ? key[1] : "NA"}
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

</style>