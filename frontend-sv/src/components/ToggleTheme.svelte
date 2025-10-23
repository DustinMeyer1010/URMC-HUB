<script lang="ts">
  import { theme } from "$lib/theme";
  import { get } from "svelte/store";

  function toggleTheme() {
    const current = get(theme);
    theme.set(current === "light" ? "dark" : "light");
  }
</script>

<button on:click={toggleTheme}>
  Toggle to { $theme === "light" ? "dark" : "light" } mode
</button>