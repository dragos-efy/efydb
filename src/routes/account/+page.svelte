<section>
    {#if token}
    <div>
        {userInfo.name}
    </div>
    {:else}
    <div id="signUp">
        <input type="text" bind:value={username} placeholder="Username">
        <input type="text" bind:value={password} placeholder="Password">
        <input type="text" bind:value={bio} placeholder="Bio">
        <div id="sign-buttons">
            <button on:click={signIn}>Sign In</button>
            <button on:click={signUp}>Sign Up</button>
        </div>
    </div>
    {/if}
</section>

<script type="ts">
	import fetchJson from "$lib/fetchjs";
	import { getToken, setToken } from "$lib/token";
	import { onMount } from "svelte";

    let token = getToken();
    let userInfo: any;

    let username: string;
    let password: string;
    let bio: string;

    onMount(async () => {
        if (token) {
            userInfo = await fetchJson("/users/account", {})
        }
    })

    const signUp = async () => {
        if (!password || !username) {
            alert("Please fill out all the info!");
            return
        }

        const response = await fetchJson("/users/register", {
            method: "POST",
            body: JSON.stringify({
                name: username,
                password: password,
                bio: bio,
            })
        })
        setToken(response.token);
        token = response.token;
        console.log(response);
    }

    const signIn = async () => {
        if (!password || !username) {
            alert("Please fill out all the info!");
            return
        }

        const response = await fetchJson("/users/login", {
            method: "POST"
        })
        setToken(response.token);
        token = response.token;
        console.log(response);
    }
</script>

<style>
    #sign-buttons {
        display: flex;
        justify-content: end;
    }

    #sign-buttons button:last-child {
        margin-left: 10rem;
    }
</style>