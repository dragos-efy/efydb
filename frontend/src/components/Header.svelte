<script lang="ts">
	import { getRole, getToken } from "$lib/token";
	import { onMount } from "svelte";

    let token: string | null;
    let role: number | null;

    onMount(() => {
        token = getToken();
        role = getRole();
    })
</script>

<header>
    <a href="/">EfyDb</a>
    <nav>
        <ul>
            <li><a href="/">Themes</a></li>
            <li><a href="/account">Account</a></li>
            {#if role && role > 0}
            <li><a href="/users">Users</a></li>
            {/if}
            {#if token}
            <li><a href="/new">New</a></li>
            {/if}
        </ul>
    </nav>
</header>

<style>
    header {
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
</style>