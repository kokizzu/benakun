<script>
  import {
    UserCreateCompany,
    GuestForgotPassword,
    GuestLogin,
    GuestRegister,
    GuestResendVerificationEmail
  } from './jsApi.GEN.js';
  import { onMount, tick } from 'svelte';
  import { notifier } from './_components/xNotifier.js';
  import InputBox from './_components/InputBox.svelte';
  import SubmitButton from './_components/SubmitButton.svelte';
  import MainLayout from './_layouts/mainLayout.svelte';

  /** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/user.js').User} User */
  /** @typedef {import('./_components/types/organization.js').Org} Org */

  let user      = /** @type User */   ({/* user */});
  let segments  = /** @type Access */ ({/* segments */});
  let myCompany = /** @type Org */    ({/* myCompany */});

  console.log('User:', user);
  console.log('Segments:', segments);

  const google = '#{google}';
  const title = '#{title}';

  console.log('Google OAuth link:', google);
  console.log('Title:', title);

  /**
   * @description Get cookie
   * @param name {string}
   * @returns {string}
   */
  function getCookie(name) {
    const match = document.cookie.match(new RegExp('(^| )' + name + '=([^;]+)'));
    if (match) return match[2];
  }

  // Login payload
  let email       = '';
  let password    = '';
  let confirmPass = '';

  const LOGIN = 'LOGIN';
  const REGISTER = 'REGISTER';
  const RESEND_VERIFICATION_EMAIL = 'RESEND_VERIFICATION_EMAIL';
  const FORGOT_PASSWORD = 'FORGOT_PASSWORD';
  const USER = '';
  let mode = LOGIN, isSubmitted = false;

  async function onHashChange() {
    const auth = getCookie('akun');
    if (auth && user && !auth.startsWith('TEMP__')) {
      location.hash = '';
      mode = USER;
      return;
    }
    let hash = location.hash || '';
    if (hash[0] === '#') hash = hash.substring(1);

    if (hash === LOGIN) mode = LOGIN;
    else if (hash === REGISTER) mode = REGISTER;
    else if (hash === RESEND_VERIFICATION_EMAIL) mode = RESEND_VERIFICATION_EMAIL;
    else if (hash === FORGOT_PASSWORD) mode = FORGOT_PASSWORD;
    else location.hash = LOGIN;
    await tick();
  }

  onMount(() => onHashChange());

  async function guestRegister() {
    isSubmitted = true;
    if (!email) {
      isSubmitted = false;
      return notifier.showWarning('Email is required');
    }
    if (password.length < 12) {
      isSubmitted = false;
      return notifier.showWarning('Password must be at least 12 characters');
    }
    if (password !== confirmPass) {
      isSubmitted = false;
      return notifier.showWarning('Passwords do not match');
    }
    const i = { email, password };
    await GuestRegister(i, /** @type {import('./jsApi.GEN.js').GuestRegisterCallback}*/ async function (/** @type {any} */ o) {
      if (o.error) {
        console.log(o);
        isSubmitted = false;
        return notifier.showError(o.error);
      }
      isSubmitted = false;
      notifier.showSuccess('Registered successfully, a registration verification has been sent to your email');
      mode = LOGIN;
      password = '';
      await tick();
    });
  }

  async function guestLogin() {
    isSubmitted = true;
    if (!email) {
      isSubmitted = false;
      return notifier.showWarning('Email is required');
    }
    if (password.length < 12) {
      isSubmitted = false;
      return notifier.showWarning('Password must be at least 12 characters');
    }
    const i = { email, password };
    await GuestLogin(i, /** @type {import('./jsApi.GEN.js').GuestLoginCallback}*/ function (/** @type {any} */ o) {
      if (o.error) {
        isSubmitted = false;
        notifier.showError(o.error);
        return;
      }

      isSubmitted = false;
      notifier.showSuccess('Login success');
      user = o.user;
      segments = o.segments;
      window.document.location = '/';
    });
  }

  async function guestResendVerificationEmail() {
    isSubmitted = true;
    if (!email) {
      isSubmitted = false;
      notifier.showWarning('Email is required');
      return;
    }
    const i = { email };
    await GuestResendVerificationEmail(i, /** @type {import('./jsApi.GEN.js').GuestResendVerificationEmailCallback} */ function (/** @type {any} */ o) {
      if (o.error) {
        console.log(o);
        isSubmitted = false;
        notifier.showError(o.error);
        return;
      }
      isSubmitted = false;
      onHashChange();
      notifier.showInfo('An email verification link has been sent to your email');
    });
  }

  async function guestForgotPassword() {
    isSubmitted = true;
    if (!email) {
      isSubmitted = false;
      notifier.showWarning('Email is required');
      return;
    }
    const i = { email };
    await GuestForgotPassword(i, /** @type {import('./jsApi.GEN.js').GuestForgotPasswordCallback}*/ function (/** @type {any} */ o) {
      if (o.error) {
        console.log(o);
        isSubmitted = false;
        notifier.showError(o.error);
        return;
      }
      onHashChange();
      notifier.showInfo('A reset password link has been sent to your email');
    });
  }

  // Payload for create company
  let tenantCode  = '';
  let companyName = '';
  let headTitle   = '';

  onMount(() => {
    console.log('My Company:', myCompany);
    if (myCompany) {
      tenantCode = myCompany.tenantCode;
      companyName = myCompany.name;
      headTitle = myCompany.headTitle;
    }
  })

  let isCreatingCompany = false;
  async function SubmitCreateCompany() {
    if (!tenantCode || !companyName || !headTitle) {
      notifier.showWarning('all fields are required');
      return;
    }

    isCreatingCompany = true;

    /** @type {import('./jsApi.GEN.js').UserCreateCompanyIn} */ //@ts-ignore 
    const company = {
      tenantCode: tenantCode,
      name: companyName,
      headTitle: headTitle,
    }
    await UserCreateCompany( //@ts-ignore
      { company: company }, /** @type {import('./jsApi.GEN.js').UserCreateCompanyCallback} */
      function (/** @type {any} */ o) {
        isCreatingCompany = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return;
        }

        console.log(o);
        notifier.showSuccess('company created successfully');

        setTimeout(() => window.location.reload(), 1200);
      }
    );
  }
</script>

<svelte:window on:hashchange={onHashChange} />

{#if mode === USER}
  <MainLayout>
    <section class="create-company">
      <header>
        <h2>Create Company</h2>
        <h3>
          Use this if you have your own company you want to be associated with this email: <i>{user.email}</i>
        </h3>
      </header>
      <div class="form">
        <InputBox
          id="tenantCode"
          label="Kode Tenant / Tenant Code"
          bind:value={tenantCode}
          type="text"
          placeholder="johnxdoe"
        />
        <InputBox
          id="companyName"
          label="Nama Perusahaan / Company Name"
          bind:value={companyName}
          type="text"
          placeholder="My Company"
        />
        <InputBox
          id="headTitle"
          label="Kepala Jabatan / Head Title"
          bind:value={headTitle}
          type="text"
          placeholder="Director, CEO, President, etc"
        />
        <SubmitButton
          on:click={SubmitCreateCompany}
          isSubmitted={isCreatingCompany}
          isFullWidth
        />
      </div>
    </section>
    <!-- TODO:HABIBI list of companies i joined (see hostmapper), show in: table of tenantCode, CompanyName (link to hostmapper) -->
  </MainLayout>
{:else}
  <section class="auth-section">
    <div class="main-container">
      <div class="title-container">
        <div class="label">
          <img src="/assets/icons/benakun-logo.png" alt="Benakun" />
          <span>Benakun</span>
        </div>
        <h1>{mode.split('_').join(' ')}</h1>
      </div>
      <div class="sign-in-container">
        <div class="input-container">
          {#if mode === LOGIN || mode === REGISTER || mode === RESEND_VERIFICATION_EMAIL || mode === FORGOT_PASSWORD}
            <InputBox
              id="email"
              label="Email"
              bind:value={email}
              type="email"
            />
          {/if}
          {#if mode === LOGIN || mode === REGISTER}
            <InputBox
              id="password"
              label="Password"
              bind:value={password}
              type="password"
            />
          {/if}
          {#if mode === REGISTER}
            <InputBox
              id="confirmPass"
              label="Confirm Password"
              bind:value={confirmPass}
              type="password"
            />
          {/if}
        </div>
        <!-- Forgot Password -->
        {#if mode === LOGIN}
          <p class="forgot_password">
            Forgot Password?
            <a href="#FORGOT_PASSWORD" on:click|preventDefault={() => (mode = FORGOT_PASSWORD)}>Reset here</a>
          </p>
        {/if}
        <div class="button_container">
          {#if mode === REGISTER}
            <SubmitButton label="Register" on:click={guestRegister} isSubmitted={isSubmitted} isFullWidth />
          {/if}
          {#if mode === LOGIN}
            <SubmitButton label="Login" on:click={guestLogin} isSubmitted={isSubmitted} isFullWidth />
          {/if}
          {#if mode === RESEND_VERIFICATION_EMAIL}
            <SubmitButton label="Resend Verification Email" on:click={guestResendVerificationEmail} isSubmitted={isSubmitted} isFullWidth />
          {/if}
          {#if mode === FORGOT_PASSWORD}
            <SubmitButton label="Request Reset Password Link" on:click={guestForgotPassword} isSubmitted={isSubmitted} isFullWidth />
          {/if}
        </div>
        <!-- Oauth Buttons -->
        {#if mode === REGISTER || mode === LOGIN}
          <div class="oauth_container">
            <div class="or_separator">
              <span />
              <p>or</p>
              <span />
            </div>
            <!-- Google OAuth -->
              <a class="button" href="/">
                <img src="/assets/icons/google.svg" alt="Google" />
                <span>Continue with Google</span>
              </a>

          </div>
        {/if}
        <div class="foot_auth">
          {#if mode !== REGISTER}
            <p>Have no account? <a href="#REGISTER" on:click={() => (mode = REGISTER)}>register</a></p>
          {/if}
          {#if mode !== LOGIN}
            <p>Already have account? <a href="#LOGIN" on:click={() => (mode = LOGIN)}>login</a></p>
          {/if}
          {#if mode !== RESEND_VERIFICATION_EMAIL}
            <p>
              Email not yet verified? <a
                href="#RESEND_VERIFICATION_EMAIL"
                on:click={() => (mode = RESEND_VERIFICATION_EMAIL)}>request verification email</a
              >
            </p>
          {/if}
        </div>
      </div>
    </div>
  </section>
{/if}

<style>
  @keyframes spin {
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }

  :global(.spin) {
    animation: spin 1s cubic-bezier(0, 0, 0.2, 1) infinite;
  }

  .auth-section {
    height: 100dvh;
    width: 100%;
    background-color: var(--gray-002);
    display: flex;
    color: var(--gray-007);
  }

  .main-container {
    width: 480px;
    height: fit-content;
    padding: 20px;
    filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    border-radius: 15px;
    display: flex;
    flex-direction: column;
    background-color: white;
    margin: 50px auto;
    border: 1px solid #cbd5e1;
  }

  .title-container {
    display: flex;
    flex-direction: column;
    width: 100%;
    text-align: center;
    gap: 15px;
  }

  .title-container .label {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    gap: 8px;
  }

  .title-container .label img {
    width: 17px;
    height: auto;
    margin: 0;
  }

  .title-container .label span {
    font-size: 17px;
    font-weight: 600;
    color: var(--purple-002);
    margin-bottom: -2px;
  }

  .title-container h1 {
    margin: 5px 0 0 0;
    font-size: 22px;
    font-weight: 700;
    color: var(--blue-005);
  }

  .input-container {
    display: flex;
    flex-direction: column;
    margin-bottom: 25px;
    gap: 15px;
  }

  .forgot_password {
    margin-top: 7px;
    margin-bottom: 15px;
    width: 100%;
    text-align: center;
    font-size: 14px;
    font-weight: 600;
  }

  .forgot_password a {
    color: var(--blue-006);
    text-decoration: none;
  }

  .forgot_password a:hover {
    color: var(--blue-005);
    text-decoration: underline;
  }

  .oauth_container .or_separator {
    display: flex;
    flex-direction: row;
    align-items: center;
    width: 100%;
  }

  .oauth_container .or_separator span {
    flex-grow: 1;
    height: 0;
    border-top: 1px solid #cbd5e1;
    padding: 0;
  }

  .oauth_container .or_separator p {
    width: fit-content;
    font-weight: 600;
    padding: 0 10px;
  }

  .oauth_container .button {
    padding: 10px;
    background-color: white;
    border: 1px solid #cbd5e1;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    font-weight: 600;
    border-radius: 8px;
    text-decoration: none;
    color: #334155;
  }

  .oauth_container .button:hover {
    border: 1px solid var(--blue-006);
    background-color: var(--gray-001);
  }

  .oauth_container .button img {
    width: 20px;
    height: auto;
  }

  .oauth_container .button span {
    margin-left: 8px;
  }

  .foot_auth {
    margin-top: 10px;
    display: flex;
    flex-direction: column;
  }

  .foot_auth p {
    margin-top: 10px;
    margin-bottom: 0;
    text-align: center;
    font-weight: 600;
  }

  .foot_auth a {
    color: var(--blue-006);
    text-decoration: none;
  }

  .foot_auth a:hover {
    color: var(--blue-005);
    text-decoration: underline;
  }

  .create-company {
    display: flex;
    flex-direction: column;
    height: fit-content;
    width: 400px;
    background-color: #FFF;
    border-radius: 10px;
    border: 1px solid var(--gray-002);
    padding: 20px;
  }

  .invite_user header h2,
  .create-company header h2 {
    margin: 0;
    text-align: center;
  }

  .create-company .form {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  @media only screen and (max-width : 768px) {
    .auth-section {
      background-color: #FFF;
    }
    .main-container {
      width: 100%;
      margin: 30px auto;
      border: none;
      filter: none;
      background-color: transparent;
    }

    .create-company {
      width: 100%;
    }
  }
</style>
