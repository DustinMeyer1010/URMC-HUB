<script lang="ts">
	import Nav from '../Nav.svelte';
	import  {type Section, Sections } from '../section.js';

    let {
        data
    } : {
        data: {name: string}
    } = $props()

    let section: Section = $state("INFO")

    const swapSection = (newSection: Section) => {
        section = newSection
    }
</script>

<h1>{data.name}</h1>
<Nav sections={Sections} {swapSection} />


