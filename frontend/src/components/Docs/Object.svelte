<script lang="ts">


    let {objects} : {objects: Attributes[]} = $props()

</script>

<h1>Object</h1>
<hr>
<ul>
    {#each objects as object }
        <li>
            <b>{object.name}</b><sup>{object.type}</sup>
            <p>{object.description}</p>

        </li>
        <hr>
    {/each}
</ul>




<style>

    hr {
        width: 100%;
        height: 2px;
        margin: 5px;
        border: none;
        background: var(--color-surface-lighter)
    }

    h1 {
        font-size: 18px;
        margin-top: 10px;
        margin-bottom: 0;
    }

    b {
        padding: 0;
        margin: 0;
    }

    ul {
        padding: 0;
        margin: 0;
        list-style: none;
    }

    li {
        padding-left: 10px;
    }
    
    p {
        margin: 0;
        padding: 5px;
        font-size: 12px;
    }

    sup {
        font-weight: 100;
        font-size: 10px;
        margin-left: 2px;
    }

</style>