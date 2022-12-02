<section>
    {#if !loaded}
    <div class="spin"></div>
    {:else if token}
    <div>
        {userInfo?.name}
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

    let loaded = false;
    let token: string | null;
    let userInfo: any;

    let username: string;
    let password: string;
    let bio: string;

    onMount(async () => {
        token = getToken();
        if (token) {
            userInfo = await fetchJson("/users/account", {});
        }
        loaded = true;
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

        if (response.message) {
            alert(response.message);
            return;
        }

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

        if (response.message) {
            alert(response.message);
            return;
        }

        setToken(response.token);
        token = response.token;
        userInfo = response.
        console.log(response);
    }
</script>

<style>
    section {
        width: 100%;
        min-height: 70vh;
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
    }

    #sign-buttons {
        display: flex;
        justify-content: end;
    }

    #sign-buttons button:last-child {
        margin-left: 10rem;
    }

    .spin {
        display: inline-block;
        width: 50px;
        height: 50px;
        border: 3px solid rgba(255, 255, 255, .3);
        border-radius: 50%;
        border-top-color: var(--efy_color2);
        animation: spin 1s ease-in-out infinite;
        -webkit-animation: spin 1s ease-in-out infinite;
      }
      @keyframes spin {
        to {
          -webkit-transform: rotate(360deg);
        }
      }
      @-webkit-keyframes spin {
        to {
          -webkit-transform: rotate(360deg);
        }
    }
</style>