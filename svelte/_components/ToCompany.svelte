<script>
  // @ts-nocheck
  import InputBox from "./InputBox.svelte";
  import SubmitButton from "./SubmitButton.svelte";
  import {UserCreateCompany} from '../jsApi.GEN.js';
  import {notifier} from './notifier';
  import { onMount } from "svelte";

  export let user;

  onMount(() => {
    console.log('onMount.toCompany ', user);
  })

  const DEFAULT = 'DEFAULT', CREATE = 'CREATE', JOIN = 'JOIN';
  let mode = DEFAULT;

  let tenantCode = '', companyName = '', headTitle = '';
  let isSubmitCreateCompany = false;
  async function userCreateCompany() {
    isSubmitCreateCompany = true;
    await UserCreateCompany({tenantCode, companyName, headTitle},
      function (o) {
        if (o.error) {
          isSubmitCreateCompany = false;
          notifier.showError(o.error);
          console.log(o.error);
          return;
        }
        isSubmitCreateCompany = false;
        console.log(o);
        notifier.showSuccess('Company created successfully');
      }
    );
  }

  let invitationUrl = '', isSubmitJoinCompany = false;
  function userJoinCompany() {
    if (!invitationUrl) return notifier.showWarning('Invitation URL is required');
    window.location.href = invitationUrl;
  }
</script>

<div class="to_company_container">
  {#if mode===DEFAULT}
    <div class="join_create_company">
      <button on:click={() => mode=JOIN}>
        <img src="/assets/icons/join-company.svg" alt="" />
        <span>Join Company</span>
      </button>
      <button on:click={() => mode=CREATE}>
        <img src="/assets/icons/create-company.svg" alt="" />
        <span>Create Company</span>
      </button>
    </div>
  {/if}

  {#if mode === JOIN}
    <div style="width: 400px;">
      <InputBox id="invitationUrl" label="Invitation URL" bind:value={tenantCode} type="text" />
      <br /><br />
      <SubmitButton on:click={userJoinCompany} isSubmitted={isSubmitJoinCompany} />
    </div>
  {/if}

  {#if mode===CREATE}
  <div style="width: 400px;">
    <InputBox id="tenantCode" label="Tenant Code" bind:value={tenantCode} type="text" />
    <br /><br />
    <InputBox id="companyName" label="Company Name" bind:value={companyName} type="text" />
    <br /><br />
    <InputBox id="headTitle" label="Head Title" bind:value={headTitle} type="text" />
    <br /><br />
    <SubmitButton on:click={userCreateCompany} isSubmitted={isSubmitCreateCompany} />
  </div>
  {/if}
</div>

<style>
  .to_company_container {
    display: flex;
    width: 100%;
    height: fit-content;
  }

  .to_company_container .join_create_company {
    display: grid;
    grid-template-columns: 1fr 1fr;
    align-items: stretch;
    gap: 20px;
    width: 100%;
  }

  .to_company_container .join_create_company button {
    width: 100%;
    height: 100%;
    background-color: #FFF;
    border-radius: 20px;
    border: 1px solid var(--gray-003);
    cursor: pointer;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 30px;
    padding: 20px;
    font-size: var(--font-xl);
    color: var(--gray-007);
    font-weight: 700;
  }

  .to_company_container .join_create_company button:hover {
    border: 1px solid var(--blue-005);
    color: var(--blue-005);
  }

  .to_company_container .join_create_company button img {
    width: 360px;
    height: auto;
  }
</style>