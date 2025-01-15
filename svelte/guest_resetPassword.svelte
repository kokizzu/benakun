<script>
    import {GuestResetPassword} from './jsApi.GEN';
    import {notifier} from './_components/xNotifier.js'
    import InputBox from './_components/InputBox.svelte';

    let password = '';
    let pass2 = '';

    async function resetPassword() {
        if (password.length < 12) {
            return notifier.showWarning('password must be at least 12 characters');
        }
        if (password !== pass2) {
            return notifier.showWarning('password confirmation does not match');
        }
        const queryParam = window.location.href.split('?')[1];
        const qps = queryParam.split('&');
        let secretCode = '';
        let hash = '';
        for (let i = 0; i < qps.length; i++) {
            const [key, value] = qps[i].split('=');
            if (key === 'secretCode') secretCode = value;
            if (key === 'hash') hash = value;
        }
        let i = {secretCode, hash, password}; //@ts-ignore
        await GuestResetPassword(i, function(/** @type {any} */ o) {
            console.log(o);
            if (o.error) {
                return notifier.showError(o.error);
            }
            notifier.showSuccess('password reset successful');
            window.location.href = '/';
        });
    }
</script>


<section class="reset-password-container">
  <div class="main-content">
    <h1>
        <i class="gg-lock"/>
        <span>Reset Password</span>
    </h1>
    <div class="input-container">
      <InputBox
        id="newPass"
        type="password"
        label="New Password"
        bind:value={password}
      />
      <InputBox
        id="confirmPass"
        type="password"
        label="Confirm New Password"
        bind:value={pass2}
      />
    </div>
    <button on:click={resetPassword}>Reset Password</button>
  </div>
</section>

<style>
  .reset-password-container {
    height           : 100%;
    width            : 100%;
    background-color : #F1F5F9;
    display          : flex;
    color            : #475569;
  }
  .reset-password-container .main-content {
    width            : 420px;
    height           : fit-content;
    padding          : 20px;
    filter           : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    border-radius    : 15px;
    display          : flex;
    flex-direction   : column;
    gap: 20px;
    background-color : white;
    margin           : 50px auto;
    border           : 1px solid #CBD5E1;
  }

  .reset-password-container .main-content h1 {
    font-weight     : 700;
    font-size       : 22px;
    margin          : 0 auto 15px;
    display         : flex;
    flex-direction  : row;
    align-content   : center;
    justify-content : center;
    align-items     : center;
  }

  .reset-password-container .main-content h1 i {
    margin-right : 15px;
  }

  .reset-password-container .main-content .input-container {
    display: flex;
    flex-direction: column;
    gap: 15px;
  }

  .reset-password-container .main-content button {
    padding          : 12px 0;
    font-size        : 16px;
    border           : none;
    background-color : var(--blue-006);
    border-radius    : 8px;
    color            : white;
    cursor           : pointer;
    text-decoration  : none;
    width            : 100%;
    margin           : 10px auto 0 auto;
  }

  .reset-password-container .main-content button:hover {
    background-color : var(--blue-005);
  }
</style>
