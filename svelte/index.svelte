<script>
  //@ts-nocheck
  import { GuestForgotPassword, GuestLogin, GuestRegister, GuestResendVerificationEmail } from './jsApi.GEN.js';
  import { onMount, tick } from 'svelte';
  import FaSolidCircleNotch from 'svelte-icons-pack/fa/FaSolidCircleNotch';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import Footer from './_components/partials/Footer.svelte';
  import SideMenu from './_components/partials/SideMenu.svelte';
  import Navbar from './_components/partials/Navbar.svelte';
  import { notifier } from './_components/notifier.js';
  import InputBox from './_components/InputBox.svelte';
  import ToCompany from './_components/ToCompany.svelte';

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
</script>

<svelte:window on:hashchange={onHashChange} />
{#if mode === USER}
  <div class="root_layout">
    <div class="root_container">
      <SideMenu access={segments} />
      <div class="root_content">
        <Navbar {user} />
        <div class="content">
          <ToCompany {user} />
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
            <button on:click={guestRegister}>
              {#if isSubmitted === true}
                <Icon className="spin" color="#FFF" size={15} src={FaSolidCircleNotch} />
              {/if}
              {#if isSubmitted === false}
                <span>Register</span>
              {/if}
            </button>
          {/if}
          {#if mode === LOGIN}
            <button on:click={guestLogin}>
              {#if isSubmitted === true}
                <Icon className="spin" color="#FFF" size={15} src={FaSolidCircleNotch} />
              {/if}
              {#if isSubmitted === false}
                <span>Login</span>
              {/if}
            </button>
          {/if}
          {#if mode === RESEND_VERIFICATION_EMAIL}
            <button on:click={guestResendVerificationEmail}>
              {#if isSubmitted === true}
                <Icon className="spin" color="#FFF" size={15} src={FaSolidCircleNotch} />
              {/if}
              {#if isSubmitted === false}
                <span>Resend Verification Email</span>
              {/if}
            </button>
          {/if}
          {#if mode === FORGOT_PASSWORD}
            <button on:click={guestForgotPassword}>
              {#if isSubmitted === true}
                <Icon className="spin" color="#FFF" size={15} src={FaSolidCircleNotch} />
              {/if}
              {#if isSubmitted === false}
                <span>Request Reset Password Link</span>
              {/if}
            </button>
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
            {#if google}
              <a class="button" href={google}>
                <img src="/assets/icons/google.svg" alt="Google" />
                <span>Continue with Google</span>
              </a>
            {/if}
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

  .button_container button {
    margin: 0;
    width: 100%;
    padding: 10px;
    font-size: 16px;
    font-weight: 700;
    background-color: var(--blue-006);
    border-radius: 8px;
    color: white;
    border: none;
    cursor: pointer;
    filter: drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
  }

  .button_container button:hover {
    background-color: var(--blue-005);
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
    background-color: var(--gray-002);
    /* #94a3b8 */
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
</style>
