<script lang="ts">

    let {
        header,
        information
    } : {
        header: string
        information: string
     } = $props()

</script>

<div>
    <h2>{header}</h2>
        <p>
            {information}
        </p>

</div>

<style>
    div {
        width: 100%;
        overflow: hidden;
        border-radius: 6px;
        line-height: 1.5rem;
    }

     h2 {
        width: 100%;
        font-size: 18px;
        padding: 5px;
        margin: 0;
        box-sizing: border-box;
        padding: 5px;
     }

     p {
        font-size: 14px;
     }

</style>