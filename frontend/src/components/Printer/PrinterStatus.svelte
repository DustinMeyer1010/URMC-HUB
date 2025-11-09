<script lang="ts">

    import onlinePrinterDark from "$lib/assets/online-printer-color-accent-dark.png"
    import offlinePrinterDark from "$lib/assets/offline-printer-color-error-dark.png"
    import IconLinkDark from "$lib/assets/link-color-text-dark.png"

    let {
        isOnline,
        eRecordPrinter,
        ip
    } : {
        isOnline: boolean
        eRecordPrinter: boolean
        ip: string
    } = $props()


</script>
    
    
    
    
{#if isOnline}
    <div>
        <h2>Online</h2>
        {#if !eRecordPrinter}
            <a href={`http://${ip.replace(/_.*/, "")}`} title="Printer Settings" target="_blank"><img src={onlinePrinterDark} alt=""></a>
        {:else}
            <img src={onlinePrinterDark} alt="">
        {/if}


    </div>
        {:else}
        <h2>Offline<img src={offlinePrinterDark} alt=""></h2>

{/if}



<style>


    div {
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 5px;
    }

    a {
        text-decoration: none;
        color: var(--color-text);
        font-weight: bold;
    }

    h2 {
        display: flex;
        gap: 10px;
        font-size: 15px;
    }

    img {
        width: 15px;
    }
    
</style>