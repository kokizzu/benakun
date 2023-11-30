<script>
  // @ts-nocheck
  import InputBox from "./InputBox.svelte";
  import SubmitButton from "./SubmitButton.svelte";
  import {UserCreateCompany} from '../jsApi.GEN.js';
  import {notifier} from './notifier';
  import { onMount } from "svelte";
  import BackButton from "./BackButton.svelte";

  export let user;

  onMount(() => {
    console.log('onMount.toCompany ', user);
  })

  const DEFAULT = 'DEFAULT', CREATE = 'Create Company', JOIN = 'Join Company';
  let mode = DEFAULT;

  let tenantCode = '', companyName = '', headTitle = '';
  let isSubmitCreateCompany = false;
  async function userCreateCompany() {
    isSubmitCreateCompany = true;
    if (!tenantCode || !companyName || !headTitle) {
      isSubmitCreateCompany = false;
      notifier.showWarning('All fields are required');
      return;
    }
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
  <div class="create_company">
    <section class="container">
      <header>
        <BackButton on:click={() => mode=DEFAULT} />
        <h2>{JOIN}</h2>
      </header>
      <div>
        <InputBox id="invitationUrl" label="Invitation URL" bind:value={tenantCode} type="text" placeholder="https://benakun.com/join?companyId="/>
        <SubmitButton on:click={userJoinCompany} isSubmitted={isSubmitJoinCompany} isFullWidth />
      </div>
    </section>
  </div>
  {/if}

  {#if mode===CREATE}
  <div class="create_company">
    <section class="container">
      <header>
        <BackButton on:click={() => mode=DEFAULT} />
        <h2>{CREATE}</h2>
      </header>
      <div>
        <InputBox id="tenantCode" label="Tenant Code" bind:value={tenantCode} type="text" />
        <InputBox id="companyName" label="Company Name" bind:value={companyName} type="text" />
        <InputBox id="headTitle" label="Head Title" bind:value={headTitle} type="text" />
        <SubmitButton on:click={userCreateCompany} isSubmitted={isSubmitCreateCompany} isFullWidth />
      </div>
    </section>
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
    width: 600px;
  }

  .to_company_container .join_create_company > button {
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

  .to_company_container .join_create_company > button:hover {
    border: 1px solid var(--blue-005);
    color: var(--blue-005);
  }

  .to_company_container .join_create_company > button img {
    width: 100%;
    height: auto;
  }

  .to_company_container .create_company {
    width: 100%;
    height: fit-content;
    display: flex;
    justify-content: center;
  }

  .to_company_container .create_company .container {
    display: flex;
    flex-direction: column;
    gap: 20px;
    height: fit-content;
    width: 400px;
    background-color: #FFF;
    border-radius: 10px;
    border: 1px solid var(--gray-002);    
    padding: 20px;
  }

  .to_company_container .create_company .container header {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 15px;
  }

  .to_company_container .create_company .container header h2 {
    margin: 0;
  }

  .to_company_container .create_company .container div {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }
</style>