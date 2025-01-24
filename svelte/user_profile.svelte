<script>
  /** @typedef {import('./_components/types/user').User} User */
  /** @typedef {import('./_components/types/access').Access} Access */
  /** @typedef {import('./_components/types/user').Session} Session */
  /** @typedef {import('./_components/types/invoicePayment').InvoicePayment} InvoicePayment */

  import { notifier } from './_components/xNotifier.js';
  import { localeDatetime } from './_components/xformatter.js';
  import { UserChangePassword, UserUpdateProfile } from './jsApi.GEN.js';
  import InputBox from './_components/InputBox.svelte';
  import SubmitButton from './_components/SubmitButton.svelte';
  import MainLayout from './_layouts/mainLayout.svelte';
  import ProfileSubMenu from './_components/partials/ProfileSubMenu.svelte';

  let user      = /** @type {User} */ ({/* user */});
  let segments  = /** @type {Access} */ ({/* segments */});

  let oldPassword = '';
  let newPassword = '';
  let repeatNewPassword = '';
  let oldProfileJson = '';

  let isProfileSubmitted = false;
  let isPasswordSubmitted = false;

  async function updateProfile() {
    isProfileSubmitted = true;
    if (JSON.stringify(user) === oldProfileJson) return notifier.showWarning('No changes');
    // @ts-ignore
    await UserUpdateProfile(user, function (res) {
      isProfileSubmitted = false;
      // @ts-ignore
      if (res.error) return notifier.showError(res.error);
      // @ts-ignore
      oldProfileJson = JSON.stringify(res.user);
      // @ts-ignore
      user = res.user;
      notifier.showSuccess('Profile updated');
    });
  }

  async function changePassword() {
    isPasswordSubmitted = true;
    if (newPassword !== repeatNewPassword) {
      isPasswordSubmitted = false;
      return notifier.showWarning('New password and repeat new password must be same');
    }
    let input = {
      oldPass: oldPassword,
      newPass: newPassword,
    };
    // @ts-ignore
    await UserChangePassword(input, function (res) {
      isPasswordSubmitted = false;
      // @ts-ignore
      if (res.error) return notifier.showError(res.error);
      oldPassword = '';
      newPassword = '';
      repeatNewPassword = '';
      notifier.showSuccess('Password updated');
    });
  }
</script>

<MainLayout>
  <ProfileSubMenu />
  <div class="user-details">
    <div class="profile">
      <h3>Profile Details</h3>
      <div class="input-row">
        <InputBox id="fullname" label="Full Name" bind:value={user.fullName} type="text" />
        <InputBox id="email" label="E-Mail" bind:value={user.email} type="email" />
      </div>
      <div class="user_info">
        <div class="info">
          <span>Join At</span>
          <span>{localeDatetime(user.createdAt)}</span>
        </div>
        <div class="info">
          <span>User ID</span>
          <span>#{user.id}</span>
        </div>
        <div class="info">
          <span>Role</span>
          <span>{user.role || `--`}</span>
        </div>
        <div class="info">
          <span>Last Login</span>
          <span>{localeDatetime(user.lastLoginAt) || '--'}</span>
        </div>
        <div class="info">
          <span>Updated At</span>
          <span>{localeDatetime(user.updatedAt) || '--'}</span>
        </div>
        <div class="info">
          <span>Verified At</span>
          <span>{localeDatetime(user.verifiedAt) || '--'}</span>
        </div>
        <div class="info">
          <span>Last Updated Password</span>
          <span>{localeDatetime(user.passwordSetAt) || '--'}</span>
        </div>
        <div class="info">
          <span>Tenant Code</span>
          <span class="tenant-code">{user.tenantCode || `--`}</span>
        </div>
        <div class="info">
          <span>Support Expired At</span>
          <span>{localeDatetime(user.supportExpiredAt) || '--'}</span>
        </div>
      </div>
      <SubmitButton
        on:click={updateProfile}
        isSubmitted={isProfileSubmitted}
        isFullWidth={false}
      />
    </div>
    <div class="password-set">
      <h3>Change Password</h3>
      <InputBox id="oldPassword" label="Old Password" value={oldPassword} type="password" />
      <InputBox id="newPassword" label="New Password" value={newPassword} type="password" />
      <InputBox id="repeatNewPassword" label="Repeat New Password" value={repeatNewPassword} type="password" />
      <SubmitButton
        on:click={changePassword}
        isSubmitted={isPasswordSubmitted}
        isFullWidth={false}
      />
    </div>
  </div>
</MainLayout>


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

  .user-details h3 {
    margin: 0;
    font-size: var(--font-lg);
    width: 100%;
    text-align: left;
  }

  .user-details {
    display: flex;
    flex-direction: row;
    gap: 30px;
  }

  .user-details .profile,
  .user-details .password-set {
    padding: 20px;
    border: 1px solid var(--gray-003);
    display: flex;
    flex-direction: column;
    gap: 20px;
    height: fit-content;
    border-radius: 8px;
    background-color: #FFF;
  }

  .user-details .profile {
    flex-grow: 1;
  }

  .user-details .profile .input-row {
    display: flex;
    flex-direction: row;
    gap: 20px;
  }

  .user-details .profile .user_info {
    border: 1px solid var(--gray-003);
    display: flex;
    flex-direction: column;
    height: fit-content;
    border-radius: 8px;
    overflow: hidden;
  }

  .user-details .profile .user_info .info {
    border-bottom: 1px solid var(--gray-003);
    padding: 10px 15px;
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
    text-transform: capitalize;
  }

  .user-details .profile .user_info .info span:first-child {
    font-weight: 700;
    width: 200px;
  }

  .user-details .profile .user_info .info:nth-child(odd) {
    background-color: var(--gray-001);
  }

  .user-details .profile .user_info .info:last-child {
    border-bottom: none;
  }

  .user-details .profile .user_info .info .tenant-code {
    text-transform: lowercase !important;
  }

  .user-details .password-set {
    width: 400px;
  }

  @media only screen and (max-width : 768px) {
    .user-details {
      flex-direction: column;
      gap: 20px;
    }
  }
</style>
