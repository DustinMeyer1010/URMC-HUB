<script lang="ts">
    import type { GroupPageInfo } from "@t/group";
    let {
        group
    } : {
        group: GroupPageInfo
    } = $props();
</script>

<section>
    {group.name}
    {group.ou}
    {group.description}
    {group.information}

</section>


<style>

</style>