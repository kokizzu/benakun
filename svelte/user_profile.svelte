<script>
  import { notifier } from './_components/notifier.js';
  import Icon from 'svelte-icons-pack/Icon.svelte';
  import { datetime } from './_components/formatter';
  import { onMount } from 'svelte';
  import { UserChangePassword, UserSessionKill, UserSessionsActive, UserUpdateProfile } from './jsApi.GEN.js';
  import FaSolidTimes from 'svelte-icons-pack/fa/FaSolidTimes';
  import InputBox from './_components/InputBox.svelte';
  import SubmitButton from './_components/SubmitButton.svelte';
  import { localeDatetime } from './_components/formatter.js';
  import MainLayout from './_layouts/mainLayout.svelte';

  let user = {/* user */};
  let segments = {/* segments */};
  let sessionActiveLists = [/* activeSessions */];
  let oldPassword = '';
  let newPassword = '';
  let repeatNewPassword = '';
  let oldProfileJson = '';
  let profileSubmit = false,passwordSubmit = false;

  onMount(async () => {
    console.log('onMount.user');
    oldProfileJson = JSON.stringify(user);
    console.log('User data = ', user);
  });

  async function updateProfile() {
    profileSubmit = true;
    if (JSON.stringify(user) === oldProfileJson) return notifier.showWarning('No changes');
    // @ts-ignore
    await UserUpdateProfile(user, function (res) {
      profileSubmit = false;
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
    passwordSubmit = true;
    if (newPassword !== repeatNewPassword) {
      passwordSubmit = false;
      return notifier.showWarning('New password and repeat new password must be same');
    }
    let input = {
      oldPass: oldPassword,
      newPass: newPassword,
    };
    // @ts-ignore
    await UserChangePassword(input, function (res) {
      passwordSubmit = false;
      // @ts-ignore
      if (res.error) return notifier.showError(res.error);
      oldPassword = '';
      newPassword = '';
      repeatNewPassword = '';
      notifier.showSuccess('Password updated');
    });
  }

  async function killSession(sessionToken) {
    await UserSessionKill({ sessionTokenHash: sessionToken }, async res => {
      // @ts-ignore
      if (res.error) return notifier.showError(res.error);
      if (res.sessionTerminated < 1) return notifier.showError('No session terminated');
      notifier.showInfo(res.sessionTerminated + ' session terminated');
      // @ts-ignore
      await UserSessionsActive({ userId: user.id }, res => {
        // @ts-ignore
        if (res.error) return notifier.showError(res.error);
        sessionActiveLists = res.sessionsActive;
      });
    });
  }
</script>

<MainLayout>
  <div class="user_details_container">
    <div class="profile_details">
      <h3>Profile Details</h3>
      <div class="input_row">
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
          <span>{user.id}</span>
        </div>
        <div class="info">
          <span>Role</span>
          <span>{user.role || `--`}</span>
        </div>
        <div class="info">
          <span>Last Login</span>
          <span>{localeDatetime(user.lastLoginAt) || 0}</span>
        </div>
        <div class="info">
          <span>Updated At</span>
          <span>{localeDatetime(user.updatedAt) || 0}</span>
        </div>
        <div class="info">
          <span>Verified At</span>
          <span>{localeDatetime(user.verifiedAt) || 0}</span>
        </div>
        <div class="info">
          <span>Last Updated Password</span>
          <span>{localeDatetime(user.passwordSetAt) || 0}</span>
        </div>
        <div class="info">
          <span>Tenant Code</span>
          <span>{user.tenantCode || `--`}</span>
        </div>
      </div>
      <SubmitButton
        on:click={updateProfile}
        isSubmitted={profileSubmit}
        isFullWidth={false}
      />
    </div>
    <div class="password_set">
      <h3>Change Password</h3>
      <InputBox id="oldPassword" label="Old Password" value={oldPassword} type="password" />
      <InputBox id="newPassword" label="New Password" value={newPassword} type="password" />
      <InputBox id="repeatNewPassword" label="Repeat New Password" value={repeatNewPassword} type="password" />
      <SubmitButton
        on:click={changePassword}
        isSubmitted={passwordSubmit}
        isFullWidth={false}
      />
    </div>
  </div>
  <div class="sessions_container">
    <div class="header">
      <span>IP Address</span>
      <span>Expired At</span>
      <span>Device</span>
    </div>
    <div class="sessions">
      {#if sessionActiveLists && sessionActiveLists.length}
        {#each sessionActiveLists as session}
          <div class="session">
            <span>{session.loginIPs || 'no-data'}</span>
            <span>{datetime(session.expiredAt) || 0}</span>
            <span>{session.device || 'no-data'}</span>
            <button
              on:click={() => killSession(session.sessionToken)}
              class="kill_session"
              title="Kill this session"
            >
              <Icon color="#FFF" size="12" src={FaSolidTimes} />
            </button>
          </div>
        {/each}
      {/if}
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

  .user_details_container h3 {
    margin: 0;
    font-size: var(--font-lg);
    width: 100%;
    text-align: left;
  }

  /* +============ /// ===========+ */

  .user_details_container {
    display: flex;
    flex-direction: row;
    gap: 30px;
  }

  .user_details_container .profile_details,
  .user_details_container .password_set {
    padding: 20px;
    border: 1px solid var(--gray-003);
    display: flex;
    flex-direction: column;
    gap: 20px;
    height: fit-content;
    border-radius: 8px;
    background-color: #FFF;
  }

  .user_details_container .profile_details {
    flex-grow: 1;
  }

  .user_details_container .profile_details .input_row {
    display: flex;
    flex-direction: row;
    gap: 20px;
  }

  .user_details_container .profile_details .user_info {
    border: 1px solid var(--gray-003);
    display: flex;
    flex-direction: column;
    height: fit-content;
    border-radius: 8px;
    overflow: hidden;
  }

  .user_details_container .profile_details .user_info .info {
    border-bottom: 1px solid var(--gray-003);
    padding: 10px 15px;
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 10px;
    text-transform: capitalize;
  }

  .user_details_container .profile_details .user_info .info span:first-child {
    font-weight: 700;
    width: 200px;
  }

  .user_details_container .profile_details .user_info .info:nth-child(odd),
  .sessions_container .sessions .session:nth-child(odd) {
    background-color: var(--gray-001);
  }

  .user_details_container .profile_details .user_info .info:last-child {
    border-bottom: none;
  }

  .user_details_container .password_set {
    width: 400px;
  }

  /* +=======================================+ */

  .sessions_container {
    border: 1px solid var(--gray-003);
    display: flex;
    flex-direction: column;
    height: fit-content;
    border-radius: 8px;
    overflow: hidden;
    margin-top: 30px;
    background-color: #FFF;
  }

  .sessions_container .header {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    font-weight: bold;
    padding: 12px 15px;
    border-bottom: 1px solid #cbd5e1;
  }

  .sessions_container .sessions {
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .sessions_container .sessions .session {
    text-align: left;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    padding: 12px 15px;
    position: relative;
    width: 100%;
  }

  .sessions_container .sessions .session span:nth-child(3),
  .sessions_container .header span:nth-child(3) {
    flex-grow: 1;
    margin-right: 30px;
  }

  .sessions_container .sessions .session span,
  .sessions_container .header span {
    width: 200px;
  }

  .sessions_container .sessions .session .kill_session {
    border: none;
    background-color: var(--red-002);
    padding: 5px;
    border-radius: 50%;
    position: absolute;
    right: 15px;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .sessions_container .sessions .session .kill_session:hover {
    background-color: var(--red-001);
  }

  /* Responsive to mobile device */
  @media (max-width: 768px) {
  }
</style>
