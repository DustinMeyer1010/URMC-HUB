<script lang="ts">

    let {
        isOnline,
        pingResults
    } : {
        isOnline: boolean
        pingResults: string
    } = $props()

</script>

<div>
{#if isOnline}

        <h1>Ping Results</h1>
        <span>{pingResults}</span>
{:else}
    <p>Computer is not responding to pings.</p>
    <h2>Causes: </h2>
    <ul>
        <li>Customer is remote</li>
        <li>Computer is disbaled</li>
        <li>Computer is powered off</li>
    </ul>
{/if}
</div>

<style>


    span{
        width: 75%;
        align-self: center;
    }

    p {
        padding: 0;

        margin-top: 10px;
    }


    h1,h2 {
        padding: 10px;
        margin: 0;
    }

    h1 {
        text-align: center;
    }



    div {
        display: flex;
        flex-direction: column;
    }

    ul {
        list-style-type: disc;
        font-size: 12px;
    }

    li {
        margin-top: 3px;
    }



</style>