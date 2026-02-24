<script lang="ts">
  import { theme } from "$lib/stores/theme";
  import { get } from "svelte/store";

  function toggleTheme() {
    const current = get(theme);
    theme.set(current === "light" ? "dark" : "light");
  }
</script>

<button on:click={toggleTheme}>
  Toggle to { $theme === "light" ? "dark" : "light" } mode
</button>


<style>
  button {
    background: none;
    border: none;
    padding: 1rem;
    box-sizing: border-box;
    color: var(--color-text);
  }

  button:hover {
    color: var(--color-text-hover)
  }
</style>