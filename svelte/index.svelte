<script>
  import { TenantAdminInviteJoin, UserCreateCompany, GuestForgotPassword, GuestLogin, GuestRegister, GuestResendVerificationEmail } from './jsApi.GEN.js';
  import { onMount, tick } from 'svelte';
  import Footer from './_components/partials/Footer.svelte';
  import SideMenu from './_components/partials/SideMenu.svelte';
  import Navbar from './_components/partials/Navbar.svelte';
  import { notifier } from './_components/notifier.js';
  import InputBox from './_components/InputBox.svelte';
  import SubmitButton from './_components/SubmitButton.svelte';

  let user = {/* user */};
  let segments = {/* segments */};
  let google = '#{google}';

  function getCookie(name) {
    const match = document.cookie.match(new RegExp('(^| )' + name + '=([^;]+)'));
    if (match) return match[2];
  }
  // server state
  const title = '#{title}'; // /*! title */ {/* title */} [/* title */]
  // TODO: print session or fetch from cookie

  // local state
  let email = '', password = '', confirmPass = '';

  const LOGIN = 'LOGIN';
  const REGISTER = 'REGISTER';
  const RESEND_VERIFICATION_EMAIL = 'RESEND_VERIFICATION_EMAIL';
  const FORGOT_PASSWORD = 'FORGOT_PASSWORD';
  const USER = '';
  let mode = LOGIN, isSubmitted = false;

  async function onHashChange() {
    console.log('onHashChange.start');
    const auth = getCookie('akun');
    console.log(auth, user);
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
    console.log('onHashChange.tick');
    await tick();
  }

  onMount(() => {
    console.log('onMount.index');
    onHashChange();
    console.log('User = ', user);
  });

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
    // TODO: send to backend
    const i = { email, password };
    await GuestRegister(i, async function (o) {
      // TODO: codegen commonResponse (o.error, etc)
      // TODO: codegen list of possible errors
      console.log(o);
      if (o.error) {
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
    await GuestLogin(i, function (o) {
      console.log('o.segments=', o.segments);
      if (o.error) {
        isSubmitted = false;
        return notifier.showError(o.error);
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
    await GuestResendVerificationEmail(i, function (o) {
      console.log(o);
      if (o.error) {
        isSubmitted = false;
        return notifier.showError(o.error);
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
    await GuestForgotPassword(i, function (o) {
      console.log(o);
      if (o.error) {
        isSubmitted = false;
        return notifier.showError(o.error);
      }
      onHashChange();
      notifier.showInfo('A reset password link has been sent to your email');
    });
  }

  let emailToInvite = '', isSubmitInvite = false;
  async function inviteUser() {
    isSubmitInvite = true;
    await TenantAdminInviteJoin({email: emailToInvite}, function (o) {
      if (o.error) {
        isSubmitInvite = false;
        notifier.showError(o.error);
        return;
      }
      isSubmitInvite = false;
      notifier.showSuccess('User invited successfully');
    })
  }

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
</script>

<svelte:window on:hashchange={onHashChange} />
{#if mode === USER}
  <div class="root_layout">
    <div class="root_container">
      <SideMenu access={segments} />
      <div class="root_content">
        <Navbar {user} />
        <div class="content">
          <!-- Invite user to join company -->
          {#if user.tenantCode}
            <section class="invite_user">
              <header>
                <h2>Invite user</h2>
              </header>
              <div class="form">
                <InputBox id="emailToInvite" label="Email to invite" bind:value={emailToInvite} type="email" placeholder="user@example.com" />
                <SubmitButton on:click={inviteUser} isSubmitted={isSubmitInvite} isFullWidth />
              </div>
            </section>
          {/if}

          <!-- Create company -->
          {#if !user.tenantCode && !user.invitationState}
            <section class="create_company">
              <header>
                <h2>Create Company</h2>
              </header>
              <div class="form">
                <InputBox id="tenantCode" label="Tenant Code" bind:value={tenantCode} type="text" placeholder="axrpr" />
                <InputBox id="companyName" label="Company Name" bind:value={companyName} type="text" placeholder="My Company" />
                <InputBox id="headTitle" label="Head Title" bind:value={headTitle} type="text" placeholder="Mr. Smith" />
                <SubmitButton on:click={userCreateCompany} isSubmitted={isSubmitCreateCompany} isFullWidth />
              </div>
            </section>
          {/if}
        </div>
        <Footer />
      </div>
    </div>
  </div>
{:else}
  <section class="auth_section">
    <div class="main_container">
      <div class="title_container">
        <p>{title}</p>
        <h1>{mode.split('_').join(' ')}</h1>
      </div>
      <div class="sign_in_container">
        <div class="input_container">
          {#if mode === LOGIN || mode === REGISTER || mode === RESEND_VERIFICATION_EMAIL || mode === FORGOT_PASSWORD}
            <InputBox id="email" label="Email" bind:value={email} type="email" />
          {/if}
          {#if mode === LOGIN || mode === REGISTER}
            <InputBox id="password" label="Password" bind:value={password} type="password" />
          {/if}
          {#if mode === REGISTER}
            <InputBox id="confirmPass" label="Confirm Password" bind:value={confirmPass} type="password" />
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
              <a class="button" href="#">
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

  .auth_section {
    height: 100dvh;
    width: 100%;
    background-color: var(--gray-002);
    display: flex;
    color: var(--gray-007);
  }

  .main_container {
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

  .title_container {
    display: flex;
    flex-direction: column;
    width: 100%;
    text-align: center;
  }

  .title_container p {
    font-size: 16px;
    font-weight: 600;
    color: var(--purple-002);
    margin: 0;
  }

  .title_container h1 {
    margin: 5px 0 0 0;
    font-size: 22px;
    font-weight: 700;
  }

  .input_container {
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

  .invite_user,
  .create_company {
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

  .invite_user header h2,
  .create_company header h2 {
    margin: 0;
    text-align: center;
  }

  .invite_user .form,
  .create_company .form {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }
</style>
